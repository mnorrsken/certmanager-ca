package k8s

import (
	"context"
	"testing"

	certmanagerv1 "github.com/cert-manager/cert-manager/pkg/apis/certmanager/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/dynamic/fake"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Client", func() {
	var (
		client    *fake.FakeDynamicClient
		namespace string
		name      string
		c         *Client
	)

	BeforeEach(func() {
		client = fake.NewSimpleDynamicClient(runtime.NewScheme())
		namespace = "default"
		name = "test-certificate"
		c = &Client{
			client: client,
		}
	})

	Context("CreateCertificate", func() {
		It("should create a certificate successfully", func() {
			err := c.CreateCertificate(context.Background(), namespace, name)
			Expect(err).To(BeNil())

			unstruct, err := client.Resource(certmanagerv1.SchemeGroupVersion.WithResource("certificates")).Namespace(namespace).Get(context.Background(), name, metav1.GetOptions{})
			Expect(err).To(BeNil())

			cer := &certmanagerv1.Certificate{}
			err = runtime.DefaultUnstructuredConverter.FromUnstructured(unstruct.Object, cer)

			By("checking if the Certificate resource was created successfully")
			Expect(err).To(BeNil())

			Expect(cer.Name).To(Equal(name))
			Expect(cer.Spec.CommonName).To(Equal("example.com"))

		})
	})
})

func TestSuite(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Client")
}
