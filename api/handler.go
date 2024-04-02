package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

type Handler struct {
	clientset *kubernetes.Clientset
}

func NewHandler() (*Handler, error) {
	config, err := clientcmd.BuildConfigFromFlags("", clientcmd.RecommendedHomeFile)
	if err != nil {
		return nil, fmt.Errorf("failed to build Kubernetes config: %v", err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, fmt.Errorf("failed to create Kubernetes client: %v", err)
	}

	return &Handler{
		clientset: clientset,
	}, nil
}

func (h *Handler) CreateCertificate(w http.ResponseWriter, r *http.Request) {
	// Implement the logic to create a cert-manager Certificate resource
}

func (h *Handler) DeleteCertificate(w http.ResponseWriter, r *http.Request) {
	// Implement the logic to delete a cert-manager Certificate resource
}

func (h *Handler) DownloadSecretData(w http.ResponseWriter, r *http.Request) {
	// Implement the logic to download data from the created TLS secrets
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	router := mux.NewRouter()
	router.HandleFunc("/certificates", h.CreateCertificate).Methods("POST")
	router.HandleFunc("/certificates/{name}", h.DeleteCertificate).Methods("DELETE")
	router.HandleFunc("/secrets/{name}/download", h.DownloadSecretData).Methods("GET")

	router.ServeHTTP(w, r)
}