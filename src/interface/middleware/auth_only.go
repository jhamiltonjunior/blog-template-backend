package middleware

import (
	"net/http"
)

// don't use it's working, it's not working
func AuthOnly(next http.Handler) http.Handler {
	return http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {

		// token := uuid.NewString()
		// expiresAt := time.Now().Add(120 * time.Second)

		// http.SetCookie(response, &http.Cookie{
		// 	Name:    "session_token",
		// 	Value:   token,
		// 	Expires: expiresAt,
		// })

		next.ServeHTTP(response, request)
	})
}
