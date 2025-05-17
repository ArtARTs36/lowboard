package middlewares

import (
	"net/http"

	"github.com/rs/cors"
)

func CORS(handler http.Handler, origins []string) http.Handler {
	handleCORS := cors.New(cors.Options{
		AllowCredentials:    true,
		AllowPrivateNetwork: true,
		AllowedHeaders:      []string{"*"},
		AllowedOrigins:      origins,
		AllowedMethods: []string{
			http.MethodHead,
			http.MethodGet,
			http.MethodPost,
			http.MethodPatch,
			http.MethodDelete,
		},
	}).Handler

	return handleCORS(handler)
}
