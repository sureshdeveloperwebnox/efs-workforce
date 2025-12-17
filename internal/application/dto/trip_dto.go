package dto

// CreateTripRequest represents the request to create a trip
type CreateTripRequest struct {
	UserID       string  `json:"user_id" validate:"required,uuid"`
	StartLocation string `json:"start_location" validate:"required,min=1,max=255"`
	EndLocation   string `json:"end_location" validate:"required,min=1,max=255"`
	StartTime     string `json:"start_time" validate:"required"`
	EndTime       *string `json:"end_time" validate:"omitempty"`
	Purpose       string  `json:"purpose" validate:"omitempty"`
	DistanceKm    *float64 `json:"distance_km" validate:"omitempty,min=0"`
	CreatedBy     *string  `json:"created_by" validate:"omitempty,uuid"`
}

// UpdateTripRequest represents the request to update a trip
type UpdateTripRequest struct {
	StartLocation string  `json:"start_location" validate:"omitempty,min=1,max=255"`
	EndLocation   string  `json:"end_location" validate:"omitempty,min=1,max=255"`
	StartTime     string  `json:"start_time" validate:"omitempty"`
	EndTime       *string `json:"end_time" validate:"omitempty"`
	Purpose       string  `json:"purpose" validate:"omitempty"`
	DistanceKm    *float64 `json:"distance_km" validate:"omitempty,min=0"`
}

// TripResponse represents the trip response
type TripResponse struct {
	ID           string       `json:"id"`
	UserID       string       `json:"user_id"`
	User         *UserResponse `json:"user,omitempty"`
	StartLocation string       `json:"start_location"`
	EndLocation   string       `json:"end_location"`
	StartTime     string       `json:"start_time"`
	EndTime       *string      `json:"end_time"`
	Purpose       string       `json:"purpose"`
	DistanceKm    *float64     `json:"distance_km"`
	CreatedBy     *string      `json:"created_by"`
	CreatedAt     string       `json:"created_at"`
	UpdatedAt     string       `json:"updated_at"`
}

