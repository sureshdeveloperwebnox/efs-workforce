package repositories

import (
	"efs-workforce/internal/domain"
	"github.com/google/uuid"
)

// CrewRepository defines the interface for crew data operations
type CrewRepository interface {
	Create(crew *domain.Crew) error
	FindByID(id uuid.UUID) (*domain.Crew, error)
	FindAll() ([]*domain.Crew, error)
	Update(crew *domain.Crew) error
	Delete(id uuid.UUID) error
}

// CrewMemberRepository defines the interface for crew member data operations
type CrewMemberRepository interface {
	Create(crewMember *domain.CrewMember) error
	FindByID(id uuid.UUID) (*domain.CrewMember, error)
	FindByCrewID(crewID uuid.UUID) ([]*domain.CrewMember, error)
	FindByUserID(userID uuid.UUID) ([]*domain.CrewMember, error)
	Delete(id uuid.UUID) error
	DeleteByCrewAndUser(crewID, userID uuid.UUID) error
}

