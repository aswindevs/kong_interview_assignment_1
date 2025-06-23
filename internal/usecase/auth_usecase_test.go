package usecase_test

import (
	"context"
	"testing"

	"github.com/aswindevs/kong_interview-assignment_1/config"
	"github.com/aswindevs/kong_interview-assignment_1/internal/entity"
	"github.com/aswindevs/kong_interview-assignment_1/internal/usecase"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"golang.org/x/crypto/bcrypt"
)

// MockUserRepo is a mock implementation of the UserRepo interface
type MockUserRepo struct {
	mock.Mock
}

func (m *MockUserRepo) FindById(ctx context.Context, id int) (*entity.User, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.User), args.Error(1)
}

func (m *MockUserRepo) FindByEmail(ctx context.Context, email string) (*entity.User, error) {
	args := m.Called(ctx, email)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.User), args.Error(1)
}

func (m *MockUserRepo) FindPermissionsByUserId(ctx context.Context, userID int) ([]entity.Permission, error) {
	args := m.Called(ctx, userID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]entity.Permission), args.Error(1)
}

type AuthUsecaseTestSuite struct {
	suite.Suite
	userRepo *MockUserRepo
	usecase  *usecase.AuthUsecase
	cfg      *config.Auth
}

func (suite *AuthUsecaseTestSuite) SetupTest() {
	suite.userRepo = new(MockUserRepo)
	suite.cfg = &config.Auth{
		SecretKey:           "test-secret",
		TokenExpirationTime: 1,
	}
	suite.usecase = usecase.NewAuthUsecase(suite.userRepo, suite.cfg)
}

func (suite *AuthUsecaseTestSuite) TestLogin_Success() {
	password := "password"
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	suite.NoError(err)

	mockUser := &entity.User{
		Email:    "test@example.com",
		Password: string(hashedPassword),
	}

	suite.userRepo.On("FindByEmail", mock.Anything, "test@example.com").Return(mockUser, nil)

	token, err := suite.usecase.Login(context.Background(), "test@example.com", password)

	suite.NoError(err)
	suite.NotEmpty(token)
	suite.userRepo.AssertExpectations(suite.T())
}

func TestAuthUsecase(t *testing.T) {
	suite.Run(t, new(AuthUsecaseTestSuite))
}
