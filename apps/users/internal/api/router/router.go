package router

import (
	"net/http"
)

func New(createUser, getUsers, getUser func(w http.ResponseWriter, r *http.Request)) *http.ServeMux {
	router := http.NewServeMux()

	router.HandleFunc("POST /user", createUser)

	router.HandleFunc("GET /users", getUsers)

	router.HandleFunc("GET /user/{id}", getUser)

	// router.HandleFunc("PATCH /todos/{id}", func(w http.ResponseWriter, r *http.Request) {
	// 	id := r.PathValue("id")
	// 	fmt.Println("update a todo by id", id)
	// })

	// router.HandleFunc("DELETE /todos/{id}", func(w http.ResponseWriter, r *http.Request) {
	// 	id := r.PathValue("id")
	// 	fmt.Println("delete a todo by id", id)
	// })

	return router
}
