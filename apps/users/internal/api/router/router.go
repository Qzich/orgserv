package router

import (
	"fmt"
	"net/http"
)

func New() *http.ServeMux {
	router := http.NewServeMux()

	router.HandleFunc("POST /todos", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("create a todo")
	})

	router.HandleFunc("GET /todos", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("get all todos")
	})

	router.HandleFunc("PATCH /todos/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		fmt.Println("update a todo by id", id)
	})

	router.HandleFunc("DELETE /todos/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		fmt.Println("delete a todo by id", id)
	})

	return router
}
