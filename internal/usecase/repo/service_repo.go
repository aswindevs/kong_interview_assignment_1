package repo

import (
	"context"

	"github.com/aswindevs/kong_interview-assignment_1/internal/entity"
	appError "github.com/aswindevs/kong_interview-assignment_1/internal/errors"
	"gorm.io/gorm"
)

type ServiceRepo interface {
	FindAll(ctx context.Context, search string, sortBy string, OrderBy string, page int, pageSize int) ([]entity.Service, int, error)
	FindById(ctx context.Context, id int) (*entity.Service, error)
	CreateVersion(ctx context.Context, serviceVersion entity.ServiceVersion) (*entity.ServiceVersion, error)
	FindAllVersionsById(ctx context.Context, serviceId int) ([]entity.ServiceVersion, error)
	Create(ctx context.Context, service entity.Service) (*entity.Service, error)
	Update(ctx context.Context, id int, service entity.Service) (*entity.Service, error)
	Delete(ctx context.Context, id int) error
}

type serviceRepoImpl struct {
	DB *gorm.DB
}

func NewServiceRepoImpl(DB *gorm.DB) *serviceRepoImpl {
	return &serviceRepoImpl{DB}
}

func (r *serviceRepoImpl) FindAll(ctx context.Context, search string, sortBy string, orderBy string, page int, pageSize int) ([]entity.Service, int, error) {
	var services []entity.Service
	result := r.DB.Where("name LIKE ?", "%"+search+"%").Order(sortBy + " " + orderBy).Offset((page - 1) * pageSize).Limit(pageSize).Find(&services)
	if result.Error != nil {
		return nil, 0, result.Error
	}
	var total int64
	r.DB.Model(&entity.Service{}).Where("name LIKE ?", "%"+search+"%").Count(&total)

	return services, int(total), nil
}

func (r *serviceRepoImpl) FindById(ctx context.Context, id int) (*entity.Service, error) {
	var service entity.Service
	result := r.DB.First(&service, id)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, appError.NewNotFoundError("service not found")
		}
		return nil, result.Error
	}
	return &service, nil
}

func (r *serviceRepoImpl) Create(ctx context.Context, service entity.Service) (*entity.Service, error) {
	result := r.DB.Create(&service)
	if result.Error != nil {
		return nil, result.Error
	}
	return &service, nil
}
func (r *serviceRepoImpl) CreateVersion(ctx context.Context, serviceVersion entity.ServiceVersion) (*entity.ServiceVersion, error) {
	result := r.DB.Create(&serviceVersion)
	if result.Error != nil {
		return nil, result.Error
	}
	return &serviceVersion, nil
}

func (r *serviceRepoImpl) FindAllVersionsById(ctx context.Context, serviceId int) ([]entity.ServiceVersion, error) {
	var serviceVersions []entity.ServiceVersion
	result := r.DB.Where("service_id = ?", serviceId).Find(&serviceVersions)
	if result.Error != nil {
		return nil, result.Error
	}
	return serviceVersions, nil
}

func (r *serviceRepoImpl) Update(ctx context.Context, id int, service entity.Service) (*entity.Service, error) {
	result := r.DB.Model(&entity.Service{}).Where("id = ?", id).Updates(service)
	if result.Error != nil {
		return nil, result.Error
	}
	return &service, nil
}

func (r *serviceRepoImpl) Delete(ctx context.Context, id int) error {
	result := r.DB.Delete(&entity.Service{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
