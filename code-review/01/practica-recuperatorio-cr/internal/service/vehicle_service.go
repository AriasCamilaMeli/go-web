package service

import "app/pkg/models"

// VehicleService is an interface that represents a vehicle service
type VehicleService interface {
	// FindAll is a method that returns a map of all vehicles
	FindAll() (v map[int]models.Vehicle, err error)
	// Save is a method that stores a vehicle
	Save(v models.Vehicle) (newV models.Vehicle, err error)
}
