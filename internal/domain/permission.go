package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Permission represents a permission entity in the workforce domain
type Permission struct {
	ID         uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	RoleID     uuid.UUID `gorm:"type:uuid;index;not null" json:"role_id"`
	Role       *Role     `gorm:"foreignKey:RoleID" json:"role,omitempty"`
	ModuleName string    `gorm:"type:varchar(50);not null" json:"module_name"`
	CanCreate  bool      `gorm:"default:false" json:"can_create"`
	CanRead    bool      `gorm:"default:false" json:"can_read"`
	CanUpdate  bool      `gorm:"default:false" json:"can_update"`
	CanDelete  bool      `gorm:"default:false" json:"can_delete"`
	CreatedAt  time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

// BeforeCreate hook to generate UUID if not set
func (p *Permission) BeforeCreate(tx *gorm.DB) error {
	if p.ID == uuid.Nil {
		p.ID = uuid.New()
	}
	return nil
}

// TableName specifies the table name for GORM
func (Permission) TableName() string {
	return "permissions"
}

