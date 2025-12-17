package postgres

import (
	"efs-workforce/internal/domain"
	"efs-workforce/internal/ports/repositories"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// CrewRepository implements the crew repository interface using PostgreSQL
type CrewRepository struct {
	db *gorm.DB
}

// NewCrewRepository creates a new PostgreSQL crew repository
func NewCrewRepository(db *gorm.DB) repositories.CrewRepository {
	return &CrewRepository{db: db}
}

// Create creates a new crew
func (r *CrewRepository) Create(crew *domain.Crew) error {
	return r.db.Create(crew).Error
}

// FindByID finds a crew by ID
func (r *CrewRepository) FindByID(id uuid.UUID) (*domain.Crew, error) {
	var crew domain.Crew
	if err := r.db.Preload("Creator").Preload("Members").Preload("Members.User").First(&crew, "id = ?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &crew, nil
}

// FindAll finds all crews
func (r *CrewRepository) FindAll() ([]*domain.Crew, error) {
	var crews []*domain.Crew
	if err := r.db.Preload("Creator").Preload("Members").Find(&crews).Error; err != nil {
		return nil, err
	}
	return crews, nil
}

// Update updates an existing crew
func (r *CrewRepository) Update(crew *domain.Crew) error {
	return r.db.Save(crew).Error
}

// Delete deletes a crew by ID
func (r *CrewRepository) Delete(id uuid.UUID) error {
	return r.db.Delete(&domain.Crew{}, "id = ?", id).Error
}

// CrewMemberRepository implements the crew member repository interface using PostgreSQL
type CrewMemberRepository struct {
	db *gorm.DB
}

// NewCrewMemberRepository creates a new PostgreSQL crew member repository
func NewCrewMemberRepository(db *gorm.DB) repositories.CrewMemberRepository {
	return &CrewMemberRepository{db: db}
}

// Create creates a new crew member
func (r *CrewMemberRepository) Create(crewMember *domain.CrewMember) error {
	return r.db.Create(crewMember).Error
}

// FindByID finds a crew member by ID
func (r *CrewMemberRepository) FindByID(id uuid.UUID) (*domain.CrewMember, error) {
	var crewMember domain.CrewMember
	if err := r.db.Preload("Crew").Preload("User").First(&crewMember, "id = ?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &crewMember, nil
}

// FindByCrewID finds crew members by crew ID
func (r *CrewMemberRepository) FindByCrewID(crewID uuid.UUID) ([]*domain.CrewMember, error) {
	var crewMembers []*domain.CrewMember
	if err := r.db.Preload("User").Where("crew_id = ?", crewID).Find(&crewMembers).Error; err != nil {
		return nil, err
	}
	return crewMembers, nil
}

// FindByUserID finds crew members by user ID
func (r *CrewMemberRepository) FindByUserID(userID uuid.UUID) ([]*domain.CrewMember, error) {
	var crewMembers []*domain.CrewMember
	if err := r.db.Preload("Crew").Where("user_id = ?", userID).Find(&crewMembers).Error; err != nil {
		return nil, err
	}
	return crewMembers, nil
}

// Delete deletes a crew member by ID
func (r *CrewMemberRepository) Delete(id uuid.UUID) error {
	return r.db.Delete(&domain.CrewMember{}, "id = ?", id).Error
}

// DeleteByCrewAndUser deletes a crew member by crew ID and user ID
func (r *CrewMemberRepository) DeleteByCrewAndUser(crewID, userID uuid.UUID) error {
	return r.db.Where("crew_id = ? AND user_id = ?", crewID, userID).Delete(&domain.CrewMember{}).Error
}

