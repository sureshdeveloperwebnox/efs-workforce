package postgres

import (
	"efs-workforce/internal/domain"
	"efs-workforce/internal/ports/repositories"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// UserRepository implements the user repository interface using PostgreSQL
type UserRepository struct {
	db *gorm.DB
}

// NewUserRepository creates a new PostgreSQL user repository
func NewUserRepository(db *gorm.DB) repositories.UserRepository {
	return &UserRepository{db: db}
}

// Create creates a new user
func (r *UserRepository) Create(user *domain.User) error {
	return r.db.Create(user).Error
}

// FindByID finds a user by ID
func (r *UserRepository) FindByID(id uuid.UUID) (*domain.User, error) {
	var user domain.User
	if err := r.db.Preload("Role").Preload("Creator").First(&user, "id = ?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

// FindByEmail finds a user by email
func (r *UserRepository) FindByEmail(email string) (*domain.User, error) {
	var user domain.User
	if err := r.db.Preload("Role").Where("email = ?", email).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

// FindByEmployeeID finds a user by employee ID
func (r *UserRepository) FindByEmployeeID(employeeID string) (*domain.User, error) {
	var user domain.User
	if err := r.db.Preload("Role").Where("employee_id = ?", employeeID).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

// FindAll finds all users
func (r *UserRepository) FindAll() ([]*domain.User, error) {
	var users []*domain.User
	if err := r.db.Preload("Role").Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

// FindByRoleID finds users by role ID
func (r *UserRepository) FindByRoleID(roleID uuid.UUID) ([]*domain.User, error) {
	var users []*domain.User
	if err := r.db.Preload("Role").Where("role_id = ?", roleID).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

// Update updates an existing user
func (r *UserRepository) Update(user *domain.User) error {
	return r.db.Save(user).Error
}

// Delete deletes a user by ID
func (r *UserRepository) Delete(id uuid.UUID) error {
	return r.db.Delete(&domain.User{}, "id = ?", id).Error
}

