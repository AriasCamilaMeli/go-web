package handler

import (
	"app/internal/service"
	"app/pkg/models"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/bootcamp-go/web/response"
	"github.com/go-chi/chi/v5"
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
			var (
				statusCode int
				errorMsg   string
			)

			switch {
			case errors.Is(err, models.BadRequestErr):
				statusCode = http.StatusBadRequest
				errorMsg = err.Error()
			case errors.Is(err, models.AlreadyExistErr):
				statusCode = http.StatusConflict
				errorMsg = err.Error()
			default:
				statusCode = http.StatusInternalServerError
				errorMsg = err.Error()
			}

			response.JSON(w, statusCode, map[string]any{
				"Error description": errorMsg,
			})
			return
		}

		// response
		response.JSON(w, http.StatusOK, map[string]any{
			"message": "success",
			"data":    newV,
		})

	}

}

// handler para buscar vehiculos por color y año
func (h *VehicleDefault) GetByColorAndYear() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		color := chi.URLParam(r, "color")
		year, err := strconv.Atoi(chi.URLParam(r, "year"))

		if err != nil {
			response.JSON(w, http.StatusBadRequest, map[string]any{
				"Error description": models.BadRequestErr.Error(),
			})
			return
		}

		v, err := h.sv.GetByColorAndYear(color, year)

		if err != nil {
			if errors.Is(err, models.NotFoundErr) {
				response.JSON(w, http.StatusNotFound, map[string]any{
					"Error description": err.Error(),
				})
			} else {
				response.JSON(w, http.StatusInternalServerError, map[string]any{
					"Error description": models.InternalErr.Error(),
				})
			}
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

func (h *VehicleDefault) GetByBrandAndYears() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		brand := chi.URLParam(r, "brand")
		startYear, err := strconv.Atoi(chi.URLParam(r, "start_year"))

		if err != nil {
			response.JSON(w, http.StatusBadRequest, map[string]any{
				"error description": models.BadRequestErr.Error(),
			})
			return
		}

		endYear, err := strconv.Atoi(chi.URLParam(r, "end_year"))

		if err != nil {
			response.JSON(w, http.StatusBadRequest, map[string]any{
				"error description": models.BadRequestErr.Error(),
			})
			return
		}

		v, err := h.sv.GetByBrandAndYears(brand, startYear, endYear)

		if err != nil {

			var (
				statusCode int
				errorMsg   string
			)
			switch {
			case errors.Is(err, models.BadRequestErr):
				statusCode = http.StatusBadRequest
				errorMsg = err.Error()
			case errors.Is(err, models.NotFoundErr):
				statusCode = http.StatusNotFound
				errorMsg = err.Error()
			default:
				statusCode = http.StatusInternalServerError
				errorMsg = err.Error()
			}

			response.JSON(w, statusCode, map[string]any{
				"Error description": errorMsg,
			})
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

func (h *VehicleDefault) GetVelocityAVGByBrand() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		brand := chi.URLParam(r, "brand")

		avg, err := h.sv.GetVelocityAVGByBrand(brand)

		if err != nil {
			var statusCode int

			switch {
			case errors.Is(err, models.BadRequestErr):
				statusCode = http.StatusBadRequest
			case errors.Is(err, models.NotFoundErr):
				statusCode = http.StatusNotFound
			default:
				statusCode = http.StatusInternalServerError
			}

			response.JSON(w, statusCode, map[string]any{
				"Error description": err.Error(),
			})

			return
		}

		response.JSON(w, http.StatusOK, map[string]any{
			"Average capacity by brand": avg,
		})
	}
}

func (h *VehicleDefault) CreateInBatch() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var v []models.Vehicle
		err := json.NewDecoder(r.Body).Decode(&v)
		if err != nil {
			response.JSON(w, http.StatusInternalServerError, map[string]any{
				"message": models.InternalErr.Error(),
			})
			return
		}

		err = h.sv.CreateInBatch(v)
		if err != nil {
			var statusCode int
			switch {
			case errors.Is(err, models.AlreadyExistErr):
				statusCode = http.StatusConflict
			case errors.Is(err, models.BadRequestErr):
				statusCode = http.StatusBadRequest
			default:
				statusCode = http.StatusInternalServerError
			}

			response.JSON(w, statusCode, map[string]any{
				"message": err.Error(),
			})
			return

		}
		response.JSON(w, http.StatusCreated, map[string]any{
			"message": "Vehículos creados exitosamente.",
		})

	}
}

func (h *VehicleDefault) UpdateSpeed() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			response.JSON(w, http.StatusBadRequest, map[string]any{
				"message": models.BadRequestErr.Error(),
			})
			return
		}
		var v models.Vehicle
		err = json.NewDecoder(r.Body).Decode(&v)
		if err != nil {
			response.JSON(w, http.StatusBadRequest, map[string]any{
				"message": models.BadRequestErr.Error(),
			})
			return
		}

		err = h.sv.UpdateSpeed(id, v.MaxSpeed)

		if err != nil {
			var statusCode int
			switch {
			case errors.Is(err, models.BadRequestErr):
				statusCode = http.StatusBadRequest
			case errors.Is(err, models.NotFoundErr):
				statusCode = http.StatusNotFound
			default:
				statusCode = http.StatusInternalServerError
			}
			response.JSON(w, statusCode, map[string]any{
				"message": err.Error(),
			})
			return
		}
		response.JSON(w, http.StatusOK, map[string]any{
			"message": "Velocidad del vehículo actualizada exitosamente",
		})
	}
}
