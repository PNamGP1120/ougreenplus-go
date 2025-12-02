package greennews

import (
	"github.com/PNamGP1120/ougreenplus-go/internal/database"
)

type Repository interface {
	List(month, year int) ([]Greennews, error)
	GetByID(id uint) (*Greennews, error)
	Create(g *Greennews) error
	Update(g *Greennews) error
	Delete(id uint) error
}

type GormRepository struct{}

func NewRepository() Repository {
	return &GormRepository{}
}

func (r *GormRepository) List(month, year int) ([]Greennews, error) {
	db := database.DB
	var list []Greennews

	if month > 0 {
		db = db.Where("month = ?", month)
	}
	if year > 0 {
		db = db.Where("year = ?", year)
	}

	if err := db.Order("year DESC, month DESC").Find(&list).Error; err != nil {
		return nil, err
	}

	return list, nil
}

func (r *GormRepository) GetByID(id uint) (*Greennews, error) {
	var g Greennews
	if err := database.DB.First(&g, id).Error; err != nil {
		return nil, err
	}
	return &g, nil
}

func (r *GormRepository) Create(g *Greennews) error {
	return database.DB.Create(g).Error
}

func (r *GormRepository) Update(g *Greennews) error {
	return database.DB.Save(g).Error
}

func (r *GormRepository) Delete(id uint) error {
	return database.DB.Delete(&Greennews{}, id).Error
}
