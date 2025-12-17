package postgres

import (
	"efs-workforce/internal/domain"
	"efs-workforce/internal/ports/repositories"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// RoleRepository implements the role repository interface using PostgreSQL
type RoleRepository struct {
	db *gorm.DB
}

// NewRoleRepository creates a new PostgreSQL role repository
func NewRoleRepository(db *gorm.DB) repositories.RoleRepository {
	return &RoleRepository{db: db}
}

// Create creates a new role
func (r *RoleRepository) Create(role *domain.Role) error {
	return r.db.Create(role).Error
}

// FindByID finds a role by ID
func (r *RoleRepository) FindByID(id uuid.UUID) (*domain.Role, error) {
	var role domain.Role
	if err := r.db.First(&role, "id = ?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &role, nil
}

// FindByName finds a role by name
func (r *RoleRepository) FindByName(name string) (*domain.Role, error) {
	var role domain.Role
	if err := r.db.Where("role_name = ?", name).First(&role).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &role, nil
}

// FindAll finds all roles
func (r *RoleRepository) FindAll() ([]*domain.Role, error) {
	var roles []*domain.Role
	if err := r.db.Find(&roles).Error; err != nil {
		return nil, err
	}
	return roles, nil
}

// Update updates an existing role
func (r *RoleRepository) Update(role *domain.Role) error {
	return r.db.Save(role).Error
}

// Delete deletes a role by ID
func (r *RoleRepository) Delete(id uuid.UUID) error {
	return r.db.Delete(&domain.Role{}, "id = ?", id).Error
}

