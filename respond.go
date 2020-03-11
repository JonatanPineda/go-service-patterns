package main

import (
	"context"
	"encoding/json"
	"net/http"
)

type contextKey string

type Response struct {
	body interface{}
}

const (
	contextResponseKey = contextKey("response")
)

func ResponseCtx(ctx context.Context) *Response {
	response := ctx.Value(contextResponseKey).(*Response)
	return response
}

func respond(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		ctx = context.WithValue(ctx, contextResponseKey, &Response{})
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
		response := ResponseCtx(r.Context())
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response.body)
	}
}
