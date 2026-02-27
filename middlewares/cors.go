package middlewares

import "net/http"

// function name (variable type) return
func CORS(next http.HandlerFunc)http.HandlerFunc{ 
	return func(w http.ResponseWriter, r *http.Request) { 
		w.Header().Set("Access-Controll-Allow-Origin", "*")
		w.Header().Set("Access-Controll-Allow-Method", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Controll-Allow-Origin", "X-API-Key, Content-Type")
	
		if r.Method == "OPTIONS"{
			w.WriteHeader(http.StatusOK)
			return
		}

		next(w,r)
	
	}
}