package dto

// CreateRoleRequest represents the request to create a role
type CreateRoleRequest struct {
	RoleName    string `json:"role_name" validate:"required,min=1,max=50"`
	Description string `json:"description" validate:"max=255"`
}

// UpdateRoleRequest represents the request to update a role
type UpdateRoleRequest struct {
	RoleName    string `json:"role_name" validate:"omitempty,min=1,max=50"`
	Description string `json:"description" validate:"omitempty,max=255"`
}

// RoleResponse represents the role response
type RoleResponse struct {
	ID          string `json:"id"`
	RoleName    string `json:"role_name"`
	Description string `json:"description"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

