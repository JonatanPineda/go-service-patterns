package main

import "context"

type contextKey string

const (
	contextResponseKey = contextKey("response")
)

type Response struct {
	body interface{}
	status int
}

func (s *Server) ResponseCtx(ctx context.Context) *Response {
	response := ctx.Value(contextResponseKey).(*Response)
	return response
}