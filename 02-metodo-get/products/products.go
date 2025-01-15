package products

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"

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
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		productID := chi.URLParam(r, "id")
		id, err := strconv.Atoi(productID)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "{'response':'error'}")
			return
		}

		for _, product := range products {
			if product.ID == id {
				w.Header().Set("Content-Type", "application/json")
				json.NewEncoder(w).Encode(product)
				return
			}
		}
		http.Error(w, "Product not found", http.StatusNotFound)
	}
}

func GetAllProductsHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		products, err := LoadProducts("products.json")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(products)
	}
}

func SearchProductsHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		products, err := LoadProducts("products.json")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		priceGt := r.URL.Query().Get("priceGt")
		priceGtInt, err := strconv.Atoi(priceGt)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "{'response':'error'}")
			return
		}

		var filteredProducts []Product
		for _, product := range products {
			if product.Price > float64(priceGtInt) {
				filteredProducts = append(filteredProducts, product)
			}
		}

		if len(filteredProducts) < 1 {
			w.Header().Set("Content-Type", "application/json")
			http.Error(w, "Product not found", http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(filteredProducts)
	}
}
