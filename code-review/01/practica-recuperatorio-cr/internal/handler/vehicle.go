package handler

import (
	"app/internal/service"
	"app/pkg/models"
	"encoding/json"
	"errors"
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

// Create is a method that returns a handler for the route POST /vehicles
func (h *VehicleDefault) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		var v models.Vehicle
		err := json.NewDecoder(r.Body).Decode(&v)
		if err != nil {
			response.JSON(w, http.StatusBadRequest, map[string]any{
				"message":     "Error",
				"description": models.BadRequestErr.Error(),
			})
		}
		// process
		newV, err := h.sv.Save(v)
		if err != nil {
			if errors.Is(err, models.BadRequestErr) {
				response.JSON(w, http.StatusBadRequest, map[string]any{
					"Error description": err.Error(),
				})
			} else if errors.Is(err, models.AlreadyExistErr) {
				response.JSON(w, http.StatusConflict, map[string]any{
					"Error description": err.Error(),
				})
			} else {
				response.JSON(w, http.StatusInternalServerError, map[string]any{
					"Error description": err.Error(),
				})
			}
			return

		}

		// response
		response.JSON(w, http.StatusOK, map[string]any{
			"message": "success",
			"data":    newV,
		})

	}

}
