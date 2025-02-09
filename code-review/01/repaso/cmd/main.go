package main

import (
	"net/http"

	"github.com/alkemy/repaso/internal/transport/http/chi/router"
	"github.com/go-chi/chi/v5"
)

func main() {
	rt := chi.NewRouter()

	rt.Mount("/movies", router.NewChiMoviesRouter())

	if err := http.ListenAndServe(":8080", rt); err != nil {
		panic(err)
	}
}