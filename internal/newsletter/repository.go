package newsletter

import "github.com/PNamGP1120/ougreenplus-go/internal/database"

type Repository interface {
	Subscribe(s *Subscriber) error
	Unsubscribe(email string) error
	All() ([]Subscriber, error)
}

type GormRepository struct{}

func NewRepository() Repository {
	return &GormRepository{}
}

func (r *GormRepository) Subscribe(s *Subscriber) error {
	return database.DB.Create(s).Error
}

func (r *GormRepository) Unsubscribe(email string) error {
	return database.DB.Model(&Subscriber{}).Where("email = ?", email).Update("is_active", false).Error
}

func (r *GormRepository) All() ([]Subscriber, error) {
	var list []Subscriber
	err := database.DB.Order("created_at DESC").Find(&list).Error
	return list, err
}
