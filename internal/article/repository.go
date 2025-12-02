package article

import (
	"errors"

	"gorm.io/gorm"

	"github.com/PNamGP1120/ougreenplus-go/internal/database"
)

type Repository interface {
	List(page, size int, categoryID uint, status Status) ([]Article, int64, error)
	GetByID(id uint) (*Article, error)
	Create(a *Article) error
	Update(a *Article) error
	Delete(id uint) error
	ListRelated(baseID, categoryID uint, limit int) ([]Article, error)
}

type GormRepository struct {
	db *gorm.DB
}

func NewRepository() Repository {
	return &GormRepository{db: database.DB}
}

func (r *GormRepository) List(page, size int, categoryID uint, status Status) ([]Article, int64, error) {
	var items []Article
	var total int64

	query := r.db.Model(&Article{})

	if categoryID != 0 {
		query = query.Where("category_id = ?", categoryID)
	}
	if status != "" {
		query = query.Where("status = ?", status)
	}

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
		Order("published_at DESC NULLS LAST, created_at DESC").
		Limit(size).
		Offset(offset).
		Find(&items).Error; err != nil {
		return nil, 0, err
	}

	return items, total, nil
}

func (r *GormRepository) GetByID(id uint) (*Article, error) {
	var a Article
	if err := r.db.First(&a, id).Error; err != nil {
		return nil, err
	}
	return &a, nil
}

func (r *GormRepository) Create(a *Article) error {
	return r.db.Create(a).Error
}

func (r *GormRepository) Update(a *Article) error {
	if a.ID == 0 {
		return errors.New("missing article ID")
	}
	return r.db.Save(a).Error
}

func (r *GormRepository) Delete(id uint) error {
	return r.db.Delete(&Article{}, id).Error
}

func (r *GormRepository) ListRelated(baseID, categoryID uint, limit int) ([]Article, error) {
	var items []Article
	if limit <= 0 || limit > 20 {
		limit = 5
	}
	err := r.db.
		Where("category_id = ? AND id <> ?", categoryID, baseID).
		Order("published_at DESC NULLS LAST, created_at DESC").
		Limit(limit).
		Find(&items).Error
	return items, err
}
