package repositories

import (
	"address-book-server-v2/internal/models"

	"gorm.io/gorm"
)

type IAddressRepository interface {
	Create(address *models.Address) error
	FindByUser(userID uint64) ([]models.Address, error)
	FindByID(id uint64, userID uint64) (*models.Address, error)
	Update(address *models.Address) error
	Delete(address *models.Address) error
	FindAllForExport(userID uint64) ([]models.Address, error)
	FindFiltered(requestQuery *models.ListAddressQuery) ([]models.Address, int64, error)
}

type AddressRepository struct {
	db *gorm.DB
}

func NewAddressRepository(db *gorm.DB) *AddressRepository {
	return &AddressRepository{db}
}

func (r *AddressRepository) Create(address *models.Address) error {
	return r.db.Create(address).Error
}

func (r *AddressRepository) FindAllByUser(userID uint64) ([]models.Address, error) {
	var addresses []models.Address
	err := r.db.Where("user_id = ?", userID).Find(&addresses).Error
	return addresses, err
}

func (r *AddressRepository) FindByID(id, userID uint64) (*models.Address, error) {
	var address models.Address
	err := r.db.Where("id=? AND user_id=?", id, userID).First(&address).Error
	if err != nil {
		return nil, err
	}
	return &address, nil
}

func (r *AddressRepository) Update(address *models.Address) error {
	return r.db.Save(address).Error
}

func (r *AddressRepository) Delete(address *models.Address) error {
	return r.db.Delete(address).Error // soft delete
}

func (r *AddressRepository) FindAllForExport(fields []string, userID uint64) ([]map[string]interface{}, error) {

	var results []map[string]interface{}

	err := r.db.
		Select(fields).
		Where("user_id = ?", userID).
		Find(&results).
		Error

	return results, err
}

// Query format:
// GET /addreses?page=1&limit=10&city=Ahmedabad
func (r *AddressRepository) FindFiltered(
	userID uint64,
	listAddressQuery models.ListAddressQuery,
) ([]models.Address, int64, error) {

	offset := (listAddressQuery.Page - 1) * listAddressQuery.Limit

	query := r.db.Model(&models.Address{}).Where("user_id = ?", userID)

	// SEARCH (across multiple fields)
	if listAddressQuery.Search != "" {
		like := "%" + listAddressQuery.Search + "%"
		query = query.Where(`
			first_name ILIKE ? OR 
			last_name ILIKE ? OR 
			email ILIKE ? OR
			phone ILIKE ? OR
			city ILIKE ? OR
			state ILIKE ? OR
			country ILIKE ?`,
			like, like, like, like, like, like, like,
		)
	}

	// FILTERS
	if listAddressQuery.City != "" {
		query = query.Where("city ILIKE ?", listAddressQuery.City)
	}
	if listAddressQuery.State != "" {
		query = query.Where("state ILIKE ?", listAddressQuery.State)
	}
	if listAddressQuery.Country != "" {
		query = query.Where("country ILIKE ?", listAddressQuery.Country)
	}

	// fmt.Println(query)

	var total int64
	query.Count(&total) // get total records

	// fmt.Println(total)

	// PAGINATION
	var addresses []models.Address
	err := query.
		Limit(listAddressQuery.Limit).
		Offset(offset).
		Order("created_at DESC").
		Find(&addresses).Error

	// fmt.Println("inside repo:",err)

	return addresses, total, err
}
