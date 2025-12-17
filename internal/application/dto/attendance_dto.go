package dto

// CreateAttendanceRequest represents the request to create an attendance record
type CreateAttendanceRequest struct {
	UserID   string  `json:"user_id" validate:"required,uuid"`
	CheckIn  *string `json:"check_in" validate:"omitempty"`
	CheckOut *string `json:"check_out" validate:"omitempty"`
	Status   string  `json:"status" validate:"omitempty,oneof=Present Absent On Leave"`
	CreatedBy *string `json:"created_by" validate:"omitempty,uuid"`
}

// UpdateAttendanceRequest represents the request to update an attendance record
type UpdateAttendanceRequest struct {
	CheckIn  *string `json:"check_in" validate:"omitempty"`
	CheckOut *string `json:"check_out" validate:"omitempty"`
	Status   string  `json:"status" validate:"omitempty,oneof=Present Absent On Leave"`
}

// AttendanceResponse represents the attendance response
type AttendanceResponse struct {
	ID        string       `json:"id"`
	UserID    string       `json:"user_id"`
	User      *UserResponse `json:"user,omitempty"`
	CheckIn   *string      `json:"check_in"`
	CheckOut  *string      `json:"check_out"`
	Status    string       `json:"status"`
	CreatedBy *string      `json:"created_by"`
	CreatedAt string       `json:"created_at"`
	UpdatedAt string       `json:"updated_at"`
}

