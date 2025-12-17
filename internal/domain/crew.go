package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Crew represents a crew entity in the workforce domain
type Crew struct {
	ID        uuid.UUID  `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	CrewName  string     `gorm:"type:varchar(100);not null" json:"crew_name"`
	CreatedBy *uuid.UUID `gorm:"type:uuid;index" json:"created_by"`
	Creator   *User      `gorm:"foreignKey:CreatedBy" json:"creator,omitempty"`
	Members   []CrewMember `gorm:"foreignKey:CrewID" json:"members,omitempty"`
	CreatedAt time.Time  `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time  `gorm:"autoUpdateTime" json:"updated_at"`
}

// BeforeCreate hook to generate UUID if not set
func (c *Crew) BeforeCreate(tx *gorm.DB) error {
	if c.ID == uuid.Nil {
		c.ID = uuid.New()
	}
	return nil
}

// TableName specifies the table name for GORM
func (Crew) TableName() string {
	return "crews"
}

// CrewMember represents a crew member relationship
type CrewMember struct {
	ID         uuid.UUID  `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	CrewID     uuid.UUID  `gorm:"type:uuid;index;not null" json:"crew_id"`
	Crew       *Crew       `gorm:"foreignKey:CrewID" json:"crew,omitempty"`
	UserID     uuid.UUID  `gorm:"type:uuid;index;not null" json:"user_id"`
	User       *User       `gorm:"foreignKey:UserID" json:"user,omitempty"`
	AssignedAt time.Time  `gorm:"autoCreateTime" json:"assigned_at"`
}

// BeforeCreate hook to generate UUID if not set
func (cm *CrewMember) BeforeCreate(tx *gorm.DB) error {
	if cm.ID == uuid.Nil {
		cm.ID = uuid.New()
	}
	return nil
}

// TableName specifies the table name for GORM
func (CrewMember) TableName() string {
	return "crew_members"
}

