package entity

import (
	"gorm.io/gorm"
)

type Service struct {
	gorm.Model
	Name        string           `json:"name"`
	Description string           `json:"description"`
	Versions    []ServiceVersion `json:"versions"`
}

type ServiceVersion struct {
	gorm.Model
	ServiceID uint   `json:"service_id"`
	Version   int    `json:"version"`
}
