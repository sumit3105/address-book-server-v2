package controllers

import (
	"address-book-server-v2/internal/common/fault"
	"address-book-server-v2/internal/common/utils"
	"address-book-server-v2/internal/core/config"
	"address-book-server-v2/internal/models"
	"address-book-server-v2/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type IAuthController interface {
	Register(ctx *gin.Context)
	Login(ctx *gin.Context)
}

type AuthController struct {
	authService *services.AuthService
}

func NewAuthController(serverCfg *config.ServerConfig, db *gorm.DB) *AuthController {
	authService := services.NewAuthService(serverCfg, db)
	return &AuthController{authService}
}

func (c *AuthController) Register(ctx *gin.Context) {
	var req models.RegisterRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.Error(ctx, 400, fault.BadRequest(
			"Invalid request body",
			err,
		),
		)
		return
	}

	if err := utils.Validate.Struct(req); err != nil {
		utils.Error(ctx, 400, fault.NewValidationError(utils.FormatValidationErrors(err)))
		return
	}

	if err := c.authService.Register(req.Email, req.Password); err != nil {
		utils.Error(ctx, 400, err)
		return
	}

	utils.Success(ctx, http.StatusCreated, gin.H{"message": "user registered"})
}

func (c *AuthController) Login(ctx *gin.Context) {
	var req models.LoginRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.Error(ctx, http.StatusBadRequest, fault.BadRequest(
			"Invalid request body",
			err,
		))
		return
	}

	if err := utils.Validate.Struct(req); err != nil {
		utils.Error(ctx, 400, fault.NewValidationError(utils.FormatValidationErrors(err)))
		return
	}

	token, err := c.authService.Login(req.Email, req.Password)
	if err != nil {
		utils.Error(ctx, http.StatusUnauthorized, err)
		return
	}

	utils.Success(ctx, http.StatusOK, gin.H{
		"token": token,
	})
}
