package repository

import "app/pkg/models"

// VehicleRepository is an interface that represents a vehicle repository
type VehicleRepository interface {
	// FindAll is a method that returns a map of all vehicles
	FindAll() (v map[int]models.Vehicle, err error)
	// Store is a method that stores a vehicle
	Store(v models.Vehicle) (newV models.Vehicle, err error)
	// GetByColorAndYear metodo para obtener vehicolos por color y año dado
	GetByColorAndYear(color string, year int) (v map[int]models.Vehicle, err error)
	//GetByBrandAndYears
	GetByBrandAndYears(brand string, startYear, endYear int) (v map[int]models.Vehicle, err error)
	//GetVelocityAVGByBrand()
	GetVelocityAVGByBrand(brand string) (avg float32, err error)
	//CreateInBatch()
	CreateInBatch(v []models.Vehicle) (err error)
	//UpdateSpeed()
	UpdateSpeed(id int, speed float64) (err error)
}
