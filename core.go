package main

import "net/http"

func (s *Server) getUsers(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		users := s.storage.FindUsers()
		response := s.ResponseCtx(r.Context())
		response.status = http.StatusOK
		response.body = users
		next.ServeHTTP(w, r)
	}
}
