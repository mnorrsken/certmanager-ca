package k8s

import (
	"context"
	"os"
	"path/filepath"
	"time"

	certmanagerv1 "github.com/cert-manager/cert-manager/pkg/apis/certmanager/v1"
	cmmeta "github.com/cert-manager/cert-manager/pkg/apis/meta/v1"
	gerrors "github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/tools/clientcmd"
)

// Client represents the Kubernetes client.
type Client struct {
	client dynamic.Interface
}

// NewClient creates a new Kubernetes client.
func NewClient() (*Client, error) {
	kubeconfig := filepath.Join(os.Getenv("HOME"), ".kube", "config")
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		return nil, gerrors.Wrap(err, "failed to build kubeconfig")
	}

	client, err := dynamic.NewForConfig(config)
	if err != nil {
		return nil, gerrors.Wrap(err, "failed to create clientset")
	}

	return &Client{
		client: client,
	}, nil
}

func (c *Client) toUnstructured(obj interface{}) (*unstructured.Unstructured, error) {
	converted, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, gerrors.Wrap(err, "failed to convert object to Unstructured")
	}

	return &unstructured.Unstructured{Object: converted}, nil
}

// CreateCertificate creates a cert-manager Certificate resource.
func (c *Client) CreateCertificate(ctx context.Context, namespace, name string) error {
	certificate := &certmanagerv1.Certificate{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: namespace,
			Name:      name,
		},
		Spec: certmanagerv1.CertificateSpec{
			SecretName: name + "-tls",
			IssuerRef: cmmeta.ObjectReference{
				Name: "ca",
				Kind: "ClusterIssuer",
			},
			CommonName:  "example.com",
			DNSNames:    []string{"example.com"},
			IPAddresses: []string{"10.0.0.1"},
			Duration: &metav1.Duration{
				Duration: 90 * 24 * time.Hour,
			},
		},
	}

	obj, err := c.toUnstructured(certificate)
	if err != nil {
		return err
	}

	_, err = c.client.Resource(certmanagerv1.SchemeGroupVersion.WithResource("certificates")).Namespace(namespace).Create(ctx, obj, metav1.CreateOptions{})
	if err != nil {
		return gerrors.Wrap(err, "failed to create Certificate resource")
	}
	return nil
}
