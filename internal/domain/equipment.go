package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// EquipmentStatus represents equipment status enum
type EquipmentStatus string

const (
	EquipmentStatusActive          EquipmentStatus = "Active"
	EquipmentStatusInactive        EquipmentStatus = "Inactive"
	EquipmentStatusUnderMaintenance EquipmentStatus = "Under Maintenance"
)

// Equipment represents an equipment entity in the workforce domain
type Equipment struct {
	ID            uuid.UUID       `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	Name          string          `gorm:"type:varchar(100);not null" json:"name"`
	SerialNumber  string          `gorm:"type:varchar(50);uniqueIndex" json:"serial_number"`
	AssignedToUser *uuid.UUID     `gorm:"type:uuid;index" json:"assigned_to_user"`
	User          *User           `gorm:"foreignKey:AssignedToUser" json:"user,omitempty"`
	Status        EquipmentStatus `gorm:"type:varchar(50);default:'Active'" json:"status"`
	CreatedBy     *uuid.UUID      `gorm:"type:uuid;index" json:"created_by"`
	Creator       *User           `gorm:"foreignKey:CreatedBy" json:"creator,omitempty"`
	CreatedAt     time.Time       `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt     time.Time       `gorm:"autoUpdateTime" json:"updated_at"`
}

// BeforeCreate hook to generate UUID if not set
func (e *Equipment) BeforeCreate(tx *gorm.DB) error {
	if e.ID == uuid.Nil {
		e.ID = uuid.New()
	}
	return nil
}

// IsActive checks if the equipment is active
func (e *Equipment) IsActive() bool {
	return e.Status == EquipmentStatusActive
}

// TableName specifies the table name for GORM
func (Equipment) TableName() string {
	return "equipment"
}

