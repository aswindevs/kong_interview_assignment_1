package usecase

import (
	"context"

	// "fmt"
	"time"

	"github.com/aswindevs/kong_interview-assignment_1/config"
	"github.com/aswindevs/kong_interview-assignment_1/internal/entity"
	appError "github.com/aswindevs/kong_interview-assignment_1/internal/errors"
	"github.com/aswindevs/kong_interview-assignment_1/internal/usecase/repo"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	// appErr "github.com/aswindevs/kong_interview-assignment_1/internal/errors"
)

type AuthUsecase struct {
	userRepo repo.UserRepo

	cfg *config.Auth
}

func NewAuthUsecase(userRepo repo.UserRepo, cfg *config.Auth) *AuthUsecase {
	return &AuthUsecase{
		userRepo: userRepo,
		cfg:      cfg,
	}
}

func (a *AuthUsecase) Login(ctx context.Context, email string, password string) (string, error) {
	user, err := a.userRepo.FindByEmail(ctx, email)
	if err != nil || user == nil {
		return "", appError.NewAuthenticationError("Email or password wrong")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", appError.NewAuthenticationError("Email or password wrong")

	}

	return a.generateJWT(user)
}
func (a *AuthUsecase) generateJWT(user *entity.User) (string, error) {
	claims := jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Hour * time.Duration(a.cfg.TokenExpirationTime)).Unix(),
		"iat":     time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(a.cfg.SecretKey))
}

func (a *AuthUsecase) GetPermissionsByUserId(ctx context.Context, userId int) ([]entity.Permission, error) {
	permissions, err := a.userRepo.FindPermissionsByUserId(ctx, userId)
	if err != nil {
		return nil, err
	}
	return permissions, nil
}
