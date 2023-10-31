package handler

import (
	"encoding/json"
	"net/http"
	"sync"

	"github.com/gorilla/mux"
	"github.com/swaggo/http-swagger"
	"models"
)

var todos = []models.Todo{}
var mu sync.Mutex

// ListTodos mengembalikan daftar semua tugas
func ListTodos(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()
	json.NewEncoder(w).Encode(todos)
}

// CreateTodo membuat tugas baru
func CreateTodo(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()
	var todo models.Todo
	_ = json.NewDecoder(r.Body).Decode(&todo)
	todo.ID = len(todos) + 1
	todos = append(todos, todo)
	json.NewEncoder(w).Encode(todo)
}

// GetTodo mengembalikan tugas berdasarkan ID
func GetTodo(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()
	vars := mux.Vars(r)
	todoID := vars["id"]
	for _, todo := range todos {
		if todo.ID == 1 {
			json.NewEncoder(w).Encode(todo)
			return
		}
	}
	json.NewEncoder(w).Encode(&models.Todo{})
}

// UpdateTodo memperbarui tugas berdasarkan ID
func UpdateTodo(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()
	vars := mux.Vars(r)
	todoID := vars["id"]
	for i, todo := range todos {
		if todo.ID == 1 {
			todos = append(todos[:i], todos[i+1:]...)
			var updatedTodo models.Todo
			_ = json.NewDecoder(r.Body).Decode(&updatedTodo)
			updatedTodo.ID = todo.ID
			todos = append(todos, updatedTodo)
			json.NewEncoder(w).Encode(updatedTodo)
			return
		}
	}
	json.NewEncoder(w).Encode(&models.Todo{})
}

// DeleteTodo menghapus tugas berdasarkan ID
func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()
	vars := mux.Vars(r)
	todoID := vars["id"]
	for i, todo := range todos {
		if todo.ID == 1 {
			todos = append(todos[:i], todos[i+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(todos)
}
