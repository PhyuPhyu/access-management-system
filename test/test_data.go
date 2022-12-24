package test

import (
	"access-management-system/models"
	"access-management-system/utils"
	"fmt"
	"net/http"

	"github.com/outbrain/golib/log"
)

type RequestTest struct {
	RequestData      string
	ExpectedCode     int
	ExpectedResponse string
}

type AuthRequestTest struct {
	Request          func(*http.Request)
	RequestData      string
	ExpectedCode     int
	ExpectedResponse string
}

// -------------------------------------Testing for user register ------------------------------------------------------------
var RegisterTestData = []RequestTest{
	{
		`{"name": "Phyu Phyu","email": "phyuphyuthin24hha.dev@gmail.com","password": "password"}`,
		http.StatusCreated,
		`{"message":"Email sent successfully to phyuphyuthin24hha.dev@gmail.com"}`,
	},
	{
		`{"name": "Phyu Phyu","email": "phyuphyuthin24hha.dev@gmail.com","password": "password"}`,
		http.StatusBadRequest,
		`{"error":"Error 1062 (23000): Duplicate entry 'phyuphyuthin24hha.dev@gmail.com' for key 'users.email'"}`,
	},
	{
		`{"name": "Phyu Phyu","email": "phyuphyuthin24hha.dev@gmail.com","password": ""}`,
		http.StatusBadRequest,
		`{"error":"Password should not be empty!"}`,
	},
}

// -------------------------------------Testint for verify email ------------------------------------------------------------
var VerifyEmailTestData = []RequestTest{
	{
		`c46SWqQ9PN6jVZSDUlfS`,
		http.StatusOK,
		`{"message":"Email verified successfully"}`,
	},
	{
		`c46SWqQ9PN6jVZSDUlfS`,
		http.StatusBadRequest,
		`{"error":"record not found"}`,
	},
}

// -------------------------------------Testing for user login ------------------------------------------------------------
var UserLoginTestData = []RequestTest{
	{
		`{"email": "phyuphyuthin24hha.dev@gmail.com","password": "password"}`,
		http.StatusOK,
		`{"message":"User login successfully","token":"([a-zA-Z0-9-_.]{115})"}`,
	},
	{
		`{"email": "phyuphyuthin24hha1.dev@gmail.com","password": "password"}`,
		http.StatusBadRequest,
		`{"error":"Input email is not registered yet"}`,
	},
	{
		`{"email": "phyuphyuthin24hha.dev@gmail.com","password": "password"}`,
		http.StatusBadRequest,
		`{"error":"Invalid password"}`,
	},
}

// User token create and set for user
func HeaderUserTokenMock(req *http.Request, userId int) {
	var user models.User
	err := models.GetUser(&user, userId)
	if err != nil {
		log.Error("Get user by id error: ", err.Error())
	}

	token, err := utils.CreateToken(user)
	if err != nil {
		log.Error("Create user token error: ", err.Error())
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", token))
}

// -------------------------------------Testing for get me ------------------------------------------------------------
var GetMeTestData = []AuthRequestTest{
	{
		func(req *http.Request) {
			HeaderUserTokenMock(req, 1)
		},
		``,
		http.StatusOK,
		`{"id":1,"name":"Phyu Phyu","email":"phyuphyuthin24hha.dev@gmail.com"}`,
	},
	{
		func(req *http.Request) {
		},
		``,
		http.StatusUnauthorized,
		`{"error":"Authorization token required"}`,
	},
}

// -------------------------------------Testing for admin login ------------------------------------------------------------
var AdminLoginTestData = []RequestTest{
	{
		`{"name": "admin1","password": "password"}`,
		http.StatusOK,
		`{"message":"Admin login successfully","token":"([a-zA-Z0-9-_.]{115})"}`,
	},
	{
		`{"name": "admin1","password": "passwor"}`,
		http.StatusBadRequest,
		`{"error":"Invalid password"}`,
	},
	{
		`{"name": "admin4","password": "password"}`,
		http.StatusBadRequest,
		`{"error":"record not found"}`,
	},
}

// Admin token create and set for user
func HeaderAdminTokenMock(req *http.Request, adminName string) {
	var admin models.Admin
	err := models.GetAdmin(&admin, adminName)
	if err != nil {
		log.Error("Get admin by name error: ", err.Error())
	}

	token, err := utils.CreateAdminToken(admin)
	if err != nil {
		log.Error("Create admin token error: ", err.Error())
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", token))
}

// -------------------------------------Testing for admin routes------------------------------------------------------------
var AdminRoutesTestData = []AuthRequestTest{
	{
		func(req *http.Request) {
			HeaderAdminTokenMock(req, "admin1")
		},
		"1",
		http.StatusOK,
		``,
	},
	{
		func(req *http.Request) {
		},
		"1",
		http.StatusUnauthorized,
		`{"error":"Authorization token required"}`,
	},
	{
		func(req *http.Request) {
			HeaderAdminTokenMock(req, "admin2")
		},
		"1",
		http.StatusUnauthorized,
		`{"error":"Unauthorized access"}`,
	},
}
