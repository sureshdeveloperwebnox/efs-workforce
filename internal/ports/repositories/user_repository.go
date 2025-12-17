package repositories

import (
	"efs-workforce/internal/domain"
	"github.com/google/uuid"
)

// UserRepository defines the interface for user data operations
type UserRepository interface {
	Create(user *domain.User) error
	FindByID(id uuid.UUID) (*domain.User, error)
	FindByEmail(email string) (*domain.User, error)
	FindByEmployeeID(employeeID string) (*domain.User, error)
	FindAll() ([]*domain.User, error)
	FindByRoleID(roleID uuid.UUID) ([]*domain.User, error)
	Update(user *domain.User) error
	Delete(id uuid.UUID) error
}

