package main

import (
	"log"
	"net/http"
)

func (app *application) enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		allowedOrigins := map[string]bool{
			"https://logocomp.netlify.app": true,
		}

		origin := r.Header.Get("Origin")

		if _, ok := allowedOrigins[origin]; ok {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
			log.Printf("CORS allowed for origin: %s\n", origin)
		} else {
			log.Printf("CORS not allowed for origin: %s\n", origin)
		}

		next.ServeHTTP(w, r)
	})
}