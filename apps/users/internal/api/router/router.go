package router

import (
	"net/http"
)

func New(createUser, getUsers, getUser, authUser func(w http.ResponseWriter, r *http.Request)) *http.ServeMux {
	router := http.NewServeMux()

	router.HandleFunc("POST /user", createUser)

	// TODO: don't think that we need it now since API gateway should takes it
	router.HandleFunc("GET /users", getUsers)

	router.HandleFunc("GET /user/{id}", getUser)

	router.HandleFunc("POST /user/auth", authUser)

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
