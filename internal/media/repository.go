package media

import "github.com/PNamGP1120/ougreenplus-go/internal/database"

type Repository interface {
	Create(m *Media) error
	List() ([]Media, error)
	Get(id uint) (*Media, error)
	Delete(id uint) error
}

type GormRepository struct{}

func NewRepository() Repository {
	return &GormRepository{}
}

func (r *GormRepository) Create(m *Media) error {
	return database.DB.Create(m).Error
}

func (r *GormRepository) List() ([]Media, error) {
	var items []Media
	err := database.DB.Order("created_at DESC").Find(&items).Error
	return items, err
}

func (r *GormRepository) Get(id uint) (*Media, error) {
	var m Media
	err := database.DB.First(&m, id).Error
	return &m, err
}

func (r *GormRepository) Delete(id uint) error {
	return database.DB.Delete(&Media{}, id).Error
}
