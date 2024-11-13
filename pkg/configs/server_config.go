package configs

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/shou-nian/EzCashier/pkg/utils"
)

// ServerConfig func for configuration net/http app.
func ServerConfig(router *gin.Engine) *http.Server {
	// Define server settings:
	serverConnURL, _ := utils.ConnectionURLBuilder("server")
	readTimeoutSecondsCount, _ := strconv.Atoi(os.Getenv("SERVER_READ_TIMEOUT"))

	// Return server configuration.
	return &http.Server{
		Handler:     router,
		Addr:        serverConnURL,
		ReadTimeout: time.Second * time.Duration(readTimeoutSecondsCount),
	}
}
