package handler

import (
	"app/internal/service"
	"app/pkg/models"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/bootcamp-go/web/response"
	"github.com/go-chi/chi/v5"
)

// CreateBatch is a method that returns a handler for the route POST /vehicles/batch
func (h *VehicleDefault) CreateBatch() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Leer el cuerpo de la solicitud
		var vehicles []models.Vehicle
		err := json.NewDecoder(r.Body).Decode(&vehicles)
		if err != nil {
			response.JSON(w, http.StatusBadRequest, map[string]any{
				"message": "Datos de algún vehículo mal formados o incompletos.",
			})
			return
		}

		// Enviar los datos al servicio para la creación
		err = h.sv.CreateBatch(vehicles)
		if err != nil {
			// Manejo de errores de validación
			if err.Error() == "Identificador del vehículo ya existente." {
				response.JSON(w, http.StatusConflict, map[string]any{
					"message": err.Error(),
				})
			} else {
				response.JSON(w, http.StatusBadRequest, map[string]any{
					"message":     "Error",
					"description": err.Error(),
				})
			}
			return
		}

		// Responder con los vehículos creados
		response.JSON(w, http.StatusCreated, map[string]any{
			"message": "Vehículos creados exitosamente",
		})
	}
}

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
			response.JSON(w, http.StatusNotFound, map[string]any{
				"message": "No se encontró el vehículo.",
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

func (h *VehicleDefault) GetByWeight() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		min, err := strconv.ParseFloat(r.URL.Query().Get("min"), 64)
		if err != nil {
			response.JSON(w, http.StatusInternalServerError, map[string]any{
				"message": "Hubo en error en los parametros de entrada.",
			})
			return
		}
		max, err := strconv.ParseFloat(r.URL.Query().Get("max"), 64)
		if err != nil {
			response.JSON(w, http.StatusInternalServerError, map[string]any{
				"message": "Hubo en error en los parametros de entrada.",
			})
			return
		}

		v, err := h.sv.GetByWeight(min, max)
		if err != nil {
			response.JSON(w, http.StatusNotFound, map[string]any{
				"message": "No se encontró el vehículo.",
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

		if len(data) < 1 {

			response.JSON(w, http.StatusNotFound, map[string]any{
				"message": "No se encontraron vehículos en ese rango de peso.",
			})
			return
		}

		response.JSON(w, http.StatusOK, map[string]any{
			"message": "success",
			"data":    data,
		})

	}
}

func (h *VehicleDefault) GetDimensions() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		lengthRange := r.URL.Query().Get("length")

		lengths := strings.Split(lengthRange, "-")
		if len(lengths) != 2 {
			response.JSON(w, http.StatusBadRequest, map[string]any{
				"message": "Formato de longitud inválido.",
			})
			return
		}

		min_length, err := strconv.ParseFloat(lengths[0], 64)
		if err != nil {
			response.JSON(w, http.StatusBadRequest, map[string]any{
				"message": "Valor mínimo de longitud inválido.",
			})
			return
		}

		max_length, err := strconv.ParseFloat(lengths[1], 64)
		if err != nil {
			response.JSON(w, http.StatusBadRequest, map[string]any{
				"message": "Valor máximo de longitud inválido.",
			})
			return
		}

		widthRange := r.URL.Query().Get("width")

		widths := strings.Split(widthRange, "-")
		if len(widths) != 2 {
			response.JSON(w, http.StatusBadRequest, map[string]any{
				"message": "Formato de hancho inválido.",
			})
			return
		}

		min_width, err := strconv.ParseFloat(widths[0], 64)
		if err != nil {
			response.JSON(w, http.StatusBadRequest, map[string]any{
				"message": "Valor mínimo de hancho inválido.",
			})
			return
		}

		max_width, err := strconv.ParseFloat(widths[1], 64)
		if err != nil {
			response.JSON(w, http.StatusBadRequest, map[string]any{
				"message": "Valor máximo de hancho inválido.",
			})
			return
		}

		v, err := h.sv.GetDimensions(min_length, max_length, min_width, max_width)
		if err != nil {
			response.JSON(w, http.StatusNotFound, map[string]any{
				"message": "No se encontró el vehículo.",
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

		if len(data) < 1 {
			response.JSON(w, http.StatusNotFound, map[string]any{
				"message": "No se encontraron vehículos con esas dimensiones.",
			})
			return
		}

		response.JSON(w, http.StatusOK, map[string]any{
			"message": "success",
			"data":    data,
		})

	}
}

func (h *VehicleDefault) GetAverageCapacity() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		brand := chi.URLParam(r, "brand")

		average, err := h.sv.GetAverageCapacity(brand)
		if err != nil {
			response.JSON(w, http.StatusNotFound, map[string]any{
				"message": err.Error(),
			})
			return
		}

		response.JSON(w, http.StatusOK, map[string]any{
			"message": fmt.Sprintf("La capacidad promedio es %.2f", average),
		})
	}
}

func (h *VehicleDefault) UpdateFuel() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			response.JSON(w, http.StatusInternalServerError, nil)
			return
		}

		var v models.Vehicle
		err = json.NewDecoder(r.Body).Decode(&v)

		if err != nil {
			response.JSON(w, http.StatusBadRequest, map[string]any{
				"message": "Tipo de combustible mal formado o no admitido.",
			})
			return
		}

		err = h.sv.UpdateFuel(id, v.FuelType)

		if err != nil {
			if err.Error() == "No se encontró el vehículo." {
				response.JSON(w, http.StatusNotFound, map[string]any{
					"message": err.Error(),
				})
			} else {
				response.JSON(w, http.StatusBadRequest, map[string]any{
					"message": "Tipo de combustible mal formado o no admitido.",
				})
			}
			return
		}

		response.JSON(w, http.StatusOK, map[string]any{
			"message": "Tipo de combustible del vehículo actualizado exitosamente.",
		})

	}
}

func (h *VehicleDefault) GetByTransmission() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		type_t := chi.URLParam(r, "type")

		v, err := h.sv.GetByTransmission(type_t)
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

		if len(data) < 1 {
			response.JSON(w, http.StatusNotFound, map[string]any{
				"message": "No se encontraron vehículos con ese tipo de transmisión.",
			})
			return
		}

		response.JSON(w, http.StatusOK, map[string]any{
			"message": "success",
			"data":    data,
		})

	}
}

func (h *VehicleDefault) GetAverageSpped() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		brand := chi.URLParam(r, "brand")

		average_speed, err := h.sv.GetAverageSpped(brand)

		if err != nil {
			response.JSON(w, http.StatusNotFound, map[string]any{
				"message": err.Error(),
			})
			return
		}
		response.JSON(w, http.StatusOK, map[string]any{
			"message":       "success",
			"average speed": average_speed,
		})

	}
}

func (h *VehicleDefault) GetByColorAndYear() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		color := chi.URLParam(r, "color")
		year, err := strconv.Atoi(chi.URLParam(r, "year"))

		if err != nil {
			response.JSON(w, http.StatusInternalServerError, nil)
			return
		}

		v, err := h.sv.GetByColorAndYear(color, year)
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

		if len(data) < 1 {
			response.JSON(w, http.StatusNotFound, map[string]any{
				"message": "error",
				"data":    "No se encontraron vehículos con esos criterios.",
			})
		} else {

			response.JSON(w, http.StatusOK, map[string]any{
				"message": "success",
				"data":    data,
			})
		}

	}
}

func (h *VehicleDefault) GetByBranForRange() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		brand := chi.URLParam(r, "brand")

		start_year, err := strconv.Atoi(chi.URLParam(r, "start_year"))
		if err != nil {
			response.JSON(w, http.StatusInternalServerError, map[string]any{
				"description": err.Error(),
			})
			return
		}

		end_year, err := strconv.Atoi(chi.URLParam(r, "end_year"))

		if err != nil {
			response.JSON(w, http.StatusInternalServerError, map[string]any{
				"description": err.Error(),
			})
			return
		}

		v, err := h.sv.GetByBranForRange(brand, start_year, end_year)

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

		if len(data) < 1 {
			response.JSON(w, http.StatusNotFound, map[string]any{
				"message": "No se encontraron vehículos con esos criterios.",
			})
		} else {

			response.JSON(w, http.StatusOK, map[string]any{
				"message": "success",
				"data":    data,
			})
		}
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

func (h *VehicleDefault) UpdateSpeed() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			response.JSON(w, http.StatusInternalServerError, nil)
			return
		}

		var v models.Vehicle
		err = json.NewDecoder(r.Body).Decode(&v)
		if err != nil {
			response.JSON(w, http.StatusBadRequest, map[string]any{
				"message": "Velocidad mal formada o fuera de rango.",
			})
			return
		}

		err = h.sv.UpdateSpeed(id, v.MaxSpeed)

		if err != nil {
			if err.Error() == "Velocidad mal formada o fuera de rango." {
				response.JSON(w, http.StatusBadRequest, map[string]any{
					"message": err.Error(),
				})
			}
			if err.Error() == "No se encontró el vehículo" {
				response.JSON(w, http.StatusNotFound, map[string]any{
					"message": err.Error(),
				})
			}
			response.JSON(w, http.StatusBadRequest, map[string]any{
				"message": "unknow error",
			})
		}

		response.JSON(w, http.StatusOK, map[string]any{
			"message": "Velocidad del vehículo actualizada exitosamente.",
		})

	}

}
func (h *VehicleDefault) GetByFuelType() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fuel_type := chi.URLParam(r, "type")

		v, err := h.sv.GetByFuelType(fuel_type)

		if err != nil {
			response.JSON(w, http.StatusBadRequest, map[string]any{
				"message": "Bad request",
			})
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

		if len(data) < 1 {
			response.JSON(w, http.StatusNotFound, map[string]any{
				"message": " No se encontraron vehículos con ese tipo de combustible.",
			})
		} else {

			response.JSON(w, http.StatusOK, map[string]any{
				"message": "success",
				"data":    data,
			})
		}

	}
}

func (h *VehicleDefault) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))

		if err != nil {
			response.JSON(w, http.StatusBadRequest, map[string]any{
				"message": "Bad request",
			})
		}

		err = h.sv.Delete(id)

		if err != nil {
			if err.Error() == "No se encontró el vehículo" {
				response.JSON(w, http.StatusNotFound, map[string]any{
					"message": err.Error(),
				})
			} else {

				response.JSON(w, http.StatusBadRequest, map[string]any{
					"message": "Bad request",
				})
			}
			return
		}

		response.JSON(w, http.StatusNoContent, map[string]any{
			"message": "Vehículo eliminado exitosamente",
		})
	}
}
