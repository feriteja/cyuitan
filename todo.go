package main

// Todo struct represents a todo item
type Todo struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

var todos []Todo

// initTodos initializes the todos slice with some dummy data
func initTodos() {
	todos = []Todo{
		{ID: "1", Title: "Buy groceries", Done: false},
		{ID: "2", Title: "Read a book", Done: true},
	}
}
