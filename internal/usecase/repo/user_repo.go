package repo

import (
	"context"

	"github.com/aswindevs/kong_interview-assignment_1/internal/entity"

	"gorm.io/gorm"
)

type UserRepo interface {
	FindById(ctx context.Context, id int) (*entity.User, error)
	FindByEmail(ctx context.Context, email string) (*entity.User, error)
	FindPermissionsByUserId(ctx context.Context, userId int) ([]entity.Permission, error)
}

type UserRepoImpl struct {
	DB *gorm.DB
}

func NewUserRepoImpl(DB *gorm.DB) *UserRepoImpl {
	return &UserRepoImpl{DB}
}


func (r *UserRepoImpl) FindById(ctx context.Context, id int) (*entity.User, error) {
	var user entity.User
	result := r.DB.Preload("Roles").First(&user, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (r *UserRepoImpl) FindByEmail(ctx context.Context, email string) (*entity.User, error) {
	var user entity.User
	result := r.DB.Preload("Roles").First(&user, "email = ?", email)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (r *UserRepoImpl) FindPermissionsByUserId(ctx context.Context, userId int) ([]entity.Permission, error) {
	var permissions []entity.Permission
	result := r.DB.Table("permissions").
		Joins("JOIN role_permissions ON permissions.id = role_permissions.permission_id").
		Joins("JOIN user_roles ON role_permissions.role_id = user_roles.role_id").
		Where("user_roles.user_id = ?", userId).
		Find(&permissions)
	if result.Error != nil {
		return nil, result.Error
	}
	return permissions, nil
}