package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Role represents a role entity in the workforce domain
type Role struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	RoleName    string    `gorm:"type:varchar(50);uniqueIndex;not null" json:"role_name"`
	Description string    `gorm:"type:varchar(255)" json:"description"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

// BeforeCreate hook to generate UUID if not set
func (r *Role) BeforeCreate(tx *gorm.DB) error {
	if r.ID == uuid.Nil {
		r.ID = uuid.New()
	}
	return nil
}

// TableName specifies the table name for GORM
func (Role) TableName() string {
	return "roles"
}

