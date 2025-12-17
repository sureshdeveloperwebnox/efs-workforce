package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Trip represents a trip entity in the workforce domain
type Trip struct {
	ID          uuid.UUID  `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	UserID      uuid.UUID  `gorm:"type:uuid;index;not null" json:"user_id"`
	User        *User      `gorm:"foreignKey:UserID" json:"user,omitempty"`
	StartLocation string   `gorm:"type:varchar(255);not null" json:"start_location"`
	EndLocation   string   `gorm:"type:varchar(255);not null" json:"end_location"`
	StartTime     time.Time `gorm:"type:timestamp;not null" json:"start_time"`
	EndTime       *time.Time `gorm:"type:timestamp" json:"end_time"`
	Purpose       string    `gorm:"type:text" json:"purpose"`
	DistanceKm    *float64  `gorm:"type:decimal(8,2)" json:"distance_km"`
	CreatedBy     *uuid.UUID `gorm:"type:uuid;index" json:"created_by"`
	Creator       *User     `gorm:"foreignKey:CreatedBy" json:"creator,omitempty"`
	CreatedAt     time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt     time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

// BeforeCreate hook to generate UUID if not set
func (t *Trip) BeforeCreate(tx *gorm.DB) error {
	if t.ID == uuid.Nil {
		t.ID = uuid.New()
	}
	return nil
}

// TableName specifies the table name for GORM
func (Trip) TableName() string {
	return "trips"
}

