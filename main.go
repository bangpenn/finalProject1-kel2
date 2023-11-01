package main

import (
	"log"
	"net/http"

	"finalProject-1/todo-api/docs"
	"finalProject-1/todo-api/handler"

	"github.com/gorilla/mux"
	"github.com/swaggo/http-swagger"
)

func main() {

	docs.SwaggerInfo.Title = "Final Project 1"
	docs.SwaggerInfo.Description = "Final Project 1 tentang ToDo"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.Schemes = []string{"http"}

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

	// Handle not found (404) errors
	router.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Not Found", http.StatusNotFound)
	})

	// Handle other errors
	router.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				if r := recover(); r != nil {
					http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				}
			}()
			next.ServeHTTP(w, r)
		})
	})

	// Mulai server HTTP
	port := ":8080"
	log.Println("Server berjalan di port", port)
	log.Fatal(http.ListenAndServe(port, router))
}
