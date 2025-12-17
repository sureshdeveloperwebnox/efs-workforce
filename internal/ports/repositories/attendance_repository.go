package repositories

import (
	"efs-workforce/internal/domain"
	"time"
	"github.com/google/uuid"
)

// AttendanceRepository defines the interface for attendance data operations
type AttendanceRepository interface {
	Create(attendance *domain.Attendance) error
	FindByID(id uuid.UUID) (*domain.Attendance, error)
	FindByUserID(userID uuid.UUID) ([]*domain.Attendance, error)
	FindByUserIDAndDate(userID uuid.UUID, date time.Time) (*domain.Attendance, error)
	FindByDateRange(userID uuid.UUID, startDate, endDate time.Time) ([]*domain.Attendance, error)
	FindAll() ([]*domain.Attendance, error)
	Update(attendance *domain.Attendance) error
	Delete(id uuid.UUID) error
}

