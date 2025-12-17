package application

import (
	"efs-workforce/internal/application/dto"
	"efs-workforce/internal/domain"
	"efs-workforce/internal/ports/external"
	"efs-workforce/internal/ports/repositories"
	"fmt"
	"time"

	"github.com/google/uuid"
)

// RoleService implements role business logic
type RoleService struct {
	roleRepo      repositories.RoleRepository
	eventPublisher external.EventPublisher
}

// NewRoleService creates a new role service
func NewRoleService(
	roleRepo repositories.RoleRepository,
	eventPublisher external.EventPublisher,
) *RoleService {
	return &RoleService{
		roleRepo:      roleRepo,
		eventPublisher: eventPublisher,
	}
}

// CreateRole creates a new role
func (s *RoleService) CreateRole(req *dto.CreateRoleRequest) (*dto.RoleResponse, error) {
	// Check if role already exists
	existing, err := s.roleRepo.FindByName(req.RoleName)
	if err != nil {
		return nil, fmt.Errorf("failed to check if role exists: %w", err)
	}
	if existing != nil {
		return nil, fmt.Errorf("role with name %s already exists", req.RoleName)
	}

	// Create domain entity
	role := &domain.Role{
		RoleName:    req.RoleName,
		Description: req.Description,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	// Save to repository
	if err := s.roleRepo.Create(role); err != nil {
		return nil, fmt.Errorf("failed to create role: %w", err)
	}

	// Publish event
	if s.eventPublisher != nil {
		event := &domain.Event{
			Type:      "RoleCreated",
			Payload:   map[string]interface{}{"role_id": role.ID.String(), "role_name": role.RoleName},
			Timestamp: time.Now(),
		}
		_ = s.eventPublisher.Publish(event)
	}

	return s.toDTO(role), nil
}

// GetRole retrieves a role by ID
func (s *RoleService) GetRole(id string) (*dto.RoleResponse, error) {
	roleID, err := uuid.Parse(id)
	if err != nil {
		return nil, fmt.Errorf("invalid role ID: %w", err)
	}

	role, err := s.roleRepo.FindByID(roleID)
	if err != nil {
		return nil, fmt.Errorf("failed to get role: %w", err)
	}
	if role == nil {
		return nil, fmt.Errorf("role not found")
	}

	return s.toDTO(role), nil
}

// ListRoles retrieves all roles
func (s *RoleService) ListRoles() ([]*dto.RoleResponse, error) {
	roles, err := s.roleRepo.FindAll()
	if err != nil {
		return nil, fmt.Errorf("failed to list roles: %w", err)
	}

	dtos := make([]*dto.RoleResponse, len(roles))
	for i, role := range roles {
		dtos[i] = s.toDTO(role)
	}

	return dtos, nil
}

// UpdateRole updates an existing role
func (s *RoleService) UpdateRole(id string, req *dto.UpdateRoleRequest) (*dto.RoleResponse, error) {
	roleID, err := uuid.Parse(id)
	if err != nil {
		return nil, fmt.Errorf("invalid role ID: %w", err)
	}

	// Get existing role
	role, err := s.roleRepo.FindByID(roleID)
	if err != nil {
		return nil, fmt.Errorf("failed to get role: %w", err)
	}
	if role == nil {
		return nil, fmt.Errorf("role not found")
	}

	// Update fields
	if req.RoleName != "" {
		// Check if new name already exists
		if req.RoleName != role.RoleName {
			existing, err := s.roleRepo.FindByName(req.RoleName)
			if err != nil {
				return nil, fmt.Errorf("failed to check if role exists: %w", err)
			}
			if existing != nil {
				return nil, fmt.Errorf("role with name %s already exists", req.RoleName)
			}
		}
		role.RoleName = req.RoleName
	}
	if req.Description != "" {
		role.Description = req.Description
	}
	role.UpdatedAt = time.Now()

	// Save to repository
	if err := s.roleRepo.Update(role); err != nil {
		return nil, fmt.Errorf("failed to update role: %w", err)
	}

	// Publish event
	if s.eventPublisher != nil {
		event := &domain.Event{
			Type:      "RoleUpdated",
			Payload:   map[string]interface{}{"role_id": role.ID.String(), "role_name": role.RoleName},
			Timestamp: time.Now(),
		}
		_ = s.eventPublisher.Publish(event)
	}

	return s.toDTO(role), nil
}

// DeleteRole deletes a role by ID
func (s *RoleService) DeleteRole(id string) error {
	roleID, err := uuid.Parse(id)
	if err != nil {
		return fmt.Errorf("invalid role ID: %w", err)
	}

	// Check if role exists
	role, err := s.roleRepo.FindByID(roleID)
	if err != nil {
		return fmt.Errorf("failed to get role: %w", err)
	}
	if role == nil {
		return fmt.Errorf("role not found")
	}

	// Delete from repository
	if err := s.roleRepo.Delete(roleID); err != nil {
		return fmt.Errorf("failed to delete role: %w", err)
	}

	// Publish event
	if s.eventPublisher != nil {
		event := &domain.Event{
			Type:      "RoleDeleted",
			Payload:   map[string]interface{}{"role_id": roleID.String()},
			Timestamp: time.Now(),
		}
		_ = s.eventPublisher.Publish(event)
	}

	return nil
}

// toDTO converts a domain entity to a DTO
func (s *RoleService) toDTO(role *domain.Role) *dto.RoleResponse {
	return &dto.RoleResponse{
		ID:          role.ID.String(),
		RoleName:    role.RoleName,
		Description: role.Description,
		CreatedAt:   role.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   role.UpdatedAt.Format(time.RFC3339),
	}
}

