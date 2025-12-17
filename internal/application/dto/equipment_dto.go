package dto

// CreateEquipmentRequest represents the request to create equipment
type CreateEquipmentRequest struct {
	Name          string  `json:"name" validate:"required,min=1,max=100"`
	SerialNumber  string  `json:"serial_number" validate:"omitempty,max=50"`
	AssignedToUser *string `json:"assigned_to_user" validate:"omitempty,uuid"`
	Status        string  `json:"status" validate:"omitempty,oneof=Active Inactive Under Maintenance"`
	CreatedBy     *string `json:"created_by" validate:"omitempty,uuid"`
}

// UpdateEquipmentRequest represents the request to update equipment
type UpdateEquipmentRequest struct {
	Name          string  `json:"name" validate:"omitempty,min=1,max=100"`
	SerialNumber  string  `json:"serial_number" validate:"omitempty,max=50"`
	AssignedToUser *string `json:"assigned_to_user" validate:"omitempty,uuid"`
	Status        string  `json:"status" validate:"omitempty,oneof=Active Inactive Under Maintenance"`
}

// EquipmentResponse represents the equipment response
type EquipmentResponse struct {
	ID            string       `json:"id"`
	Name          string       `json:"name"`
	SerialNumber  string       `json:"serial_number"`
	AssignedToUser *string     `json:"assigned_to_user"`
	User          *UserResponse `json:"user,omitempty"`
	Status        string       `json:"status"`
	CreatedBy     *string      `json:"created_by"`
	CreatedAt     string       `json:"created_at"`
	UpdatedAt     string       `json:"updated_at"`
}

