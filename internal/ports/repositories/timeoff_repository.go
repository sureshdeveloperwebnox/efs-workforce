package repositories

import (
	"efs-workforce/internal/domain"
	"time"
	"github.com/google/uuid"
)

// TimeOffRepository defines the interface for time off data operations
type TimeOffRepository interface {
	Create(timeOff *domain.TimeOff) error
	FindByID(id uuid.UUID) (*domain.TimeOff, error)
	FindByUserID(userID uuid.UUID) ([]*domain.TimeOff, error)
	FindByStatus(status domain.TimeOffStatus) ([]*domain.TimeOff, error)
	FindByDateRange(startDate, endDate time.Time) ([]*domain.TimeOff, error)
	FindAll() ([]*domain.TimeOff, error)
	Update(timeOff *domain.TimeOff) error
	Delete(id uuid.UUID) error
}

