package services

import (
	"address-book-server-v2/internal/common/fault"
	"address-book-server-v2/internal/common/utils"
	"address-book-server-v2/internal/core/config"
	"address-book-server-v2/internal/models"
	"address-book-server-v2/internal/repositories"

	"gorm.io/gorm"
)

type IAuthService interface {
	Register(email, password string) error
	Login(email, password string) (string, error)
}

type AuthService struct{
	userRepo *repositories.UserRepository
	serverCfg *config.ServerConfig
}

func NewAuthService(serverCfg *config.ServerConfig, db *gorm.DB) *AuthService{
	userRepo := repositories.NewUserRepository(db)
	return &AuthService{userRepo: userRepo, serverCfg: serverCfg}
}

func (s *AuthService) Register(email, password string) error {
	exist, err := s.userRepo.ExistsByEmail(email)
	if err != nil {
		return fault.Internal("database error", err)
	}
	if exist {
		return fault.BadRequest("email already registered", nil)
	}

	hashedPass, err := utils.HashPassword(password)
	if err != nil {
		return fault.Internal("internal error", err)
	}

	user := &models.User{
		Email: email,
		Password: hashedPass,
	}
	return s.userRepo.Create(user)
}

func (s *AuthService) Login(email, password string) (string, error) {
	user, err := s.userRepo.FindByEmail(email)
	if err != nil {
		return "", fault.Internal("database error", err)
	}

	if user.Email == ""{
		return "", fault.Unauthorized("invalid credentials: Please correct Email", nil)
	}

	if err := utils.ComparePassword(user.Password, password); err != nil {
		return "", fault.Unauthorized("invalid credentials: Please correct Password", nil)
	}

	token, err := utils.GenerateToken(s.serverCfg.JwtSecret, user.ID, user.Email)
	if err != nil {
		return "",fault.Internal("internal error", err)
	}
	return token, nil
}