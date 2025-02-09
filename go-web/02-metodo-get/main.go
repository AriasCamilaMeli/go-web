package main

import (
	"fmt"
	"net/http"

	"github.com/AriasCamilaMeli/go-web/02-metodo-get/products"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// Controlador para manejar la ruta /ping
func pingHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong"))
}

func main() {

	// Crear un nuevo enrutador
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	// Definir rutas y asignar controladores
	r.Get("/ping", pingHandler)
	r.Get("/products/{id}", products.GetProductByIDHandler())
	r.Get("/products", products.GetAllProductsHandler())
	r.Get("/products/search", products.SearchProductsHandler())

	// Iniciar el servidor
	fmt.Println("Servidor encendido en el puerto http://localhost:8080")
	http.ListenAndServe(":8080", r)
}
