package service

import (
	"app/internal/repository"
	"app/pkg/models"
	"errors"
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
		return newV, fmt.Errorf("%w: %v", models.BadRequestErr, err)
	}
	newV, err = s.rp.Store(v)
	return
}

// Función que valida los campos del vehículo
func validateVehicle(v models.Vehicle) error {
	// Verificar si los campos obligatorios no están vacíos
	if v.Brand == "" {
		return errors.New("La marca del vehículo es obligatoria.")
	}
	if v.Model == "" {
		return errors.New("El modelo del vehículo es obligatorio.")
	}
	if v.Registration == "" {
		return errors.New("La matrícula del vehículo es obligatoria.")
	}
	if v.Color == "" {
		return errors.New("El color del vehículo es obligatorio.")
	}

	currentYear := 2025
	if v.FabricationYear > currentYear {
		return errors.New("El año de fabricación es inválido.")
	}

	if v.MaxSpeed <= 0 {
		return errors.New("La velocidad máxima debe ser mayor que cero.")
	}

	if v.Capacity <= 0 {
		return errors.New("La capacidad de personas debe ser mayor que cero.")
	}

	if v.Height <= 0 || v.Length <= 0 || v.Width <= 0 {
		return errors.New("Las dimensiones del vehículo deben ser mayores que cero.")
	}

	return nil
}
