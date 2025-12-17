package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// LeaveType represents leave type enum
type LeaveType string

const (
	LeaveTypeSickLeave   LeaveType = "Sick Leave"
	LeaveTypeCasualLeave LeaveType = "Casual Leave"
	LeaveTypePaidLeave   LeaveType = "Paid Leave"
	LeaveTypeUnpaidLeave LeaveType = "Unpaid Leave"
)

// TimeOffStatus represents time off status enum
type TimeOffStatus string

const (
	TimeOffStatusPending  TimeOffStatus = "Pending"
	TimeOffStatusApproved TimeOffStatus = "Approved"
	TimeOffStatusRejected TimeOffStatus = "Rejected"
)

// TimeOff represents a time off/leave entity in the workforce domain
type TimeOff struct {
	ID        uuid.UUID    `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	UserID    uuid.UUID    `gorm:"type:uuid;index;not null" json:"user_id"`
	User      *User        `gorm:"foreignKey:UserID" json:"user,omitempty"`
	LeaveType LeaveType    `gorm:"type:varchar(50);not null" json:"leave_type"`
	StartDate time.Time    `gorm:"type:date;not null" json:"start_date"`
	EndDate   time.Time    `gorm:"type:date;not null" json:"end_date"`
	Reason    string       `gorm:"type:text" json:"reason"`
	Status    TimeOffStatus `gorm:"type:varchar(50);default:'Pending'" json:"status"`
	CreatedBy *uuid.UUID   `gorm:"type:uuid;index" json:"created_by"`
	Creator   *User        `gorm:"foreignKey:CreatedBy" json:"creator,omitempty"`
	CreatedAt time.Time    `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time    `gorm:"autoUpdateTime" json:"updated_at"`
}

// BeforeCreate hook to generate UUID if not set
func (t *TimeOff) BeforeCreate(tx *gorm.DB) error {
	if t.ID == uuid.Nil {
		t.ID = uuid.New()
	}
	return nil
}

// IsPending checks if the time off request is pending
func (t *TimeOff) IsPending() bool {
	return t.Status == TimeOffStatusPending
}

// TableName specifies the table name for GORM
func (TimeOff) TableName() string {
	return "time_off"
}

