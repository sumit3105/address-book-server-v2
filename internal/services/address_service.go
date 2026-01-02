package services

import (
	"address-book-server-v2/internal/common/fault"
	logger "address-book-server-v2/internal/common/log"
	"address-book-server-v2/internal/common/utils"
	"address-book-server-v2/internal/core/config"
	"address-book-server-v2/internal/models"
	"address-book-server-v2/internal/repositories"
	"fmt"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type IAddressService interface {
	Create(userID uint64, address *models.Address) error
	GetAll(userID uint64) ([]models.AddressResponse, error)
	GetByID(userID, addressID uint64) (*models.AddressResponse, error)
	Update(userID uint64, id uint64, updated *models.Address) error
	Delete(userID uint64, id uint64) error
	ExportAddressesCustomAsync(userID uint64, fields []string, sendTo string)
	GetFilteredAddresses(userID uint64, listAddressQuery models.ListAddressQuery) ([]models.AddressResponse, int64, error)
}

type AddressService struct {
	addressRepo *repositories.AddressRepository
	serverCfg   *config.ServerConfig
	smtpCfg     *config.SMTPConfig
}

func NewAddressService(db *gorm.DB, serverCfg *config.ServerConfig, smtpCfg *config.SMTPConfig) *AddressService {
	addressRepository := repositories.NewAddressRepository(db)
	return &AddressService{addressRepository, serverCfg, smtpCfg}
}

func (s *AddressService) Create(userID uint64, req *models.CreateAddressRequest) error {

	address := &models.Address{
		UserID: userID,
		FirstName: req.FirstName,
		LastName: req.LastName,
		Email: req.Email,
		Phone: req.Phone,
		AddressLine1: req.AddressLine1,
		AddressLine2: req.AddressLine2,
		City: req.City,
		State: req.State,
		Country: req.Country,
		Pincode: req.Pincode,
	}

	address.UserID = userID
	err := s.addressRepo.Create(address)
	if err != nil {
		return fault.Internal("failed to create address", err)
	}
	return nil
}

func (s *AddressService) GetAll(userID uint64) ([]models.AddressResponse, error) {
	addresses, err := s.addressRepo.FindAllByUser(userID)
	if err != nil {
		return nil, fault.Internal("database error", err)
	}

	// Map Address to AddressResponse
	var result []models.AddressResponse
	for _, c := range addresses {
		result = append(result, models.AddressResponse{
			Id:           c.ID,
			FirstName:    c.FirstName,
			LastName:     c.LastName,
			Email:        c.Email,
			Phone:        c.Phone,
			AddressLine1: c.AddressLine1,
			AddressLine2: c.AddressLine2,
			City:         c.City,
			State:        c.State,
			Country:      c.Country,
			Pincode:      c.Pincode,
		})
	}

	return result, nil
}

func (s *AddressService) GetByID(userID, addressID uint64) (*models.AddressResponse, error) {
	address, err := s.addressRepo.FindByID(addressID, userID)
	if err != nil {
		return nil, fault.Internal("database error", err)
	}

	if address.ID == 0 {
		return nil, fault.NotFound("address not found", nil)
	}

	// ownership check
	if address.UserID != userID {
		return nil, fault.Forbidden("address not found for this user", nil)
	}

	// Map to DTO
	resp := &models.AddressResponse{
		Id:           address.ID,
		FirstName:    address.FirstName,
		LastName:     address.LastName,
		Email:        address.Email,
		Phone:        address.Phone,
		AddressLine1: address.AddressLine1,
		AddressLine2: address.AddressLine2,
		City:         address.City,
		State:        address.State,
		Country:      address.Country,
		Pincode:      address.Pincode,
	}

	return resp, nil
}

func (s *AddressService) Update(userID uint64, id uint64, req *models.UpdateAddressRequest) error {
	address, err := s.addressRepo.FindByID(id, userID)
	if err != nil {
		return fault.Internal("database error", err)
	}

	if address.ID == 0 {
		return fault.NotFound("address not found", nil)
	}

	// ownership check
	if address.UserID != userID {
		return fault.Forbidden("address not found for this user", nil)
	}

	if req.FirstName != nil {
		address.FirstName = *req.FirstName
	}
	if req.LastName != nil {
		address.LastName = *req.LastName
	}
	if req.Email != nil {
		address.Email = *req.Email
	}
	if req.Phone != nil {
		address.Phone = *req.Phone
	}
	if req.AddressLine1 != nil {
		address.AddressLine1 = *req.AddressLine1
	}
	if req.AddressLine2 != nil {
		address.AddressLine2 = *req.AddressLine2
	}
	if req.City != nil {
		address.City = *req.City
	}
	if req.State != nil {
		address.State = *req.State
	}
	if req.Country != nil {
		address.Country = *req.Country
	}
	if req.Pincode != nil {
		address.Pincode = *req.Pincode
	}

	err = s.addressRepo.Update(address)
	if err != nil {
		return fault.Internal("database error", err)
	}
	return nil
}

func (s *AddressService) Delete(userID uint64, id uint64) error {
	address, err := s.addressRepo.FindByID(id, userID)
	if err != nil {
		return fault.Internal("database error", err)
	}

	if address.ID == 0 {
		return fault.NotFound("address not found", nil)
	}

	if address.UserID != userID {
		return fault.Forbidden("address not found for this user", nil)
	}

	err = s.addressRepo.Delete(address)
	if err != nil {
		return fault.Internal("database error", err)
	}
	return nil
}

func (s *AddressService) ExportAddressesCustomAsync(
	userID uint64,
	fields []string,
	sendTo string,
) {

	go func() {
		defer func() {
			if r := recover(); r != nil {
				logger.Logger.Error("panic in async custom export", zap.Any("error", r))
			}
		}()

		logger.Logger.Info("custom address export started", zap.Uint64("user_id", userID))

		// 1. Fetch all addresses for user
		addresses, err := s.addressRepo.FindAllForExport(fields, userID)
		if err != nil {
			logger.Logger.Error("failed to fetch addresses", zap.Error(err))
			return
		}

		// 2. Generate CSV from filtered data
		filePath, fileName, err := utils.GenerateCustomAddressesCSV(userID, fields, addresses)
		if err != nil {
			logger.Logger.Error("failed to generate custom csv", zap.Error(err))
			return
		}

		// 3. Create download URL
		downloadURL := fmt.Sprintf(
			"%s/downloads/%s",
			s.serverCfg.AppURL,
			fileName,
		)

		// 4. Email with ATTACHMENT + LINK
		emailBody := fmt.Sprintf(
			"Attached is the custom address report you requested.\n\n"+
				"You can also download it using the link below:\n%s",
			downloadURL,
		)

		// 5. Email with attachment
		err = utils.SendEmailWithAttachment(
			s.smtpCfg.SMTPHost,
			s.smtpCfg.SMTPPort,
			s.smtpCfg.SMTPUser,
			s.smtpCfg.SMTPPass,
			sendTo,
			"Custom Address CSV Export",
			emailBody,
			filePath,
		)
		if err != nil {
			logger.Logger.Error("failed to send export email", zap.Error(err))
			return
		}

		logger.Logger.Info(
			"custom address export completed",
			zap.Uint64("user_id", userID),
			zap.String("file", filePath),
		)
	}()
}

func (s *AddressService) GetFilteredAddresses(
	userID uint64, listAddressQuery models.ListAddressQuery,
) ([]models.AddressResponse, int64, error) {

	addresses, total, err := s.addressRepo.FindFiltered(
		userID,
		listAddressQuery,
	)

	if err != nil {
		return nil, 0, fault.Internal("database error", err)
	}

	var responseData []models.AddressResponse
	for _, a := range addresses {
		responseData = append(responseData, models.AddressResponse{
			Id:           a.ID,
			FirstName:    a.FirstName,
			LastName:     a.LastName,
			Email:        a.Email,
			Phone:        a.Phone,
			AddressLine1: a.AddressLine1,
			AddressLine2: a.AddressLine2,
			City:         a.City,
			State:        a.State,
			Country:      a.Country,
			Pincode:      a.Pincode,
		})
	}

	return responseData, total, nil
}
