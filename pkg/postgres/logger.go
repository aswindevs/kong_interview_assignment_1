package postgres

import (
	"context"
	"fmt"
	"time"

	"github.com/aswindevs/kong_interview-assignment_1/pkg/logger"

	gormlogger "gorm.io/gorm/logger"
)

type GormLogger struct {
	logger        logger.Interface
	level         gormlogger.LogLevel
	slowThreshold time.Duration
}

func NewGormLogger(logger logger.Interface, level gormlogger.LogLevel) *GormLogger {
	return &GormLogger{
		logger:        logger,
		level:         level,
		slowThreshold: 200 * time.Millisecond,
	}
}

func (l *GormLogger) LogMode(level gormlogger.LogLevel) gormlogger.Interface {
	newLogger := *l
	newLogger.level = level
	return &newLogger
}

func (l *GormLogger) Info(ctx context.Context, msg string, data ...interface{}) {
	if l.level >= gormlogger.Info {
		l.logger.Info(ctx, fmt.Sprintf(msg, data...))
	}
}

func (l *GormLogger) Warn(ctx context.Context, msg string, data ...interface{}) {
	if l.level >= gormlogger.Warn {
		l.logger.Warn(ctx, fmt.Sprintf(msg, data...))
	}
}

func (l *GormLogger) Error(ctx context.Context, msg string, data ...interface{}) {
	if l.level >= gormlogger.Error {
		l.logger.Error(ctx, fmt.Sprintf(msg, data...))
	}
}

func (l *GormLogger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	if l.level <= gormlogger.Silent {
		return
	}

	elapsed := time.Since(begin)
	sql, rows := fc()

	if elapsed > l.slowThreshold {
		l.logger.Warn(ctx, fmt.Sprintf("Slow SQL: %v | Elapsed: %v | Rows: %v", sql, elapsed, rows))
		return
	}

	l.logger.Debug(ctx, fmt.Sprintf("SQL: %v | Elapsed: %v | Rows: %v", sql, elapsed, rows))
}

// Convert string level to GORM LogLevel
func getGormLogLevel(level string) gormlogger.LogLevel {
	switch level {
	case "debug":
		return gormlogger.Info // GORM's most verbose level
	case "info":
		return gormlogger.Info
	case "warn":
		return gormlogger.Warn
	case "error":
		return gormlogger.Error
	case "silent":
		return gormlogger.Silent
	default:
		return gormlogger.Silent // Default to less verbose
	}
}
