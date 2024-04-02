package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Server represents the API server.
type Server struct {
	router *mux.Router
}

// NewServer creates a new instance of the API server.
func NewServer() *Server {
	return &Server{
		router: mux.NewRouter(),
	}
}

// Start starts the API server and listens for incoming requests.
func (s *Server) Start(addr string) error {
	log.Printf("Starting server on %s\n", addr)
	return http.ListenAndServe(addr, s.router)
}

// InitializeRoutes initializes the API routes.
func (s *Server) InitializeRoutes() {
	s.router.HandleFunc("/certificates", s.createCertificate).Methods("POST")
	s.router.HandleFunc("/certificates/{name}", s.deleteCertificate).Methods("DELETE")
	s.router.HandleFunc("/secrets/{name}", s.downloadSecretData).Methods("GET")
}

func (s *Server) createCertificate(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement createCertificate handler
}

func (s *Server) deleteCertificate(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement deleteCertificate handler
}

func (s *Server) downloadSecretData(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement downloadSecretData handler
}

func main() {
	server := NewServer()
	server.InitializeRoutes()

	err := server.Start(":8080")
	if err != nil {
		log.Fatal(err)
	}
}