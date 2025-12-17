package dto

// CreateTimeOffRequest represents the request to create a time off record
type CreateTimeOffRequest struct {
	UserID    string  `json:"user_id" validate:"required,uuid"`
	LeaveType string  `json:"leave_type" validate:"required,oneof=Sick Leave Casual Leave Paid Leave Unpaid Leave"`
	StartDate string  `json:"start_date" validate:"required"`
	EndDate   string  `json:"end_date" validate:"required"`
	Reason    string  `json:"reason" validate:"omitempty"`
	CreatedBy *string `json:"created_by" validate:"omitempty,uuid"`
}

// UpdateTimeOffRequest represents the request to update a time off record
type UpdateTimeOffRequest struct {
	LeaveType string `json:"leave_type" validate:"omitempty,oneof=Sick Leave Casual Leave Paid Leave Unpaid Leave"`
	StartDate string `json:"start_date" validate:"omitempty"`
	EndDate   string `json:"end_date" validate:"omitempty"`
	Reason    string `json:"reason" validate:"omitempty"`
	Status    string `json:"status" validate:"omitempty,oneof=Pending Approved Rejected"`
}

// TimeOffResponse represents the time off response
type TimeOffResponse struct {
	ID        string       `json:"id"`
	UserID    string       `json:"user_id"`
	User      *UserResponse `json:"user,omitempty"`
	LeaveType string       `json:"leave_type"`
	StartDate string       `json:"start_date"`
	EndDate   string       `json:"end_date"`
	Reason    string       `json:"reason"`
	Status    string       `json:"status"`
	CreatedBy *string      `json:"created_by"`
	CreatedAt string       `json:"created_at"`
	UpdatedAt string       `json:"updated_at"`
}

