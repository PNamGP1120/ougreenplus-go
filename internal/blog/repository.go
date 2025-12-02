package blog

import (
	"errors"

	"gorm.io/gorm"

	"github.com/PNamGP1120/ougreenplus-go/internal/database"
)

type Repository interface {
	List(page, size int) ([]Blog, int64, error)
	GetByID(id uint) (*Blog, error)
	Create(b *Blog) error
	Update(b *Blog) error
	Delete(id uint) error
}

type GormRepository struct {
	db *gorm.DB
}

func NewRepository() Repository {
	return &GormRepository{db: database.DB}
}

func (r *GormRepository) List(page, size int) ([]Blog, int64, error) {
	var items []Blog
	var total int64

	query := r.db.Model(&Blog{})

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if page <= 0 {
		page = 1
	}
	if size <= 0 || size > 100 {
		size = 10
	}
	offset := (page - 1) * size

	if err := query.
		Order("created_at DESC").
		Limit(size).
		Offset(offset).
		Find(&items).Error; err != nil {
		return nil, 0, err
	}

	return items, total, nil
}

func (r *GormRepository) GetByID(id uint) (*Blog, error) {
	var b Blog
	if err := r.db.First(&b, id).Error; err != nil {
		return nil, err
	}
	return &b, nil
}

func (r *GormRepository) Create(b *Blog) error {
	return r.db.Create(b).Error
}

func (r *GormRepository) Update(b *Blog) error {
	if b.ID == 0 {
		return errors.New("missing blog ID")
	}
	return r.db.Save(b).Error
}

func (r *GormRepository) Delete(id uint) error {
	return r.db.Delete(&Blog{}, id).Error
}
