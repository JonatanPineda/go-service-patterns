package main

import (
	"github.com/go-chi/chi"
	"net/http"
)

func main() {
	db := NewInMemoryStorage()
	getUsers := NewGetUsers(db)

	r := chi.NewRouter()

	r.Get("/", chain(logger, respond, getUsers))
	http.ListenAndServe(":3000", r)
}

func NewGetUsers(db Storage) func(next http.HandlerFunc) http.HandlerFunc {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			users := db.FindUsers()
			response := ResponseCtx(r.Context())
			response.body = users
			next.ServeHTTP(w, r)
		}
	}
}
