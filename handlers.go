package main

import (
	"encoding/json"
	"net/http"
)

func todosHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getTodos(w, r)
	case http.MethodPost:
		createTodo(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func todoHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getTodoByID(w, r)
	case http.MethodPut:
		updateTodo(w, r)
	case http.MethodDelete:
		deleteTodo(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func getTodos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)
}

func createTodo(w http.ResponseWriter, r *http.Request) {
	var newTodo Todo
	err := json.NewDecoder(r.Body).Decode(&newTodo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Generate ID (In real-world scenario, use UUID or any other unique identifier)
	newTodo.ID = "3"

	todos = append(todos, newTodo)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newTodo)
}

func getTodoByID(w http.ResponseWriter, r *http.Request) {
	// Extract todo ID from URL
	id := r.URL.Path[len("/todos/"):]

	for _, todo := range todos {
		if todo.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(todo)
			return
		}
	}

	http.NotFound(w, r)
}

func updateTodo(w http.ResponseWriter, r *http.Request) {
	// Extract todo ID from URL
	id := r.URL.Path[len("/todos/"):]

	for i, todo := range todos {
		if todo.ID == id {
			var updatedTodo Todo
			err := json.NewDecoder(r.Body).Decode(&updatedTodo)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			todos[i] = updatedTodo
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(updatedTodo)
			return
		}
	}

	http.NotFound(w, r)
}

func deleteTodo(w http.ResponseWriter, r *http.Request) {
	// Extract todo ID from URL
	id := r.URL.Path[len("/todos/"):]

	for i, todo := range todos {
		if todo.ID == id {
			todos = append(todos[:i], todos[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}

	http.NotFound(w, r)
}
