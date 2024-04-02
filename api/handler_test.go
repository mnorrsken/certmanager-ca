package api

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Handler", func() {
	var (
		router *gin.Engine
		w      *httptest.ResponseRecorder
	)

	BeforeEach(func() {
		gin.SetMode(gin.TestMode)
		router = SetupRouter()
		w = httptest.NewRecorder()
	})

	Describe("CreateCertificate", func() {
		Context("with valid certificate data", func() {
			It("should return a 201 status code", func() {
				req, _ := http.NewRequest("POST", "/certificates", strings.NewReader(`{"name":"test","hostnames":["test.com"]}`))
				req.Header.Set("Content-Type", "application/json")
				router.ServeHTTP(w, req)

				Expect(w.Code).To(Equal(http.StatusCreated))
			})
		})

		Context("with invalid certificate data", func() {
			It("should return a 400 status code", func() {
				req, _ := http.NewRequest("POST", "/certificates", strings.NewReader(`{"name":"test","hostnames":"test.com"}`)) // invalid hostnames field
				req.Header.Set("Content-Type", "application/json")
				router.ServeHTTP(w, req)

				Expect(w.Code).To(Equal(http.StatusBadRequest))
			})
		})
	})
})

func TestSuite(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Client")
}
