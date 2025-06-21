package dto

import (
	"strconv"

	"github.com/aswindevs/kong_interview-assignment_1/internal/entity"
)

type ServiceCatalogRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type ServiceCatalogResponse struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type ServiceVersionRequest struct {
	Version string `json:"version"`
}

type ServiceVersionResponse struct {
	ID      string `json:"id"`
	Version string `json:"version"`
}

func FromServiceEntity(service entity.Service) ServiceCatalogResponse {
	return ServiceCatalogResponse{
		ID:          strconv.FormatUint(uint64(service.ID), 10),
		Name:        service.Name,
		Description: service.Description,
	}
}

func FromServiceVersionEntity(serviceVersion entity.ServiceVersion) ServiceVersionResponse {
	return ServiceVersionResponse{
		ID:      strconv.FormatUint(uint64(serviceVersion.ID), 10),
		Version: strconv.Itoa(serviceVersion.Version),
	}
}
