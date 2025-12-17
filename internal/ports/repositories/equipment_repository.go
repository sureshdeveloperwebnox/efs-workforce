package repositories

import (
	"efs-workforce/internal/domain"
	"github.com/google/uuid"
)

// EquipmentRepository defines the interface for equipment data operations
type EquipmentRepository interface {
	Create(equipment *domain.Equipment) error
	FindByID(id uuid.UUID) (*domain.Equipment, error)
	FindBySerialNumber(serialNumber string) (*domain.Equipment, error)
	FindByUserID(userID uuid.UUID) ([]*domain.Equipment, error)
	FindAll() ([]*domain.Equipment, error)
	Update(equipment *domain.Equipment) error
	Delete(id uuid.UUID) error
}

