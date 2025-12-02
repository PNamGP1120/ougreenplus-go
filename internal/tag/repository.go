package tag

import (
	"github.com/gosimple/slug"

	"github.com/PNamGP1120/ougreenplus-go/internal/database"
)

type Repository interface {
	List() ([]Tag, error)
	Create(t *Tag) error
	Update(t *Tag) error
	Delete(id uint) error
}

type GormRepository struct{}

func NewRepository() Repository {
	return &GormRepository{}
}

func (r *GormRepository) List() ([]Tag, error) {
	var items []Tag
	err := database.DB.Order("name ASC").Find(&items).Error
	return items, err
}

func (r *GormRepository) Create(t *Tag) error {
	t.Slug = slug.Make(t.Name)
	return database.DB.Create(t).Error
}

func (r *GormRepository) Update(t *Tag) error {
	t.Slug = slug.Make(t.Name)
	return database.DB.Save(t).Error
}

func (r *GormRepository) Delete(id uint) error {
	return database.DB.Delete(&Tag{}, id).Error
}
