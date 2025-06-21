package entity

import (
	"gorm.io/gorm"
)

type Role struct {
	gorm.Model
	Name string `json:"name"`
}

type Permission struct {
	gorm.Model
	Name string `json:"name"`
}

type RolePermission struct {
	gorm.Model
	RoleID       uint `json:"role_id"`
	PermissionID uint `json:"permission_id"`
}