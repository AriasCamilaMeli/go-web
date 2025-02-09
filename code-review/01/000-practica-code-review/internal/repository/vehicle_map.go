package repository

import (
	"app/pkg/models"
	"errors"
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

func (r *VehicleMap) GetByWeight(min, max float64) (v map[int]models.Vehicle, err error) {
	v = make(map[int]models.Vehicle)

	for key, value := range r.db {
		if value.Width <= max && value.Width >= min {
			v[key] = value
		}
	}

	return

}

func (r *VehicleMap) GetDimensions(min_lenght, max_lenght, min_width, max_width float64) (v map[int]models.Vehicle, err error) {
	v = make(map[int]models.Vehicle)

	for key, value := range r.db {
		if value.Length <= max_lenght && value.Length >= min_lenght && value.Width <= max_width && value.Width >= min_width {
			v[key] = value
		}
	}

	return
}

func (r *VehicleMap) GetAverageCapacity(brand string) (average float64, err error) {
	sum := 0
	len := 0

	for _, value := range r.db {
		if value.Brand == brand {
			sum += value.Capacity
			len += 1
		}
	}

	if len < 1 {
		return 0, errors.New("No se encontraron vehículos de esa marca.")
	}

	return float64(sum) / float64(len), nil
}

func (r *VehicleMap) UpdateFuel(id int, update_fuel string) (err error) {
	v, exists := r.db[id]

	if !exists {
		return errors.New("No se encontró el vehículo.")
	}

	v.FuelType = update_fuel

	r.db[id] = v

	return
}

func (r *VehicleMap) GetByTransmission(type_t string) (v map[int]models.Vehicle, err error) {
	v = make(map[int]models.Vehicle)
	for key, value := range r.db {
		if value.Transmission == type_t {
			v[key] = value
		}
	}

	return
}

func (r *VehicleMap) GetAverageSpped(brand string) (average_speed float64, err error) {

	sum := 0.0
	len := 0.0

	// copy db
	for _, value := range r.db {
		if value.Brand == brand {
			sum += value.MaxSpeed
			len += 1
		}
	}

	if len > 0 {
		return sum / len, nil
	}

	return 0, errors.New("No se encontraron vehículos de esa marca.")
}

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

func (r *VehicleMap) GetByBranForRange(brand string, start_year int, end_year int) (v map[int]models.Vehicle, err error) {
	v = make(map[int]models.Vehicle)

	for key, value := range r.db {
		if value.Brand == brand && value.FabricationYear <= end_year && value.FabricationYear >= start_year {
			v[key] = value
		}
	}
	return
}

func (r *VehicleMap) GetByFuelType(fuel_type string) (v map[int]models.Vehicle, err error) {
	v = make(map[int]models.Vehicle)

	for key, value := range r.db {
		if value.FuelType == fuel_type {
			v[key] = value
		}
	}

	return
}

func (r *VehicleMap) Create(v models.Vehicle) (err error) {
	r.db[v.Id] = v
	return
}
func (r *VehicleMap) CreateBatch(vehicles []models.Vehicle) (err error) {
	for _, v := range vehicles {
		if _, exists := r.db[v.Id]; exists {
			return errors.New("Algún vehículo tiene un identificador ya existente.")
		}
	}
	for _, v := range vehicles {
		r.db[v.Id] = v
	}
	return
}

func (r *VehicleMap) UpdateSpeed(id int, speed float64) (err error) {
	v, exists := r.db[id]

	if !exists {
		return errors.New("No se encontró el vehículo")
	}

	v.MaxSpeed = speed

	r.db[id] = v

	return
}

func (r *VehicleMap) Delete(id int) (err error) {

	if _, exists := r.db[id]; !exists {
		return errors.New("No se encontró el vehículo")
	}

	delete(r.db, id)

	return
}
