package usecase_test

import (
	"context"
	"testing"
	"time"

	"github.com/aswindevs/kong_interview-assignment_1/internal/dto"
	"github.com/aswindevs/kong_interview-assignment_1/internal/entity"
	"github.com/aswindevs/kong_interview-assignment_1/internal/usecase"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
)

type MockServiceRepo struct {
	mock.Mock
}

func (m *MockServiceRepo) FindAll(ctx context.Context, search string, sortBy string, OrderBy string, page int, pageSize int) ([]entity.Service, int, error) {
	args := m.Called(ctx, search, sortBy, OrderBy, page, pageSize)
	return args.Get(0).([]entity.Service), args.Int(1), args.Error(2)
}

func (m *MockServiceRepo) FindById(ctx context.Context, id int) (*entity.Service, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.Service), args.Error(1)
}
func (m *MockServiceRepo) CreateVersion(ctx context.Context, serviceVersion entity.ServiceVersion) (*entity.ServiceVersion, error) {
	args := m.Called(ctx, serviceVersion)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.ServiceVersion), args.Error(1)
}

func (m *MockServiceRepo) FindAllVersionsById(ctx context.Context, serviceId int) ([]entity.ServiceVersion, error) {
	args := m.Called(ctx, serviceId)
	return args.Get(0).([]entity.ServiceVersion), args.Error(1)
}
func (m *MockServiceRepo) Create(ctx context.Context, service entity.Service) (*entity.Service, error) {
	args := m.Called(ctx, service)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.Service), args.Error(1)
}
func (m *MockServiceRepo) Update(ctx context.Context, id int, service entity.Service) (*entity.Service, error) {
	args := m.Called(ctx, id, service)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.Service), args.Error(1)
}
func (m *MockServiceRepo) Delete(ctx context.Context, id int) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

type ServiceCatalogUsecaseTestSuite struct {
	suite.Suite
	serviceRepo *MockServiceRepo
	usecase     *usecase.ServiceCatalogUsecase
}

func (suite *ServiceCatalogUsecaseTestSuite) SetupTest() {
	suite.serviceRepo = new(MockServiceRepo)
	suite.usecase = usecase.NewServiceCatalogUsecase(suite.serviceRepo)
}

func (suite *ServiceCatalogUsecaseTestSuite) TestCreateService_Success() {
	mockService := dto.ServiceCatalogRequest{
		Name:        "Test Service",
		Description: "A service for testing",
	}
	expectedService := &entity.Service{
		Model: gorm.Model{
			ID:        1,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		Name:        "Test Service",
		Description: "A service for testing",
		Versions: []entity.ServiceVersion{
			{
				Version: 1,
			},
		},
	}
	suite.serviceRepo.On("Create", mock.Anything, mock.AnythingOfType("entity.Service")).Return(expectedService, nil)

	createdService, err := suite.usecase.CreateService(context.Background(), mockService)

	suite.NoError(err)
	suite.NotNil(createdService)
	suite.Equal(expectedService.Name, createdService.Name)
	suite.serviceRepo.AssertExpectations(suite.T())
}

func (suite *ServiceCatalogUsecaseTestSuite) TestGetServiceById_Success() {
	expectedService := &entity.Service{
		Model: gorm.Model{
			ID: 1,
		},
		Name: "Test Service",
	}
	suite.serviceRepo.On("FindById", mock.Anything, 1).Return(expectedService, nil)

	service, err := suite.usecase.GetServiceById(context.Background(), 1)

	suite.NoError(err)
	suite.NotNil(service)
	suite.Equal(expectedService.ID, service.ID)
	suite.serviceRepo.AssertExpectations(suite.T())
}

func TestServiceCatalogUsecase(t *testing.T) {
	suite.Run(t, new(ServiceCatalogUsecaseTestSuite))
}
