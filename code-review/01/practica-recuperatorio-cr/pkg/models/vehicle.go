package models

import "errors"

// Dimensions is a struct that represents a dimension in 3d
type Dimensions struct {
	// Height is the height of the dimension
	Height float64 `json:"height"`
	// Length is the length of the dimension
	Length float64 `json:"length"`
	// Width is the width of the dimension
	Width float64 `json:"width"`
}

// VehicleAttributes is a struct that represents the attributes of a vehicle
type VehicleAttributes struct {
	// Brand is the brand of the vehicle
	Brand string `json:"brand"`
	// Model is the model of the vehicle
	Model string `json:"model"`
	// Registration is the registration of the vehicle
	Registration string `json:"registration"`
	// Color is the color of the vehicle
	Color string `json:"color"`
	// FabricationYear is the fabrication year of the vehicle
	FabricationYear int `json:"year"`
	// Capacity is the capacity of people of the vehicle
	Capacity int `json:"passengers"`
	// MaxSpeed is the maximum speed of the vehicle
	MaxSpeed float64 `json:"max_speed"`
	// FuelType is the fuel type of the vehicle
	FuelType string `json:"fuel_type"`
	// Transmission is the transmission of the vehicle
	Transmission string `json:"transmission"`
	// Weight is the weight of the vehicle
	Weight float64 `json:"weight"`
	// Dimensions is the dimensions of the vehicle
	Dimensions
}

// Vehicle is a struct that represents a vehicle in JSON format
type VehicleDoc struct {
	ID              int     `json:"id"`
	Brand           string  `json:"brand"`
	Model           string  `json:"model"`
	Registration    string  `json:"registration"`
	Color           string  `json:"color"`
	FabricationYear int     `json:"year"`
	Capacity        int     `json:"passengers"`
	MaxSpeed        float64 `json:"max_speed"`
	FuelType        string  `json:"fuel_type"`
	Transmission    string  `json:"transmission"`
	Weight          float64 `json:"weight"`
	Height          float64 `json:"height"`
	Length          float64 `json:"length"`
	Width           float64 `json:"width"`
}

// Vehicle is a struct that represents a vehicle
type Vehicle struct {
	// Id is the unique identifier of the vehicle
	Id int `json:"id"`

	// VehicleAttribue is the attributes of a vehicle
	VehicleAttributes
}

var (
	AlreadyExistErr = errors.New("Identificador del vehículo ya existente")
	BadRequestErr   = errors.New("Datos del vehículo mal formados o incompletos")
	InternalErr     = errors.New("Error del servidor")
	NotFoundErr     = errors.New("No se encontraron vehículos con esos criterios.")
)
