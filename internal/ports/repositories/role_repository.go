package repositories

import (
	"efs-workforce/internal/domain"
	"github.com/google/uuid"
)

// RoleRepository defines the interface for role data operations
type RoleRepository interface {
	Create(role *domain.Role) error
	FindByID(id uuid.UUID) (*domain.Role, error)
	FindByName(name string) (*domain.Role, error)
	FindAll() ([]*domain.Role, error)
	Update(role *domain.Role) error
	Delete(id uuid.UUID) error
}

