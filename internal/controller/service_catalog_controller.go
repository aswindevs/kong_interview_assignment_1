package controller

import (
	"net/http"
	"strconv"

	"github.com/aswindevs/kong_interview-assignment_1/config"
	"github.com/aswindevs/kong_interview-assignment_1/internal/dto"
	apperrors "github.com/aswindevs/kong_interview-assignment_1/internal/errors"
	"github.com/aswindevs/kong_interview-assignment_1/internal/usecase"
	"github.com/aswindevs/kong_interview-assignment_1/pkg/logger"

	"github.com/gin-gonic/gin"
)

const (
	defaultPage     = 1
	defaultPageSize = 10
	defaultSortBy   = "id"
	defaultOrderBy  = "asc"
)

type ServiceCatalogController struct {
	serviceCatalogUsecase *usecase.ServiceCatalogUsecase
	logger                logger.Interface
	cfg                   *config.Auth
}

func NewServiceCatalogController(
	router *gin.RouterGroup,
	serviceCatalogUsecase *usecase.ServiceCatalogUsecase,
	logger logger.Interface,
	cfg *config.Auth,
) {
	controller := &ServiceCatalogController{
		serviceCatalogUsecase: serviceCatalogUsecase,
		logger:                logger,
		cfg:                   cfg,
	}

	router.GET("", controller.GetAllServices)
	router.GET("/:id", controller.GetServiceById)
	router.GET("/:id/versions", controller.GetAllServiceVersionsById)
	router.POST("", controller.CreateService)
	router.POST("/:id/versions", controller.CreateServiceVersion)
}

func (c *ServiceCatalogController) GetAllServices(ctx *gin.Context) {
	search := ctx.Query("search")
	sortBy := ctx.Query("sortBy")
	if sortBy == "" {
		sortBy = defaultSortBy
	}
	orderBy := ctx.Query("orderBy")
	if orderBy == "" {
		orderBy = defaultOrderBy
	}
	page, err := strconv.Atoi(ctx.Query("page"))
	if err != nil || page < 1 {
		page = defaultPage
	}
	pageSize, err := strconv.Atoi(ctx.Query("pageSize"))
	if err != nil || pageSize <= 0 {
		pageSize = defaultPageSize
	}

	services, total, err := c.serviceCatalogUsecase.GetAllServices(ctx, search, sortBy, orderBy, page, pageSize)
	if err != nil {
		HandleError(ctx, c.logger, err)
		return
	}
	response := make([]dto.ServiceCatalogResponse, len(services))
	for i, service := range services {
		response[i] = dto.FromServiceEntity(service)
	}
	ctx.JSON(http.StatusOK, gin.H{"services": response, "total": total})
}

func (c *ServiceCatalogController) GetServiceById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		HandleError(ctx, c.logger, apperrors.NewBadRequestError("invalid service id"))
		return
	}

	service, err := c.serviceCatalogUsecase.GetServiceById(ctx, id)
	if err != nil {
		HandleError(ctx, c.logger, err)
		return
	}
	response := dto.FromServiceEntity(*service)
	ctx.JSON(http.StatusOK, gin.H{"service": response})
}

func (c *ServiceCatalogController) CreateService(ctx *gin.Context) {
	var service dto.ServiceCatalogRequest
	if err := ctx.ShouldBindJSON(&service); err != nil {
		HandleError(ctx, c.logger, apperrors.NewBadRequestError("invalid request body"))
		return
	}
	createdService, err := c.serviceCatalogUsecase.CreateService(ctx, service)
	if err != nil {
		HandleError(ctx, c.logger, err)
		return
	}
	response := dto.FromServiceEntity(*createdService)
	ctx.JSON(http.StatusOK, gin.H{"service": response})
}
func (c *ServiceCatalogController) CreateServiceVersion(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		HandleError(ctx, c.logger, apperrors.NewBadRequestError("invalid service id"))
		return
	}
	var serviceVersion dto.ServiceVersionRequest
	if err := ctx.ShouldBindJSON(&serviceVersion); err != nil {
		HandleError(ctx, c.logger, apperrors.NewBadRequestError("invalid request body"))
		return
	}
	createdServiceVersion, err := c.serviceCatalogUsecase.CreateServiceVersion(ctx, id, serviceVersion)
	if err != nil {
		HandleError(ctx, c.logger, err)
		return
	}
	response := dto.FromServiceVersionEntity(*createdServiceVersion)
	ctx.JSON(http.StatusOK, gin.H{"serviceVersion": response})
}

func (c *ServiceCatalogController) GetAllServiceVersionsById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		HandleError(ctx, c.logger, apperrors.NewBadRequestError("invalid service id"))
		return
	}
	serviceVersions, err := c.serviceCatalogUsecase.GetAllServiceVersionsById(ctx, id)
	if err != nil {
		HandleError(ctx, c.logger, err)
		return
	}
	response := make([]dto.ServiceVersionResponse, len(serviceVersions))
	for i, serviceVersion := range serviceVersions {
		response[i] = dto.FromServiceVersionEntity(serviceVersion)
	}
	ctx.JSON(http.StatusOK, gin.H{"serviceVersions": response})
}
