package usecase

import (
	"context"
	"strconv"

	"github.com/aswindevs/kong_interview-assignment_1/internal/dto"
	"github.com/aswindevs/kong_interview-assignment_1/internal/entity"
	appError "github.com/aswindevs/kong_interview-assignment_1/internal/errors"
	"github.com/aswindevs/kong_interview-assignment_1/internal/usecase/repo"
)

type ServiceCatalogUsecase struct {
	ServiceRepo repo.ServiceRepo
}

func NewServiceCatalogUsecase(r repo.ServiceRepo) *ServiceCatalogUsecase {
	return &ServiceCatalogUsecase{
		ServiceRepo: r,
	}
}

func (uc *ServiceCatalogUsecase) GetAllServices(ctx context.Context, search string, sortBy string, orderBy string, page int, pageSize int) ([]entity.Service, int, error) {
	services, total, err := uc.ServiceRepo.FindAll(ctx, search, sortBy, orderBy, page, pageSize)
	if err != nil {
		return nil, 0, appError.NewNotFoundError(err.Error())
	}
	return services, total, nil
}

func (uc *ServiceCatalogUsecase) GetServiceById(ctx context.Context, id int) (*entity.Service, error) {
	service, err := uc.ServiceRepo.FindById(ctx, id)
	if err != nil {
		return nil, appError.NewNotFoundError(err.Error())
	}
	return service, nil
}

func (uc *ServiceCatalogUsecase) CreateService(ctx context.Context, service dto.ServiceCatalogRequest) (*entity.Service, error) {
	serviceEntity := entity.Service{
		Name:        service.Name,
		Description: service.Description,
		Versions: []entity.ServiceVersion{
			{
				Version: 1,
			},
		},
	}
	createdService, err := uc.ServiceRepo.Create(ctx, serviceEntity)
	if err != nil {
		return nil, appError.NewAlreadyExistsError(err.Error())
	}
	return createdService, nil
}

func (uc *ServiceCatalogUsecase) UpdateService(ctx context.Context, id int, service entity.Service) (*entity.Service, error) {
	updatedService, err := uc.ServiceRepo.Update(ctx, id, service)
	if err != nil {
		return nil, appError.NewNotFoundError(err.Error())
	}
	return updatedService, nil
}

func (uc *ServiceCatalogUsecase) DeleteService(ctx context.Context, id int) error {
	err := uc.ServiceRepo.Delete(ctx, id)
	if err != nil {
		return appError.NewNotFoundError(err.Error())
	}
	return nil
}

func (uc *ServiceCatalogUsecase) CreateServiceVersion(ctx context.Context, id int, serviceVersion dto.ServiceVersionRequest) (*entity.ServiceVersion, error) {
	version, err := strconv.Atoi(serviceVersion.Version)
	if err != nil {
		return nil, appError.NewBadRequestError("invalid version format provided")
	}
	serviceVersionEntity := entity.ServiceVersion{
		ServiceID: uint(id),
		Version:   version,
	}
	createdServiceVersion, err := uc.ServiceRepo.CreateVersion(ctx, serviceVersionEntity)
	if err != nil {
		return nil, appError.NewNotFoundError(err.Error())
	}
	return createdServiceVersion, nil
}

func (uc *ServiceCatalogUsecase) GetAllServiceVersionsById(ctx context.Context, serviceId int) ([]entity.ServiceVersion, error) {
	serviceVersions, err := uc.ServiceRepo.FindAllVersionsById(ctx, serviceId)
	if err != nil {
		return nil, appError.NewNotFoundError(err.Error())
	}
	return serviceVersions, nil
}
