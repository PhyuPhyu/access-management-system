package test

import (
	"access-management-system/models"
	"access-management-system/routers"
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestHealthCheck(t *testing.T) {
	// Set Gin to Test Mode
	gin.SetMode(gin.TestMode)

	r := routers.Router()

	w := httptest.NewRecorder()

	expectedResponse := `{"message":"Server is up and running"}`

	req, _ := http.NewRequest(http.MethodGet, "/v1/ams/health-check", nil)
	r.ServeHTTP(w, req)

	// Assertion
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, expectedResponse, w.Body.String())
}

func TestUserRegister(t *testing.T) {
	r := ConnectDBForTest()
	models.MigrateUser()

	for _, testData := range RegisterTestData {
		req, _ := http.NewRequest(http.MethodPost, "/v1/ams/users/register", bytes.NewBuffer([]byte(testData.RequestData)))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, testData.ExpectedCode, w.Code)
		assert.Equal(t, testData.ExpectedResponse, w.Body.String())
	}
}

func TestVerifyEmail(t *testing.T) {
	r := ConnectDBForTest()

	for _, testData := range VerifyEmailTestData {
		// Get a verify code from email when registration
		verificationCode := testData.RequestData

		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/v1/ams/users/verify-email/"+verificationCode, nil)
		r.ServeHTTP(w, req)

		assert.Equal(t, testData.ExpectedCode, w.Code)
		assert.Equal(t, testData.ExpectedResponse, w.Body.String())
	}
}

func TestUserLogin(t *testing.T) {
	r := ConnectDBForTest()

	for _, testData := range UserLoginTestData {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPost, "/v1/ams/users/login", bytes.NewBuffer([]byte(testData.RequestData)))
		r.ServeHTTP(w, req)

		assert.Equal(t, testData.ExpectedCode, w.Code)
	}
}

func TestGetMe(t *testing.T) {
	r := ConnectDBForTest()

	for _, testData := range GetMeTestData {
		w := httptest.NewRecorder()

		req, _ := http.NewRequest(http.MethodGet, "/v1/ams/users/me", nil)
		testData.Request(req)
		r.ServeHTTP(w, req)

		assert.Equal(t, testData.ExpectedCode, w.Code)
		assert.Equal(t, testData.ExpectedResponse, w.Body.String())
	}
}
