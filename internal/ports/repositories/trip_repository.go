package repositories

import (
	"efs-workforce/internal/domain"
	"time"
	"github.com/google/uuid"
)

// TripRepository defines the interface for trip data operations
type TripRepository interface {
	Create(trip *domain.Trip) error
	FindByID(id uuid.UUID) (*domain.Trip, error)
	FindByUserID(userID uuid.UUID) ([]*domain.Trip, error)
	FindByDateRange(userID uuid.UUID, startDate, endDate time.Time) ([]*domain.Trip, error)
	FindAll() ([]*domain.Trip, error)
	Update(trip *domain.Trip) error
	Delete(id uuid.UUID) error
}

