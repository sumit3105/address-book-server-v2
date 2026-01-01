package repositories

import (
	"address-book-server-v2/internal/models"
	"errors"

	"gorm.io/gorm"
)

type IUserRepository interface {
	Create(user *models.User) error
	FindByEmail(email string) (*models.User, error)
	ExistsByEmail(email string) (bool, error)
	ExistsByID(userID uint) (bool, error)
}

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(user *models.User) error {
	return r.db.Create(user).Error
}

func (r *UserRepository) FindByEmail(email string) (*models.User, error) {
	var user models.User
	err := r.db.Where("email=?", email).Find(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) ExistsByEmail(email string) (bool, error) {
	var count int64
	err := r.db.Model(&models.User{}).Where("email=?", email).Count(&count).Error
	return count > 0, err
}

func (r *UserRepository) ExistsByID(userID uint) (bool, error) {
	var count int64
	err := r.db.Model(&models.User{}).Where("id=?", userID).Count(&count).Error
	return count > 0, err
}

func IsNotFound(err error) bool {
	return errors.Is(err, gorm.ErrRecordNotFound)
}
