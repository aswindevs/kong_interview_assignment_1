package usecase_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/aswindevs/kong_interview-assignment_1/internal/entity"
	"github.com/aswindevs/kong_interview-assignment_1/internal/usecase"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type UserUsecaseTestSuite struct {
	suite.Suite
	userRepo *MockUserRepo
	usecase  *usecase.UserUsecase
}

func (suite *UserUsecaseTestSuite) SetupTest() {
	suite.userRepo = new(MockUserRepo)
	suite.usecase = usecase.NewUserUsecase(suite.userRepo)
}

func (suite *UserUsecaseTestSuite) TestGetUserById_Success() {
	mockUser := &entity.User{
		Email: "test@example.com",
	}
	suite.userRepo.On("FindById", mock.Anything, 1).Return(mockUser, nil)

	user, err := suite.usecase.GetUserById(context.Background(), 1)

	suite.NoError(err)
	suite.NotNil(user)
	suite.Equal("test@example.com", user.Email)
	suite.userRepo.AssertExpectations(suite.T())
}

func (suite *UserUsecaseTestSuite) TestGetUserById_NotFound() {
	suite.userRepo.On("FindById", mock.Anything, 1).Return(nil, fmt.Errorf("not found"))

	user, err := suite.usecase.GetUserById(context.Background(), 1)

	suite.Error(err)
	suite.Nil(user)
	suite.userRepo.AssertExpectations(suite.T())
}

func TestUserUsecase(t *testing.T) {
	suite.Run(t, new(UserUsecaseTestSuite))
}
