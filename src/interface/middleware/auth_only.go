package middleware

import "net/http"

func AuthOnly(next http.Handler) http.Handler {
	return http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
		response.Header().Set("Content-type", "application/json")
		
		next.ServeHTTP(response, request)
	})
}