package postgres

import (
	"efs-workforce/internal/domain"
	"efs-workforce/internal/ports/repositories"
	"time"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// TimeOffRepository implements the time off repository interface using PostgreSQL
type TimeOffRepository struct {
	db *gorm.DB
}

// NewTimeOffRepository creates a new PostgreSQL time off repository
func NewTimeOffRepository(db *gorm.DB) repositories.TimeOffRepository {
	return &TimeOffRepository{db: db}
}

// Create creates a new time off record
func (r *TimeOffRepository) Create(timeOff *domain.TimeOff) error {
	return r.db.Create(timeOff).Error
}

// FindByID finds a time off record by ID
func (r *TimeOffRepository) FindByID(id uuid.UUID) (*domain.TimeOff, error) {
	var timeOff domain.TimeOff
	if err := r.db.Preload("User").Preload("Creator").First(&timeOff, "id = ?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &timeOff, nil
}

// FindByUserID finds time off records by user ID
func (r *TimeOffRepository) FindByUserID(userID uuid.UUID) ([]*domain.TimeOff, error) {
	var timeOff []*domain.TimeOff
	if err := r.db.Preload("User").Where("user_id = ?", userID).Order("start_date DESC").Find(&timeOff).Error; err != nil {
		return nil, err
	}
	return timeOff, nil
}

// FindByStatus finds time off records by status
func (r *TimeOffRepository) FindByStatus(status domain.TimeOffStatus) ([]*domain.TimeOff, error) {
	var timeOff []*domain.TimeOff
	if err := r.db.Preload("User").Where("status = ?", status).Order("start_date DESC").Find(&timeOff).Error; err != nil {
		return nil, err
	}
	return timeOff, nil
}

// FindByDateRange finds time off records by date range
func (r *TimeOffRepository) FindByDateRange(startDate, endDate time.Time) ([]*domain.TimeOff, error) {
	var timeOff []*domain.TimeOff
	if err := r.db.Preload("User").Where("start_date <= ? AND end_date >= ?", endDate, startDate).Order("start_date DESC").Find(&timeOff).Error; err != nil {
		return nil, err
	}
	return timeOff, nil
}

// FindAll finds all time off records
func (r *TimeOffRepository) FindAll() ([]*domain.TimeOff, error) {
	var timeOff []*domain.TimeOff
	if err := r.db.Preload("User").Preload("Creator").Order("start_date DESC").Find(&timeOff).Error; err != nil {
		return nil, err
	}
	return timeOff, nil
}

// Update updates an existing time off record
func (r *TimeOffRepository) Update(timeOff *domain.TimeOff) error {
	return r.db.Save(timeOff).Error
}

// Delete deletes a time off record by ID
func (r *TimeOffRepository) Delete(id uuid.UUID) error {
	return r.db.Delete(&domain.TimeOff{}, "id = ?", id).Error
}

