package postgres

import (
	"efs-workforce/internal/domain"
	"efs-workforce/internal/ports/repositories"
	"time"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// AttendanceRepository implements the attendance repository interface using PostgreSQL
type AttendanceRepository struct {
	db *gorm.DB
}

// NewAttendanceRepository creates a new PostgreSQL attendance repository
func NewAttendanceRepository(db *gorm.DB) repositories.AttendanceRepository {
	return &AttendanceRepository{db: db}
}

// Create creates a new attendance record
func (r *AttendanceRepository) Create(attendance *domain.Attendance) error {
	return r.db.Create(attendance).Error
}

// FindByID finds an attendance record by ID
func (r *AttendanceRepository) FindByID(id uuid.UUID) (*domain.Attendance, error) {
	var attendance domain.Attendance
	if err := r.db.Preload("User").Preload("Creator").First(&attendance, "id = ?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &attendance, nil
}

// FindByUserID finds attendance records by user ID
func (r *AttendanceRepository) FindByUserID(userID uuid.UUID) ([]*domain.Attendance, error) {
	var attendance []*domain.Attendance
	if err := r.db.Preload("User").Where("user_id = ?", userID).Order("created_at DESC").Find(&attendance).Error; err != nil {
		return nil, err
	}
	return attendance, nil
}

// FindByUserIDAndDate finds an attendance record by user ID and date
func (r *AttendanceRepository) FindByUserIDAndDate(userID uuid.UUID, date time.Time) (*domain.Attendance, error) {
	var attendance domain.Attendance
	startOfDay := time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, date.Location())
	endOfDay := startOfDay.Add(24 * time.Hour)
	
	if err := r.db.Preload("User").Where("user_id = ? AND created_at >= ? AND created_at < ?", userID, startOfDay, endOfDay).First(&attendance).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &attendance, nil
}

// FindByDateRange finds attendance records by user ID and date range
func (r *AttendanceRepository) FindByDateRange(userID uuid.UUID, startDate, endDate time.Time) ([]*domain.Attendance, error) {
	var attendance []*domain.Attendance
	if err := r.db.Preload("User").Where("user_id = ? AND created_at >= ? AND created_at <= ?", userID, startDate, endDate).Order("created_at DESC").Find(&attendance).Error; err != nil {
		return nil, err
	}
	return attendance, nil
}

// FindAll finds all attendance records
func (r *AttendanceRepository) FindAll() ([]*domain.Attendance, error) {
	var attendance []*domain.Attendance
	if err := r.db.Preload("User").Preload("Creator").Order("created_at DESC").Find(&attendance).Error; err != nil {
		return nil, err
	}
	return attendance, nil
}

// Update updates an existing attendance record
func (r *AttendanceRepository) Update(attendance *domain.Attendance) error {
	return r.db.Save(attendance).Error
}

// Delete deletes an attendance record by ID
func (r *AttendanceRepository) Delete(id uuid.UUID) error {
	return r.db.Delete(&domain.Attendance{}, "id = ?", id).Error
}

