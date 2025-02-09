package service

import (
	"app/internal/repository"
	"app/pkg/models"
	"errors"
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

func (s *VehicleDefault) GetByWeight(min, max float64) (v map[int]models.Vehicle, err error) {
	v, err = s.rp.GetByWeight(min, max)

	return
}

func (s *VehicleDefault) GetDimensions(min_lenght, max_lenght, min_width, max_width float64) (v map[int]models.Vehicle, err error) {
	v, err = s.rp.GetDimensions(min_lenght, max_lenght, min_width, max_width)

	return
}

func (s *VehicleDefault) GetAverageCapacity(brand string) (average float64, err error) {
	average, err = s.rp.GetAverageCapacity(brand)

	return
}

func (s *VehicleDefault) UpdateFuel(id int, update_fuel string) (err error) {
	err = s.rp.UpdateFuel(id, update_fuel)
	return
}

func (s *VehicleDefault) GetByTransmission(type_t string) (v map[int]models.Vehicle, err error) {
	v, err = s.rp.GetByTransmission(type_t)
	return
}
func (s *VehicleDefault) GetAverageSpped(brand string) (average_speed float64, err error) {
	average_speed, err = s.rp.GetAverageSpped(brand)
	return
}

// FindAll is a method that returns a map of all vehicles
func (s *VehicleDefault) GetByColorAndYear(color string, year int) (v map[int]models.Vehicle, err error) {
	v, err = s.rp.GetByColorAndYear(color, year)
	return
}
func (s *VehicleDefault) GetByBranForRange(brand string, start_year int, end_year int) (v map[int]models.Vehicle, err error) {
	v, err = s.rp.GetByBranForRange(brand, start_year, end_year)
	return
}

// Create
func (s *VehicleDefault) Create(v models.Vehicle) (err error) {
	vAll, err := s.rp.FindAll()
	if err != nil {
		return err
	}

	for _, _v := range vAll {
		if _v.Id == v.Id {
			return errors.New("Identificador del vehículo ya existente.")
		}
	}

	if err := validateVehicle(v); err != nil {
		return err
	}

	err = s.rp.Create(v)
	return
}

func (s *VehicleDefault) GetByFuelType(fuel_type string) (v map[int]models.Vehicle, err error) {
	v, err = s.rp.GetByFuelType(fuel_type)
	return
}

func (s *VehicleDefault) Delete(id int) (err error) {
	err = s.rp.Delete(id)
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

// CreateBatch is a method that creates multiple vehicles at once
func (s *VehicleDefault) CreateBatch(vehicles []models.Vehicle) (err error) {
	vAll, err := s.rp.FindAll()
	if err != nil {
		return err
	}

	for _, v := range vehicles {
		if _, exists := vAll[v.Id]; exists {
			return errors.New("Algún vehículo tiene un identificador ya existente.")
		}
	}

	for _, v := range vehicles {
		if err := validateVehicle(v); err != nil {
			return err
		}
	}

	err = s.rp.CreateBatch(vehicles)
	return
}

func (s *VehicleDefault) UpdateSpeed(id int, speed float64) (err error) {

	if speed < 0 {
		return errors.New("Velocidad mal formada o fuera de rango.")
	}

	err = s.rp.UpdateSpeed(id, speed)
	if err != nil {
		return err
	}

	return
}
