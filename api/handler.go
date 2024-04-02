package api

import (
	"net/http"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"

	"github.com/gin-gonic/gin"
)

type Certificate struct {
	Name      string   `json:"name"`
	Hostnames []string `json:"hostnames"`
}

var certificates []Certificate

func initKubernetesClient() (*kubernetes.Clientset, error) {
	// Path to the kubeconfig file
	kubeconfig := "/path/to/kubeconfig"

	// Build the client configuration
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		return nil, err
	}

	// Create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}

	return clientset, nil
}

func CreateCertificate(c *gin.Context) {
	var certificate Certificate

	if err := c.ShouldBindJSON(&certificate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// create a new certificate in kubernetes

	certificates = append(certificates, certificate)

	c.JSON(http.StatusCreated, certificate)
}

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.POST("/certificates", CreateCertificate)

	return r
}
