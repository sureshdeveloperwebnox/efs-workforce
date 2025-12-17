package dto

// CreateCrewRequest represents the request to create a crew
type CreateCrewRequest struct {
	CrewName  string  `json:"crew_name" validate:"required,min=1,max=100"`
	CreatedBy *string `json:"created_by" validate:"omitempty,uuid"`
}

// UpdateCrewRequest represents the request to update a crew
type UpdateCrewRequest struct {
	CrewName string `json:"crew_name" validate:"omitempty,min=1,max=100"`
}

// CrewResponse represents the crew response
type CrewResponse struct {
	ID        string            `json:"id"`
	CrewName  string            `json:"crew_name"`
	CreatedBy *string            `json:"created_by"`
	Members   []CrewMemberResponse `json:"members,omitempty"`
	CreatedAt string            `json:"created_at"`
	UpdatedAt string            `json:"updated_at"`
}

// CreateCrewMemberRequest represents the request to add a member to a crew
type CreateCrewMemberRequest struct {
	CrewID string `json:"crew_id" validate:"required,uuid"`
	UserID string `json:"user_id" validate:"required,uuid"`
}

// CrewMemberResponse represents the crew member response
type CrewMemberResponse struct {
	ID         string `json:"id"`
	CrewID     string `json:"crew_id"`
	UserID     string `json:"user_id"`
	User       *UserResponse `json:"user,omitempty"`
	AssignedAt string `json:"assigned_at"`
}

