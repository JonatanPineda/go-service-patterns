package main

import "net/http"

type MiddlewareFunc func(next http.HandlerFunc) http.HandlerFunc

func chain(chain ...MiddlewareFunc) http.HandlerFunc {
	return http.HandlerFunc(recurseChain(chain))
}

func recurseChain(chain []MiddlewareFunc) http.HandlerFunc {
	if len(chain) <= 0 {
		return func(_ http.ResponseWriter, _ *http.Request) {}
	}

	return chain[0](recurseChain(chain[1:]))
}
