package service

import "app/pkg/models"

// VehicleService is an interface that represents a vehicle service
type VehicleService interface {
	// FindAll is a method that returns a map of all vehicles
	FindAll() (v map[int]models.Vehicle, err error)
	GetByColorAndYear(color string, year int) (v map[int]models.Vehicle, err error)
	GetByBranForRange(brand string, start_year int, end_year int) (v map[int]models.Vehicle, err error)
	GetAverageSpped(brand string) (average_speed float64, err error)
	GetByFuelType(fuel_type string) (v map[int]models.Vehicle, err error)
	Create(v models.Vehicle) (err error)
	CreateBatch(vehicles []models.Vehicle) (err error)
	Delete(id int) (err error)
	UpdateSpeed(id int, speed float64) (err error)
	GetByTransmission(type_t string) (v map[int]models.Vehicle, err error)
	UpdateFuel(id int, update_fuel string) (err error)
	GetAverageCapacity(brand string) (average float64, err error)
	GetDimensions(min_lenght, max_lenght, min_width, max_width float64) (v map[int]models.Vehicle, err error)
	GetByWeight(min, max float64) (v map[int]models.Vehicle, err error)
}
