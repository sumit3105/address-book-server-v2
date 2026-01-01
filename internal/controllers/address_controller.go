package controllers

import (
	"address-book-server-v2/internal/common/fault"
	"address-book-server-v2/internal/common/utils"
	"address-book-server-v2/internal/core/config"
	"address-book-server-v2/internal/models"
	"address-book-server-v2/internal/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type IAddressController interface {
	Create(ctx *gin.Context)
	GetAll(ctx *gin.Context)
	GetByID(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
	Export(ctx *gin.Context)
	ExportAsync(ctx *gin.Context)
	ExportCustom(ctx *gin.Context)
	GetFiltered(ctx *gin.Context)
}

type AddressController struct {
	service *services.AddressService
}

func NewAddressController(serverCfg *config.ServerConfig, smtpCfg *config.SMTPConfig, db *gorm.DB) *AddressController {
	addressService := services.NewAddressService(db, serverCfg, smtpCfg)
	return &AddressController{addressService}
}

func (c *AddressController) Create(ctx *gin.Context) {
	userID := ctx.GetUint64("user_id")

	var address models.CreateAddressRequest
	if err := ctx.ShouldBindJSON(&address); err != nil {
		utils.Error(ctx, http.StatusBadRequest, fault.BadRequest(
			"Invalid request body",
			err,
		))
		return
	}

	if err := utils.Validate.Struct(address); err != nil {
		utils.Error(ctx, 400, fault.NewValidationError(utils.FormatValidationErrors(err)))
		return
	}

	if err := c.service.Create(userID, &address); err != nil {
		utils.Error(ctx, http.StatusInternalServerError, err)
		return
	}

	utils.Success(ctx, http.StatusCreated, gin.H{
		"Data": "address created",
	})
}

func (c *AddressController) GetAll(ctx *gin.Context) {
	userID := ctx.GetUint64("user_id")

	addresses, err := c.service.GetAll(userID)
	if err != nil {
		utils.Error(ctx, http.StatusInternalServerError, err)
		return
	}

	utils.Success(ctx, http.StatusCreated, gin.H{
		"Data": addresses,
	})
}

func (c *AddressController) GetByID(ctx *gin.Context) {
	userID := ctx.GetUint64("user_id")
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		utils.Error(ctx, http.StatusBadRequest, fault.BadRequest("invalid address id", err))
		return
	}

	address, err := c.service.GetByID(userID, id)
	if err != nil {
		utils.Error(ctx, http.StatusNotFound, err)
		return
	}

	utils.Success(ctx, http.StatusOK, gin.H{
		"data": address,
	})
}

func (c *AddressController) Update(ctx *gin.Context) {
	userID := ctx.GetUint64("user_id")
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		utils.Error(ctx, http.StatusBadRequest, fault.BadRequest("invalid address id", err))
		return
	}

	var address models.UpdateAddressRequest
	if err := ctx.ShouldBindJSON(&address); err != nil {
		utils.Error(ctx, http.StatusBadRequest, fault.BadRequest(
			"Invalid request body",
			err,
		))
		return
	}

	if err := utils.Validate.Struct(address); err != nil {
		utils.Error(ctx, 400, fault.NewValidationError(utils.FormatValidationErrors(err)))
		return
	}

	if err := c.service.Update(userID, id, &address); err != nil {
		utils.Error(ctx, http.StatusForbidden, err)
		return
	}

	utils.Success(ctx, http.StatusCreated, gin.H{"message": "address updated"})
}

func (c *AddressController) Delete(ctx *gin.Context) {
	userID := ctx.GetUint64("user_id")
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		utils.Error(ctx, http.StatusBadRequest, fault.BadRequest("invalid address id", err))
		return
	}

	if err := c.service.Delete(userID, id); err != nil {
		utils.Error(ctx, http.StatusForbidden, err)
		return
	}

	utils.Success(ctx, http.StatusOK, gin.H{"message": "address deleted"})
}

func (c *AddressController) ExportCustom(ctx *gin.Context) {
	userID := ctx.GetUint64("user_id")

	var req models.ExportAddressRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.Error(ctx, 400, fault.BadRequest(
			"Invalid request body",
			err,
		))
		return
	}
	if err := utils.Validate.Struct(req); err != nil {
		utils.Error(ctx, 400, fault.NewValidationError(utils.FormatValidationErrors(err)))
		return
	}

	c.service.ExportAddressesCustomAsync(userID, req.Fields, req.Email)

	utils.Success(ctx, 202, gin.H{
		"message": "Custom export started. CSV will be emailed shortly.",
	})
}

func (c *AddressController) GetFiltered(ctx *gin.Context) {
	userID := ctx.GetUint64("user_id")

	var query models.ListAddressQuery
	if err := ctx.ShouldBindQuery(&query); err != nil {
		utils.Error(ctx, 400, fault.BadRequest(
			"Invalid request body",
			err,
		))
		return
	}

	if query.Page <= 0 {
		query.Page = 1
	}

	if query.Limit <= 0 {
		query.Limit = 10
	}

	result, total, err := c.service.GetFilteredAddresses(userID, query)

	if err != nil {
		utils.Error(ctx, 500, fault.Internal("failed to fetch addresses", err))
		return
	}

	utils.Success(ctx, 200, gin.H{
		"result": result,
		"total":  total,
		"limit":  query.Limit,
		"page":   query.Page,
		"total_pages": (total + int64(query.Limit) - 1) / int64(query.Limit), 
	})
}
