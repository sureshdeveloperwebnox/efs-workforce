package dto

// CreateUserRequest represents the request to create a user
type CreateUserRequest struct {
	FirstName  string  `json:"first_name" validate:"required,min=1,max=50"`
	LastName   string  `json:"last_name" validate:"required,min=1,max=50"`
	EmployeeID string  `json:"employee_id" validate:"required,min=1,max=20"`
	Email      string  `json:"email" validate:"required,email,max=100"`
	Phone      string  `json:"phone" validate:"omitempty,max=20"`
	Status     string  `json:"status" validate:"omitempty,oneof=Active Inactive"`
	Profile    string  `json:"profile" validate:"omitempty,oneof=Field Agent Manager Administrator"`
	RoleID     *string `json:"role_id" validate:"omitempty,uuid"`
	CreatedBy  *string `json:"created_by" validate:"omitempty,uuid"`
}

// UpdateUserRequest represents the request to update a user
type UpdateUserRequest struct {
	FirstName  string  `json:"first_name" validate:"omitempty,min=1,max=50"`
	LastName   string  `json:"last_name" validate:"omitempty,min=1,max=50"`
	EmployeeID string  `json:"employee_id" validate:"omitempty,min=1,max=20"`
	Email      string  `json:"email" validate:"omitempty,email,max=100"`
	Phone      string  `json:"phone" validate:"omitempty,max=20"`
	Status     string  `json:"status" validate:"omitempty,oneof=Active Inactive"`
	Profile    string  `json:"profile" validate:"omitempty,oneof=Field Agent Manager Administrator"`
	RoleID     *string `json:"role_id" validate:"omitempty,uuid"`
}

// UserResponse represents the user response
type UserResponse struct {
	ID         string       `json:"id"`
	FirstName  string       `json:"first_name"`
	LastName   string       `json:"last_name"`
	EmployeeID string       `json:"employee_id"`
	Email      string       `json:"email"`
	Phone      string       `json:"phone"`
	Status     string       `json:"status"`
	Profile    string       `json:"profile"`
	RoleID     *string      `json:"role_id"`
	Role       *RoleResponse `json:"role,omitempty"`
	CreatedBy  *string      `json:"created_by"`
	CreatedAt  string       `json:"created_at"`
	UpdatedAt  string       `json:"updated_at"`
}

