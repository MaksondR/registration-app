package mdlwr

import (
	"net/http"
	"registration-app/helper"
	"strings"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.Contains(r.URL.RequestURI(), "api") {
			token := r.URL.Query().Get("token")

			if !helper.JwtValidate(token) {
				helper.RespondWithError(w, http.StatusForbidden, "Invalid authorization token!")
			}
		}
		next.ServeHTTP(w, r)
	})
}
