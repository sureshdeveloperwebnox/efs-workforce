package postgres

import (
	"efs-workforce/internal/domain"
	"efs-workforce/internal/ports/repositories"
	"time"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// TripRepository implements the trip repository interface using PostgreSQL
type TripRepository struct {
	db *gorm.DB
}

// NewTripRepository creates a new PostgreSQL trip repository
func NewTripRepository(db *gorm.DB) repositories.TripRepository {
	return &TripRepository{db: db}
}

// Create creates a new trip
func (r *TripRepository) Create(trip *domain.Trip) error {
	return r.db.Create(trip).Error
}

// FindByID finds a trip by ID
func (r *TripRepository) FindByID(id uuid.UUID) (*domain.Trip, error) {
	var trip domain.Trip
	if err := r.db.Preload("User").Preload("Creator").First(&trip, "id = ?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &trip, nil
}

// FindByUserID finds trips by user ID
func (r *TripRepository) FindByUserID(userID uuid.UUID) ([]*domain.Trip, error) {
	var trips []*domain.Trip
	if err := r.db.Preload("User").Where("user_id = ?", userID).Order("start_time DESC").Find(&trips).Error; err != nil {
		return nil, err
	}
	return trips, nil
}

// FindByDateRange finds trips by user ID and date range
func (r *TripRepository) FindByDateRange(userID uuid.UUID, startDate, endDate time.Time) ([]*domain.Trip, error) {
	var trips []*domain.Trip
	if err := r.db.Preload("User").Where("user_id = ? AND start_time >= ? AND start_time <= ?", userID, startDate, endDate).Order("start_time DESC").Find(&trips).Error; err != nil {
		return nil, err
	}
	return trips, nil
}

// FindAll finds all trips
func (r *TripRepository) FindAll() ([]*domain.Trip, error) {
	var trips []*domain.Trip
	if err := r.db.Preload("User").Preload("Creator").Order("start_time DESC").Find(&trips).Error; err != nil {
		return nil, err
	}
	return trips, nil
}

// Update updates an existing trip
func (r *TripRepository) Update(trip *domain.Trip) error {
	return r.db.Save(trip).Error
}

// Delete deletes a trip by ID
func (r *TripRepository) Delete(id uuid.UUID) error {
	return r.db.Delete(&domain.Trip{}, "id = ?", id).Error
}

