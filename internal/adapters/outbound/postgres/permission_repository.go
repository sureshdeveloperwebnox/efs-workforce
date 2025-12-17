package postgres

import (
	"efs-workforce/internal/domain"
	"efs-workforce/internal/ports/repositories"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// PermissionRepository implements the permission repository interface using PostgreSQL
type PermissionRepository struct {
	db *gorm.DB
}

// NewPermissionRepository creates a new PostgreSQL permission repository
func NewPermissionRepository(db *gorm.DB) repositories.PermissionRepository {
	return &PermissionRepository{db: db}
}

// Create creates a new permission
func (r *PermissionRepository) Create(permission *domain.Permission) error {
	return r.db.Create(permission).Error
}

// FindByID finds a permission by ID
func (r *PermissionRepository) FindByID(id uuid.UUID) (*domain.Permission, error) {
	var permission domain.Permission
	if err := r.db.Preload("Role").First(&permission, "id = ?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &permission, nil
}

// FindByRoleID finds permissions by role ID
func (r *PermissionRepository) FindByRoleID(roleID uuid.UUID) ([]*domain.Permission, error) {
	var permissions []*domain.Permission
	if err := r.db.Preload("Role").Where("role_id = ?", roleID).Find(&permissions).Error; err != nil {
		return nil, err
	}
	return permissions, nil
}

// FindByRoleIDAndModule finds a permission by role ID and module name
func (r *PermissionRepository) FindByRoleIDAndModule(roleID uuid.UUID, moduleName string) (*domain.Permission, error) {
	var permission domain.Permission
	if err := r.db.Preload("Role").Where("role_id = ? AND module_name = ?", roleID, moduleName).First(&permission).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &permission, nil
}

// FindAll finds all permissions
func (r *PermissionRepository) FindAll() ([]*domain.Permission, error) {
	var permissions []*domain.Permission
	if err := r.db.Preload("Role").Find(&permissions).Error; err != nil {
		return nil, err
	}
	return permissions, nil
}

// Update updates an existing permission
func (r *PermissionRepository) Update(permission *domain.Permission) error {
	return r.db.Save(permission).Error
}

// Delete deletes a permission by ID
func (r *PermissionRepository) Delete(id uuid.UUID) error {
	return r.db.Delete(&domain.Permission{}, "id = ?", id).Error
}

// DeleteByRoleID deletes all permissions for a role
func (r *PermissionRepository) DeleteByRoleID(roleID uuid.UUID) error {
	return r.db.Where("role_id = ?", roleID).Delete(&domain.Permission{}).Error
}

