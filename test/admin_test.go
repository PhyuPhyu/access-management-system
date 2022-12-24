package test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdminLogin(t *testing.T) {
	r := ConnectDBForTest()

	for _, testData := range AdminLoginTestData {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPost, "/v1/ams/admin/login", bytes.NewBuffer([]byte(testData.RequestData)))
		r.ServeHTTP(w, req)

		assert.Equal(t, testData.ExpectedCode, w.Code)
	}
}

func TestGetAllUsers(t *testing.T) {
	r := ConnectDBForTest()

	for _, testData := range AdminRoutesTestData {
		w := httptest.NewRecorder()

		req, _ := http.NewRequest(http.MethodGet, "/v1/ams/admin/users", nil)
		testData.Request(req)
		r.ServeHTTP(w, req)

		assert.Equal(t, testData.ExpectedCode, w.Code)
		if testData.ExpectedCode != 200 {
			assert.Equal(t, testData.ExpectedResponse, w.Body.String())
		}
	}
}

func TestGetUser(t *testing.T) {
	r := ConnectDBForTest()

	for _, testData := range AdminRoutesTestData {
		w := httptest.NewRecorder()

		userId := testData.RequestData

		req, _ := http.NewRequest(http.MethodGet, "/v1/ams/admin/users/"+userId, nil)
		testData.Request(req)
		r.ServeHTTP(w, req)

		assert.Equal(t, testData.ExpectedCode, w.Code)
		if testData.ExpectedCode != 200 {
			assert.Equal(t, testData.ExpectedResponse, w.Body.String())
		}
	}
}

func TestUserAccountApprovedByAdmin(t *testing.T) {
	r := ConnectDBForTest()

	for _, testData := range AdminRoutesTestData {
		w := httptest.NewRecorder()

		userId := testData.RequestData

		req, _ := http.NewRequest(http.MethodPut, "/v1/ams/admin/users/"+userId+"/approve", nil)
		testData.Request(req)
		r.ServeHTTP(w, req)

		assert.Equal(t, testData.ExpectedCode, w.Code)
		if testData.ExpectedCode != 200 {
			assert.Equal(t, testData.ExpectedResponse, w.Body.String())
		}
	}
}

func TestDeleteUser(t *testing.T) {
	r := ConnectDBForTest()

	for _, testData := range AdminRoutesTestData {
		w := httptest.NewRecorder()

		userId := testData.RequestData

		req, _ := http.NewRequest(http.MethodDelete, "/v1/ams/admin/users/"+userId, nil)
		testData.Request(req)
		r.ServeHTTP(w, req)

		assert.Equal(t, testData.ExpectedCode, w.Code)
		if testData.ExpectedCode != 200 {
			assert.Equal(t, testData.ExpectedResponse, w.Body.String())
		}
	}
}
