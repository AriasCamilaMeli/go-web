package repository

import (
	"app/pkg/models"
)

// NewVehicleMap is a function that returns a new instance of VehicleMap
func NewVehicleMap(db map[int]models.Vehicle) *VehicleMap {
	// default db
	defaultDb := make(map[int]models.Vehicle)
	if db != nil {
		defaultDb = db
	}
	return &VehicleMap{db: defaultDb}
}

// VehicleMap is a struct that represents a vehicle repository
type VehicleMap struct {
	// db is a map of vehicles
	db map[int]models.Vehicle
}

// FindAll is a method that returns a map of all vehicles
func (r *VehicleMap) FindAll() (v map[int]models.Vehicle, err error) {
	v = make(map[int]models.Vehicle)

	// copy db
	for key, value := range r.db {
		v[key] = value
	}

	return
}

// Store is a method that stores a vehicle
func (r *VehicleMap) Store(v models.Vehicle) (newV models.Vehicle, err error) {
	_, exists := r.db[v.Id]
	if exists {
		return newV, models.AlreadyExistErr
	}
	r.db[v.Id] = v
	newV = r.db[v.Id]
	return
}

// GetByColorAndYear metodo para obtener vehicolos por color y año dado
func (r *VehicleMap) GetByColorAndYear(color string, year int) (v map[int]models.Vehicle, err error) {
	v = make(map[int]models.Vehicle)

	// copy db
	for key, value := range r.db {
		if value.Color == color && value.FabricationYear == year {
			v[key] = value
		}
	}

	return
}

// GetByBrandAndYears
func (r *VehicleMap) GetByBrandAndYears(brand string, startYear, endYear int) (v map[int]models.Vehicle, err error) {
	v = make(map[int]models.Vehicle)

	// copy db
	for key, value := range r.db {
		if value.Brand == brand && value.FabricationYear >= startYear && value.FabricationYear <= endYear {
			v[key] = value
		}
	}

	return
}

// GetVelocityAVGByBrand()
func (r *VehicleMap) GetVelocityAVGByBrand(brand string) (avg float32, err error) {

	var (
		contador = 0
		suma     = 0
	)
	// copy db
	for _, value := range r.db {
		if value.Brand == brand {
			contador += 1
			suma += value.Capacity
		}
	}

	if contador == 0 {
		return 0, models.NotFoundErr
	}

	return (float32(suma) / float32(contador)), nil

}

// CreateInBatch()
func (r *VehicleMap) CreateInBatch(v []models.Vehicle) (err error) {
	for _, value := range v {
		_, exists := r.db[value.Id]
		if exists {
			err = models.AlreadyExistErr
			return
		}
	}
	for _, value := range v {
		r.db[value.Id] = value
	}
	return

}

// UpdateSpeed()
func (r *VehicleMap) UpdateSpeed(id int, speed float64) (err error) {
	v, exists := r.db[id]

	if !exists {
		err = models.NotFoundErr
		return
	}

	v.MaxSpeed = speed

	r.db[id] = v

	return nil
}
