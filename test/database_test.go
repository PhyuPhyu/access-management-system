package test

import (
	"access-management-system/config"
	"access-management-system/models"
	"access-management-system/routers"
	"log"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestConnectDB(t *testing.T) {
	// Set Gin to Test Mode
	gin.SetMode(gin.TestMode)

	config, err := config.LoadConfig("../.")
	if err != nil {
		assert.Error(t, err, "Load config file error")
	}

	models.ConnectDB(&config)
}

func ConnectDBForTest() *gin.Engine {
	// Set Gin to Test Mode
	gin.SetMode(gin.TestMode)

	config, err := config.LoadConfig("../.")
	if err != nil {
		log.Fatal("🚀 Could not load environment variables", err)
	}

	models.ConnectDB(&config)

	r := routers.Router()

	return r
}
