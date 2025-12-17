package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// AttendanceStatus represents attendance status enum
type AttendanceStatus string

const (
	AttendanceStatusPresent AttendanceStatus = "Present"
	AttendanceStatusAbsent  AttendanceStatus = "Absent"
	AttendanceStatusOnLeave AttendanceStatus = "On Leave"
)

// Attendance represents an attendance entity in the workforce domain
type Attendance struct {
	ID        uuid.UUID       `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	UserID    uuid.UUID       `gorm:"type:uuid;index;not null" json:"user_id"`
	User      *User           `gorm:"foreignKey:UserID" json:"user,omitempty"`
	CheckIn   *time.Time      `gorm:"type:timestamp" json:"check_in"`
	CheckOut  *time.Time      `gorm:"type:timestamp" json:"check_out"`
	Status    AttendanceStatus `gorm:"type:varchar(50);default:'Present'" json:"status"`
	CreatedBy *uuid.UUID      `gorm:"type:uuid;index" json:"created_by"`
	Creator   *User           `gorm:"foreignKey:CreatedBy" json:"creator,omitempty"`
	CreatedAt time.Time       `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time       `gorm:"autoUpdateTime" json:"updated_at"`
}

// BeforeCreate hook to generate UUID if not set
func (a *Attendance) BeforeCreate(tx *gorm.DB) error {
	if a.ID == uuid.Nil {
		a.ID = uuid.New()
	}
	return nil
}

// TableName specifies the table name for GORM
func (Attendance) TableName() string {
	return "attendance"
}

