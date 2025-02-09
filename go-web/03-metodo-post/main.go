package main

import (
	"fmt"
	"net/http"

	"github.com/AriasCamilaMeli/go-web/03-metodo-post/products"
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

	// Definir rutas para ping
	r.Get("/ping", pingHandler)

	// Definir rutas para productos
	r.Route("/products", func(r chi.Router) {
		r.Get("/", products.GetAllProductsHandler())
		r.Get("/{id}", products.GetProductByIDHandler())
		r.Get("/search", products.SearchProductsHandler())
		r.Post("/", products.AddProductHandler())
	})

	// Iniciar el servidor
	fmt.Println("Servidor encendido en el puerto http://localhost:8080")
	http.ListenAndServe(":8080", r)
}
