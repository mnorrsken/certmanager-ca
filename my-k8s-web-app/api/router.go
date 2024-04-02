package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Router returns a new router instance with registered API endpoints.
func Router() *mux.Router {
	router := mux.NewRouter()

	// Register API endpoints
	router.HandleFunc("/api/certificates", createCertificateHandler).Methods(http.MethodPost)
	router.HandleFunc("/api/certificates/{name}", deleteCertificateHandler).Methods(http.MethodDelete)
	router.HandleFunc("/api/secrets/{name}", downloadSecretHandler).Methods(http.MethodGet)

	return router
}