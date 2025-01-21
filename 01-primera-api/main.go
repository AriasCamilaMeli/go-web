package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

// Ejercicio 1 - Prueba de Ping
// Vamos a crear una aplicación web con el package net/http nativo de go, que tenga un endpoint /ping que al pegarle responda un texto que diga “pong”
// 1. El endpoint deberá ser de método GET
// 2. La respuesta de “pong” deberá ser enviada como texto, NO como JSON
func PingHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("pong"))
}

// Ejercicio 2 - Manipulando el body
// Vamos a crear un endpoint llamado /greetings. Con una pequeña estructura con nombre y apellido que al pegarle deberá responder en texto “Hello + nombre + apellido”
// 1. El endpoint deberá ser de método POST
// 2. Se deberá usar el package JSON para resolver el ejercicio
// 3. La respuesta deberá seguir esta estructura: “Hello Andrea Rivas”
// 4. La estructura deberá ser como esta:
// 	{
//		“firstName”: “Andrea”,
//		“lastName”: “Rivas”
// 	}

type Person struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

func GretingsHandler(w http.ResponseWriter, r *http.Request) {
	var p Person
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&p); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		http.Error(w, "Error to parse json", http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	response := fmt.Sprintf("Hello %s %s", p.FirstName, p.LastName)
	w.Write([]byte(response))

}

func main() {
	rt := chi.NewRouter()

	rt.Get("/ping", PingHandler)
	rt.Post("/greetings", GretingsHandler)

	fmt.Println("Server http://localhost:8080")
	http.ListenAndServe(":8080", rt)

}
