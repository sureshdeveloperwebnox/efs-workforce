package postgres

import (
	"efs-workforce/internal/domain"
	"efs-workforce/internal/ports/repositories"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// EquipmentRepository implements the equipment repository interface using PostgreSQL
type EquipmentRepository struct {
	db *gorm.DB
}

// NewEquipmentRepository creates a new PostgreSQL equipment repository
func NewEquipmentRepository(db *gorm.DB) repositories.EquipmentRepository {
	return &EquipmentRepository{db: db}
}

// Create creates a new equipment
func (r *EquipmentRepository) Create(equipment *domain.Equipment) error {
	return r.db.Create(equipment).Error
}

// FindByID finds an equipment by ID
func (r *EquipmentRepository) FindByID(id uuid.UUID) (*domain.Equipment, error) {
	var equipment domain.Equipment
	if err := r.db.Preload("User").Preload("Creator").First(&equipment, "id = ?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &equipment, nil
}

// FindBySerialNumber finds an equipment by serial number
func (r *EquipmentRepository) FindBySerialNumber(serialNumber string) (*domain.Equipment, error) {
	var equipment domain.Equipment
	if err := r.db.Preload("User").Where("serial_number = ?", serialNumber).First(&equipment).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &equipment, nil
}

// FindByUserID finds equipment by user ID
func (r *EquipmentRepository) FindByUserID(userID uuid.UUID) ([]*domain.Equipment, error) {
	var equipment []*domain.Equipment
	if err := r.db.Where("assigned_to_user = ?", userID).Find(&equipment).Error; err != nil {
		return nil, err
	}
	return equipment, nil
}

// FindAll finds all equipment
func (r *EquipmentRepository) FindAll() ([]*domain.Equipment, error) {
	var equipment []*domain.Equipment
	if err := r.db.Preload("User").Preload("Creator").Find(&equipment).Error; err != nil {
		return nil, err
	}
	return equipment, nil
}

// Update updates an existing equipment
func (r *EquipmentRepository) Update(equipment *domain.Equipment) error {
	return r.db.Save(equipment).Error
}

// Delete deletes an equipment by ID
func (r *EquipmentRepository) Delete(id uuid.UUID) error {
	return r.db.Delete(&domain.Equipment{}, "id = ?", id).Error
}

