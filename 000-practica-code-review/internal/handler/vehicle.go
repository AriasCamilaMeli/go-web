package handler

import (
	"app/internal/service"
	"app/pkg/models"
	"encoding/json"
	"net/http"

	"github.com/bootcamp-go/web/response"
)

// NewVehicleDefault is a function that returns a new instance of VehicleDefault
func NewVehicleDefault(sv service.VehicleService) *VehicleDefault {
	return &VehicleDefault{sv: sv}
}

// VehicleDefault is a struct with methods that represent handlers for vehicles
type VehicleDefault struct {
	// sv is the service that will be used by the handler
	sv service.VehicleService
}

// GetAll is a method that returns a handler for the route GET /vehicles
func (h *VehicleDefault) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		// ...

		// process
		// - get all vehicles
		v, err := h.sv.FindAll()
		if err != nil {
			response.JSON(w, http.StatusInternalServerError, nil)
			return
		}

		// response
		data := make(map[int]models.VehicleDoc)
		for key, value := range v {
			data[key] = models.VehicleDoc{
				ID:              value.Id,
				Brand:           value.Brand,
				Model:           value.Model,
				Registration:    value.Registration,
				Color:           value.Color,
				FabricationYear: value.FabricationYear,
				Capacity:        value.Capacity,
				MaxSpeed:        value.MaxSpeed,
				FuelType:        value.FuelType,
				Transmission:    value.Transmission,
				Weight:          value.Weight,
				Height:          value.Height,
				Length:          value.Length,
				Width:           value.Width,
			}
		}
		response.JSON(w, http.StatusOK, map[string]any{
			"message": "success",
			"data":    data,
		})
	}
}

func (h *VehicleDefault) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Leer el cuerpo de la solicitud
		var v models.Vehicle
		err := json.NewDecoder(r.Body).Decode(&v)
		if err != nil {
			response.JSON(w, http.StatusBadRequest, map[string]any{
				"message": "Invalid request body",
			})
			return
		}

		// Enviar los datos al servicio para la creación
		err = h.sv.Create(v)
		if err != nil {
			// Manejo de errores de validación
			if err.Error() == "Identificador del vehículo ya existente." {
				response.JSON(w, http.StatusConflict, map[string]any{
					"message": err.Error(),
				})
			} else {
				response.JSON(w, http.StatusBadRequest, map[string]any{
					"message":     "Datos del vehículo mal formados o incompletos",
					"description": err.Error(),
				})
			}
			return
		}

		// Responder con el vehículo creado
		response.JSON(w, http.StatusCreated, v)
	}
}
