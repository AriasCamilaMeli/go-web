package products

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
)

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Quantity    int     `json:"quantity"`
	CodeValue   string  `json:"code_value"`
	IsPublished bool    `json:"is_published"`
	Expiration  string  `json:"expiration"`
	Price       float64 `json:"price"`
}

type Response struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

func LoadProducts(filename string) ([]Product, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	var products []Product
	err = json.Unmarshal(data, &products)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON: %w", err)
	}

	return products, nil
}

func GetProductByIDHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		products, err := LoadProducts("products.json")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(Response{Message: "Failed to load products", Error: err.Error()})
			return
		}

		productID := chi.URLParam(r, "id")
		id, err := strconv.Atoi(productID)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(Response{Message: "Invalid product ID", Error: err.Error()})
			return
		}

		for _, product := range products {
			if product.ID == id {
				w.Header().Set("Content-Type", "application/json")
				json.NewEncoder(w).Encode(Response{Message: "Product found", Data: product})
				return
			}
		}
		w.WriteHeader(http.StatusNotFound)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(Response{Message: "Product not found"})
	}
}

func GetAllProductsHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		products, err := LoadProducts("products.json")
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(Response{Message: "Failed to load products", Error: err.Error()})
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(Response{Message: "Products retrieved successfully", Data: products})
	}
}

func SearchProductsHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		products, err := LoadProducts("products.json")
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(Response{Message: "Failed to load products", Error: err.Error()})
			return
		}

		priceGt := r.URL.Query().Get("priceGt")
		priceGtInt, err := strconv.Atoi(priceGt)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(Response{Message: "Invalid priceGt parameter", Error: err.Error()})
			return
		}

		var filteredProducts []Product
		for _, product := range products {
			if product.Price > float64(priceGtInt) {
				filteredProducts = append(filteredProducts, product)
			}
		}

		if len(filteredProducts) < 1 {
			w.WriteHeader(http.StatusNotFound)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(Response{Message: "No products found"})
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(Response{Message: "Products found", Data: filteredProducts})
	}
}

func AddProductHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var newProduct Product

		w.Header().Set("Content-Type", "application/json")

		err := json.NewDecoder(r.Body).Decode(&newProduct)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(Response{Message: "Invalid request payload", Error: err.Error()})
			return
		}

		products, err := LoadProducts("products.json")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(Response{Message: "Failed to load products", Error: err.Error()})
			return
		}

		if newProduct.Name == "" || newProduct.Quantity == 0 || newProduct.CodeValue == "" || newProduct.Expiration == "" || newProduct.Price == 0 {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(Response{Message: "All fields except is_published must be filled"})
			return
		}

		for _, product := range products {
			if product.CodeValue == newProduct.CodeValue {
				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode(Response{Message: "Code value must be unique"})
				return
			}
		}

		if _, err := time.Parse("02/01/2006", newProduct.Expiration); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(Response{Message: "Invalid expiration date format", Error: err.Error()})
			return
		}

		maxID := 0
		for _, product := range products {
			if product.ID > maxID {
				maxID = product.ID
			}
		}
		newProduct.ID = maxID + 1

		products = append(products, newProduct)

		data, err := json.Marshal(products)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(Response{Message: "Failed to marshal products", Error: err.Error()})
			return
		}

		err = os.WriteFile("products.json", data, 0644)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(Response{Message: "Failed to write products to file", Error: err.Error()})
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(Response{Message: "Product added successfully", Data: newProduct})
	}
}
