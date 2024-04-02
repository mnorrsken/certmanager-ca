package k8s

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/pkg/errors"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

// Client represents the Kubernetes client.
type Client struct {
	clientset *kubernetes.Clientset
}

// NewClient creates a new Kubernetes client.
func NewClient() (*Client, error) {
	kubeconfig := filepath.Join(os.Getenv("HOME"), ".kube", "config")
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		return nil, errors.Wrap(err, "failed to build kubeconfig")
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create clientset")
	}

	return &Client{
		clientset: clientset,
	}, nil
}

// CreateCertificate creates a cert-manager Certificate resource.
func (c *Client) CreateCertificate(namespace, name string) error {
	certificate := &certmanagerv1.Certificate{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: namespace,
			Name:      name,
		},
		Spec: certmanagerv1.CertificateSpec{
			// Add the certificate spec here
		},
	}

	_, err := c.clientset.CertmanagerV1().Certificates(namespace).Create(context.TODO(), certificate, metav1.CreateOptions{})
	if err != nil {
		return errors.Wrap(err, "failed to create Certificate resource")
	}

	return nil
}

// DeleteCertificate deletes a cert-manager Certificate resource.
func (c *Client) DeleteCertificate(namespace, name string) error {
	err := c.clientset.CertmanagerV1().Certificates(namespace).Delete(context.TODO(), name, metav1.DeleteOptions{})
	if err != nil {
		if errors.IsNotFound(err) {
			return fmt.Errorf("Certificate resource not found")
		}
		return errors.Wrap(err, "failed to delete Certificate resource")
	}

	return nil
}

// DownloadSecretData downloads data from the created TLS secrets.
func (c *Client) DownloadSecretData(namespace, name string) ([]byte, error) {
	secret, err := c.clientset.CoreV1().Secrets(namespace).Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		if errors.IsNotFound(err) {
			return nil, fmt.Errorf("Secret not found")
		}
		return nil, errors.Wrap(err, "failed to get Secret")
	}

	data, ok := secret.Data[corev1.TLSCertKey]
	if !ok {
		return nil, fmt.Errorf("TLS certificate data not found in Secret")
	}

	return data, nil
}