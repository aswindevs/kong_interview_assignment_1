package logger

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Interface defines the logging methods
type Interface interface {
	Debug(ctx context.Context, message string, args ...interface{})
	Info(ctx context.Context, message string, args ...interface{})
	Warn(ctx context.Context, message string, args ...interface{})
	Error(ctx context.Context, message string, args ...interface{})
	Fatal(ctx context.Context, message string, args ...interface{})
}

// Logger wraps zap logger
type Logger struct {
	logger *zap.Logger
	format Format
}

type Format string

const (
	JSON    Format = "json"
	Console Format = "console"
)

const TraceIDKey = "trace_id"

var _ Interface = (*Logger)(nil)

// New creates a new logger instance
func New(format string, logLevel string) (*Logger, error) {
	var config zap.Config
	logFormat := Format(format)

	switch logFormat {
	case Console:
		config = zap.NewDevelopmentConfig()
		config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
		config.Encoding = "console"
		config.EncoderConfig.ConsoleSeparator = " "
	case JSON:
		config = zap.NewProductionConfig()
	default:
		config = zap.NewProductionConfig()
	}

	level, err := zap.ParseAtomicLevel(logLevel)
	if err != nil {
		level = zap.NewAtomicLevelAt(zapcore.InfoLevel)
	}

	config.Level = level
	config.EncoderConfig.TimeKey = "timestamp"
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	logger, err := config.Build(zap.AddCallerSkip(1))
	if err != nil {
		return nil, fmt.Errorf("failed to build logger: %w", err)
	}

	return &Logger{
		logger: logger,
		format: logFormat,
	}, nil
}

// getTraceIDFromContext extracts trace ID from context
func getTraceIDFromContext(ctx context.Context) string {
	if ctx == nil {
		return ""
	}
	// Get trace ID from gin context
	if gc, ok := ctx.(*gin.Context); ok {
		if traceID := gc.GetString(TraceIDKey); traceID != "" {
			return traceID
		}
	}
	// Fallback to context value
	if traceID, ok := ctx.Value(TraceIDKey).(string); ok {
		return traceID
	}
	return ""
}

// log is a common logging function that handles different log levels
func (l *Logger) log(ctx context.Context, level zapcore.Level, message string, args ...interface{}) {
	fields := l.buildFields(ctx, args...)

	switch level {
	case zapcore.DebugLevel:
		l.logger.Debug(message, fields...)
	case zapcore.InfoLevel:
		l.logger.Info(message, fields...)
	case zapcore.WarnLevel:
		l.logger.Warn(message, fields...)
	case zapcore.ErrorLevel:
		l.logger.Error(message, fields...)
	case zapcore.FatalLevel:
		l.logger.Fatal(message, fields...)
	}
}

// Debug logs a debug message
func (l *Logger) Debug(ctx context.Context, message string, args ...interface{}) {
	l.log(ctx, zapcore.DebugLevel, message, args...)
}

// Info logs an info message
func (l *Logger) Info(ctx context.Context, message string, args ...interface{}) {
	l.log(ctx, zapcore.InfoLevel, message, args...)
}

// Warn logs a warning message
func (l *Logger) Warn(ctx context.Context, message string, args ...interface{}) {
	l.log(ctx, zapcore.WarnLevel, message, args...)
}

// Error logs an error message
func (l *Logger) Error(ctx context.Context, message string, args ...interface{}) {
	l.log(ctx, zapcore.ErrorLevel, message, args...)
}

// Fatal logs a fatal message and exits
func (l *Logger) Fatal(ctx context.Context, message string, args ...interface{}) {
	l.log(ctx, zapcore.FatalLevel, message, args...)
}

func (l *Logger) buildFields(ctx context.Context, args ...interface{}) []zap.Field {
	fields := make([]zap.Field, 0, len(args)/2+1)

	if traceID := getTraceIDFromContext(ctx); traceID != "" {
		fields = append(fields, zap.String("trace_id", traceID))
	}

	for i := 0; i < len(args); i += 2 {
		key, ok := args[i].(string)
		if !ok {
			continue
		}
		if i+1 >= len(args) {
			continue
		}
		val := args[i+1]

		switch v := val.(type) {
		case error:
			fields = append(fields, zap.Error(v))
		default:
			fields = append(fields, zap.Any(key, v))
		}
	}
	return fields
}
