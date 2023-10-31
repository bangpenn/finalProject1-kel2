package main

import (
	"log"
	"net/http"

	"finalProject-1/todo-api/handler"
	"github.com/gorilla/mux"
	"github.com/swaggo/http-swagger"
)

func main() {
	// Membuat router
	router := mux.NewRouter()

	// Rute untuk daftar tugas
	router.HandleFunc("/todos", handler.ListTodos).Methods("GET")
	router.HandleFunc("/todos", handler.CreateTodo).Methods("POST")

	// Rute untuk tugas berdasarkan ID
	router.HandleFunc("/todos/{id}", handler.GetTodo).Methods("GET")
	router.HandleFunc("/todos/{id}", handler.UpdateTodo).Methods("PUT")
	router.HandleFunc("/todos/{id}", handler.DeleteTodo).Methods("DELETE")

	// Rute untuk Swagger UI
	router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	// Mulai server HTTP
	port := ":8080"
	log.Println("Server berjalan di port", port)
	log.Fatal(http.ListenAndServe(port, router))
}
