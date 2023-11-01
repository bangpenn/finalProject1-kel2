package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"sync"

	"finalProject-1/todo-api/models"

	"github.com/gorilla/mux"
)

var users = make(map[int]models.User)
var todos []models.Todo

var mu sync.Mutex

// ListTodos mengembalikan daftar semua tugas
// ListTodos godoc
// @Tags ToDos
// @Summary Get All Todos
// @Produce json
// @Success 200 {object} []models.Todo
// @Router /todos [get]
func ListTodos(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	var todosWithUserId []map[string]interface{}

	for _, todo := range todos {
		user, exists := users[todo.ID]
		if !exists {
			// Handle this case as needed
		}

		todoWithUserId := map[string]interface{}{
			"userId":    user.ID,
			"id":        todo.ID,
			"title":     todo.Title,
			"completed": todo.Completed,
		}

		todosWithUserId = append(todosWithUserId, todoWithUserId)
	}

	json.NewEncoder(w).Encode(todos)
}

// CreateTodo membuat tugas baru
// CreateTodo godoc
// @Tags ToDos
// @Summary Create ToDo
// @ID create-todo
// @Accept json
// @Produce json
// @Param todo body models.Todo true "Todo object that needs to be created"
// @Success 200 {object} models.Todo
// @Router /todos [post]
func CreateTodo(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()
	var todo models.Todo
	_ = json.NewDecoder(r.Body).Decode(&todo)
	todo.ID = len(todos) + 1
	todos = append(todos, todo)

	// tambahkan ID todo ke daftar todos milik pengguna
	userID := todo.ID
	user, exists := users[userID]
	if exists {
		user.Todos = append(user.Todos, todo.ID)
		users[userID] = user

	} else {
		// buat pengguna baru jika belum ada
		newUser := models.User{
			ID:    userID,
			Todos: []int{todo.ID},
		}
		users[userID] = newUser
	}
	json.NewEncoder(w).Encode(todo)
}

// GetTodo mengembalikan tugas berdasarkan ID
// GetTodo godoc
// @Tags ToDos
// @Summary Get a Todo
// @Produce json
// @Param id path int true "Todo ID"
// @Success 200 {object} models.Todo
// @Router /todos/{id} [get]
func GetTodo(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()
	vars := mux.Vars(r)
	todoID := vars["id"]
	id, err := strconv.Atoi(todoID)
	if err != nil {
		http.Error(w, "Invalid todo ID", http.StatusBadRequest)
		return
	}
	for _, todo := range todos {
		if todo.ID == id {
			// Pastikan pengguna adalah pemilik tugas
			userID := todo.ID
			user, exists := users[userID]
			if !exists || user.ID != userID {
				http.Error(w, "Unauthorized access to this todo", http.StatusUnauthorized)
				return
			}
			json.NewEncoder(w).Encode(todo)
			return
		}
	}
	json.NewEncoder(w).Encode(&models.Todo{})
}

// UpdateTodo memperbarui tugas berdasarkan ID
// UpdateTodo godoc
// @Tags ToDos
// @Summary Update a Todo
// @Produce json
// @Param id path int true "Todo ID"
// @Param todo body models.Todo true "Todo object that needs to be updated"
// @Success 200 {object} models.Todo
// @Router /todos/{id} [put]
func UpdateTodo(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()
	vars := mux.Vars(r)
	todoID := vars["id"]
	id, err := strconv.Atoi(todoID)
	if err != nil {
		http.Error(w, "Invalid todo ID", http.StatusBadRequest)
		return
	}
	for i, todo := range todos {
		if todo.ID == id {
			// Pastikan pengguna adalah pemilik tugas
			userID := todo.ID
			user, exists := users[userID]
			if !exists || user.ID != userID {
				http.Error(w, "Unauthorized access to this todo", http.StatusUnauthorized)
				return
			}
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
// DeleteTodo godoc
// @Tags ToDos
// @Summary Delete a Todo
// @Produce json
// @Param id path int true "Todo ID"
// @Success 200 {object} []models.Todo
// @Router /todos/{id} [delete]
func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()
	vars := mux.Vars(r)
	todoID := vars["id"]
	id, err := strconv.Atoi(todoID)
	if err != nil {
		http.Error(w, "Invalid todo ID", http.StatusBadRequest)
		return
	}
	for i, todo := range todos {
		if todo.ID == id {
			// Pastikan pengguna adalah pemilik tugas
			userID := todo.ID
			user, exists := users[userID]
			if !exists || user.ID != userID {
				http.Error(w, "Unauthorized access to this todo", http.StatusUnauthorized)
				return
			}

			todos = append(todos[:i], todos[i+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(todos)
}
