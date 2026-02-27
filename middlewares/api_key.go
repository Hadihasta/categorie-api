package middlewares

import "net/http"

// func (api key) func handler http.handler

func APIKEY(validApiKey string) func(http.HandlerFunc) http.HandlerFunc {
	return func(next http.HandlerFunc) http.HandlerFunc {
		// before function running
		return func(w http.ResponseWriter, r *http.Request) {
			apiKey := r.Header.Get("X-API-Key")

			if apiKey == "" {
				http.Error(w, "API key Required", http.StatusUnauthorized)
				return
			}

			if apiKey != validApiKey {
				http.Error(w, "invalid Api key", http.StatusUnauthorized)
				return 
			}

			next(w,r)
			// after function running
		}
	}
}
