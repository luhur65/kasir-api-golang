package middlewares

import (
	"net/http"
	"api-kasir/dto"
)

func ApiKey(apiKey string) func(http.HandlerFunc) http.HandlerFunc {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			
			if r.Header.Get("X-API-KEY") != apiKey {
				dto.WriteError(w, http.StatusForbidden, "Invalid API Key")
				return
			}

			if r.Header.Get("X-API-KEY") == "" {
				dto.WriteError(w, http.StatusBadRequest, "Missing API Key")
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}