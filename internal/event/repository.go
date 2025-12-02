package event

import (
	"github.com/PNamGP1120/ougreenplus-go/internal/database"
)

type Repository interface {
	List(status Status) ([]Event, error)
	GetByID(id uint) (*Event, error)
	Create(e *Event) error
	Update(e *Event) error
	Delete(id uint) error

	Register(r *Registration) error
	ListRegistrations(eventID uint) ([]Registration, error)
}

type GormRepository struct{}

func NewRepository() Repository {
	return &GormRepository{}
}

func (r *GormRepository) List(status Status) ([]Event, error) {
	db := database.DB
	var items []Event

	if status != "" {
		db = db.Where("status = ?", status)
	}

	if err := db.Order("start_date ASC").Find(&items).Error; err != nil {
		return nil, err
	}
	return items, nil
}

func (r *GormRepository) GetByID(id uint) (*Event, error) {
	var e Event
	if err := database.DB.First(&e, id).Error; err != nil {
		return nil, err
	}
	return &e, nil
}

func (r *GormRepository) Create(e *Event) error {
	return database.DB.Create(e).Error
}

func (r *GormRepository) Update(e *Event) error {
	return database.DB.Save(e).Error
}

func (r *GormRepository) Delete(id uint) error {
	return database.DB.Delete(&Event{}, id).Error
}

func (r *GormRepository) Register(reg *Registration) error {
	return database.DB.Create(reg).Error
}

func (r *GormRepository) ListRegistrations(eventID uint) ([]Registration, error) {
	var list []Registration
	err := database.DB.Where("event_id = ?", eventID).Find(&list).Error
	return list, err
}
