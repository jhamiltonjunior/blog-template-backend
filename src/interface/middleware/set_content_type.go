package middleware

import "net/http"

// SetContentType middleware go set content-type json in all routes
// 
// So that it is not necessary I set
// response.Header(... in all routes, I created this middleware,
// to be more practical
func SetContentType(next http.Handler) http.Handler {
	return http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
		response.Header().Set("Content-type", "application/json")
		
		next.ServeHTTP(response, request)
	})
}
