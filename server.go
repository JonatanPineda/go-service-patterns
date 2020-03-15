package main

import (
	"context"
	"encoding/json"
	"github.com/go-chi/chi"
	"log"
	"net/http"
	"time"
)

type Server struct {
	router *chi.Mux
	storage Storage
}

func NewServer(router *chi.Mux, storage Storage) *Server {
	return &Server{
		router: router,
		storage: storage,
	}
}

func (s *Server) logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		t1 := time.Now()
		next.ServeHTTP(w, r)
		t2 := time.Now()
		log.Printf("[%s] %q %v\n", r.Method, r.URL.String(), t2.Sub(t1))
	}
}

func (s *Server) respond(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), contextResponseKey, &Response{})
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
		response := s.ResponseCtx(r.Context())
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(response.status)
		json.NewEncoder(w).Encode(response.body)
	}
}
