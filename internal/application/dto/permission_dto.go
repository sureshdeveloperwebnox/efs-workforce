package dto

// CreatePermissionRequest represents the request to create a permission
type CreatePermissionRequest struct {
	RoleID     string `json:"role_id" validate:"required,uuid"`
	ModuleName string `json:"module_name" validate:"required,min=1,max=50"`
	CanCreate  bool   `json:"can_create"`
	CanRead    bool   `json:"can_read"`
	CanUpdate  bool   `json:"can_update"`
	CanDelete  bool   `json:"can_delete"`
}

// UpdatePermissionRequest represents the request to update a permission
type UpdatePermissionRequest struct {
	ModuleName string `json:"module_name" validate:"omitempty,min=1,max=50"`
	CanCreate  bool   `json:"can_create"`
	CanRead    bool   `json:"can_read"`
	CanUpdate  bool   `json:"can_update"`
	CanDelete  bool   `json:"can_delete"`
}

// PermissionResponse represents the permission response
type PermissionResponse struct {
	ID         string       `json:"id"`
	RoleID     string       `json:"role_id"`
	Role       *RoleResponse `json:"role,omitempty"`
	ModuleName string       `json:"module_name"`
	CanCreate  bool         `json:"can_create"`
	CanRead    bool         `json:"can_read"`
	CanUpdate  bool         `json:"can_update"`
	CanDelete  bool         `json:"can_delete"`
	CreatedAt  string       `json:"created_at"`
	UpdatedAt  string       `json:"updated_at"`
}

