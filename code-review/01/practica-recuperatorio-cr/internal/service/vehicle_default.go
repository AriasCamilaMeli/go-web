package service

import (
	"app/internal/repository"
	"app/pkg/models"
	"fmt"
)

// NewVehicleDefault is a function that returns a new instance of VehicleDefault
func NewVehicleDefault(rp repository.VehicleRepository) *VehicleDefault {
	return &VehicleDefault{rp: rp}
}

// VehicleDefault is a struct that represents the default service for vehicles
type VehicleDefault struct {
	// rp is the repository that will be used by the service
	rp repository.VehicleRepository
}

// FindAll is a method that returns a map of all vehicles
func (s *VehicleDefault) FindAll() (v map[int]models.Vehicle, err error) {
	v, err = s.rp.FindAll()
	return
}

// Save is a method that stores a vehicle
func (s *VehicleDefault) Save(v models.Vehicle) (newV models.Vehicle, err error) {
	err = validateVehicle(v)
	if err != nil {
		return newV, err
	}
	newV, err = s.rp.Store(v)
	return
}

// GetByColorAndYear metodo para obtener vehicolos por color y año dado
func (s *VehicleDefault) GetByColorAndYear(color string, year int) (v map[int]models.Vehicle, err error) {
	v, err = s.rp.GetByColorAndYear(color, year)
	if err != nil {
		err = models.InternalErr
	}

	if len(v) < 1 {
		err = models.NotFoundErr
	}

	return
}

// GetByBrandAndYears
func (s *VehicleDefault) GetByBrandAndYears(brand string, startYear, endYear int) (v map[int]models.Vehicle, err error) {
	v, err = s.rp.GetByBrandAndYears(brand, startYear, endYear)
	if err != nil {
		return v, models.InternalErr
	}

	if len(v) <= 1 {
		return v, models.NotFoundErr
	}

	return
}

// GetVelocityAVGByBrand()
func (s *VehicleDefault) GetVelocityAVGByBrand(brand string) (avg float32, err error) {
	avg, err = s.rp.GetVelocityAVGByBrand(brand)
	return
}

// CreateInBatch()
func (s *VehicleDefault) CreateInBatch(v []models.Vehicle) (err error) {
	for _, value := range v {
		err = validateVehicle(value)
		if err != nil {
			return
		}
	}

	err = s.rp.CreateInBatch(v)
	return
}

// UpdateSpeed()
func (s *VehicleDefault) UpdateSpeed(id int, speed float64) (err error) {
	if speed < 1 {
		err = models.BadRequestErr
		return
	}
	err = s.rp.UpdateSpeed(id, speed)
	return
}

// Función que valida los campos del vehículo
func validateVehicle(v models.Vehicle) error {
	// Verificar si los campos obligatorios no están vacíos
	if v.Brand == "" {
		return fmt.Errorf("%w: %v", models.BadRequestErr, "La marca del vehículo es obligatoria.")
	}
	if v.Model == "" {
		return fmt.Errorf("%w: %v", models.BadRequestErr, "El modelo del vehículo es obligatorio.")
	}
	if v.Registration == "" {
		return fmt.Errorf("%w: %v", models.BadRequestErr, "La matrícula del vehículo es obligatoria.")
	}
	if v.Color == "" {
		return fmt.Errorf("%w: %v", models.BadRequestErr, "El color del vehículo es obligatorio.")
	}

	currentYear := 2025
	if v.FabricationYear > currentYear {
		return fmt.Errorf("%w: %v", models.BadRequestErr, "El año de fabricación es inválido.")
	}

	if v.MaxSpeed <= 0 {
		return fmt.Errorf("%w: %v", models.BadRequestErr, "La velocidad máxima debe ser mayor que cero.")
	}

	if v.Capacity <= 0 {
		return fmt.Errorf("%w: %v", models.BadRequestErr, "La capacidad de personas debe ser mayor que cero.")
	}

	if v.Height <= 0 || v.Length <= 0 || v.Width <= 0 {
		return fmt.Errorf("%w: %v", models.BadRequestErr, "Las dimensiones del vehículo deben ser mayores que cero.")
	}

	return nil
}
