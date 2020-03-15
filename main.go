package main

import (
	"github.com/go-chi/chi"
	"net/http"
)

func main() {
	storage := NewInMemoryStorage()
	router := chi.NewRouter()
	s := NewServer(router, storage)
	s.routes()
	http.ListenAndServe(":3000", router)
}
