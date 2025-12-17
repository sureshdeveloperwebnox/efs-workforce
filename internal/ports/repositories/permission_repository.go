package repositories

import (
	"efs-workforce/internal/domain"
	"github.com/google/uuid"
)

// PermissionRepository defines the interface for permission data operations
type PermissionRepository interface {
	Create(permission *domain.Permission) error
	FindByID(id uuid.UUID) (*domain.Permission, error)
	FindByRoleID(roleID uuid.UUID) ([]*domain.Permission, error)
	FindByRoleIDAndModule(roleID uuid.UUID, moduleName string) (*domain.Permission, error)
	FindAll() ([]*domain.Permission, error)
	Update(permission *domain.Permission) error
	Delete(id uuid.UUID) error
	DeleteByRoleID(roleID uuid.UUID) error
}

