package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// UserStatus represents user status enum
type UserStatus string

const (
	UserStatusActive   UserStatus = "Active"
	UserStatusInactive UserStatus = "Inactive"
)

// UserProfile represents user profile enum
type UserProfile string

const (
	UserProfileFieldAgent    UserProfile = "Field Agent"
	UserProfileManager       UserProfile = "Manager"
	UserProfileAdministrator UserProfile = "Administrator"
)

// User represents a user entity in the workforce domain
type User struct {
	ID         uuid.UUID   `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	FirstName  string      `gorm:"type:varchar(50);not null" json:"first_name"`
	LastName   string      `gorm:"type:varchar(50);not null" json:"last_name"`
	EmployeeID string      `gorm:"type:varchar(20);uniqueIndex;not null" json:"employee_id"`
	Email      string      `gorm:"type:varchar(100);uniqueIndex;not null" json:"email"`
	Phone      string      `gorm:"type:varchar(20)" json:"phone"`
	Status     UserStatus  `gorm:"type:varchar(20);default:'Active'" json:"status"`
	Profile    UserProfile `gorm:"type:varchar(50);default:'Field Agent'" json:"profile"`
	RoleID     *uuid.UUID  `gorm:"type:uuid;index" json:"role_id"`
	Role       *Role       `gorm:"foreignKey:RoleID" json:"role,omitempty"`
	CreatedBy  *uuid.UUID  `gorm:"type:uuid;index" json:"created_by"`
	Creator    *User       `gorm:"foreignKey:CreatedBy" json:"creator,omitempty"`
	CreatedAt  time.Time   `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt  time.Time   `gorm:"autoUpdateTime" json:"updated_at"`
}

// BeforeCreate hook to generate UUID if not set
func (u *User) BeforeCreate(tx *gorm.DB) error {
	if u.ID == uuid.Nil {
		u.ID = uuid.New()
	}
	return nil
}

// IsActive checks if the user is active
func (u *User) IsActive() bool {
	return u.Status == UserStatusActive
}

// TableName specifies the table name for GORM
func (User) TableName() string {
	return "workforce_users"
}
