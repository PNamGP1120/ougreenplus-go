package category

import (
	"errors"

	"gorm.io/gorm"

	"github.com/PNamGP1120/ougreenplus-go/internal/article"
	"github.com/PNamGP1120/ougreenplus-go/internal/database"
)

type Repository interface {
	List() ([]Category, error)
	Create(c *Category) error
	Update(c *Category) error
	Delete(id uint) error
	ListArticles(categoryID uint, page, size int) ([]article.Article, int64, error)
}

type GormRepository struct {
	db *gorm.DB
}

func NewRepository() Repository {
	return &GormRepository{db: database.DB}
}

func (r *GormRepository) List() ([]Category, error) {
	var items []Category
	if err := r.db.Order("name ASC").Find(&items).Error; err != nil {
		return nil, err
	}
	return items, nil
}

func (r *GormRepository) Create(c *Category) error {
	return r.db.Create(c).Error
}

func (r *GormRepository) Update(c *Category) error {
	if c.ID == 0 {
		return errors.New("missing category ID")
	}
	return r.db.Save(c).Error
}

func (r *GormRepository) Delete(id uint) error {
	return r.db.Delete(&Category{}, id).Error
}

func (r *GormRepository) ListArticles(categoryID uint, page, size int) ([]article.Article, int64, error) {
	var items []article.Article
	var total int64

	query := r.db.Model(&article.Article{}).Where("category_id = ?", categoryID)

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
