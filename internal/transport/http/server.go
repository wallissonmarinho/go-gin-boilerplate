package transport

import (
	"context"
	"crypto/sha256"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-kit/log"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/wallissonmarinho/go-gin-boilerplate/internal/endpoint"
	usersRepository "github.com/wallissonmarinho/go-gin-boilerplate/internal/repository/users"
	"gopkg.in/guregu/null.v4"
)

type server struct {
	usersRepository usersRepository.UsersRepository
	endpoint        *endpoint.Endpoints
	logger          *log.Logger
}

// NewService wires Go kit endpoints to the HTTP transport.
func NewService(context context.Context, db *sqlx.DB, endpoint *endpoint.Endpoints, logger *log.Logger) http.Handler {
	rest := &server{
		endpoint:        endpoint,
		logger:          logger,
		usersRepository: usersRepository.NewUsersRepository(db),
	}

	r := gin.New()

	r.Use(CorsMiddleware())

	err := r.SetTrustedProxies(strings.Split(os.Getenv("TrustedProxies"), ","))
	if err != nil {
		logrus.Error(err)
	}

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	v1 := r.Group("v1")
	v1.Use(validateAPIKey(rest))
	{
		v1.GET("/health", rest.HealthCheckHandler)
	}

	port := os.Getenv("PORT")

	err = r.Run(fmt.Sprintf("%s:%s", "0.0.0.0", port))

	logrus.Error(err)

	return r
}

func validateAPIKey(rest *server) gin.HandlerFunc {
	return func(c *gin.Context) {
		APIKey := c.Request.Header.Get("X-API-Key")

		keyHash := fmt.Sprintf("%s:%s", APIKey, viper.GetString("boilerplate_secret"))

		h := sha256.New()

		h.Write([]byte(keyHash))

		key := null.StringFrom(fmt.Sprintf("%x", h.Sum(nil)))

		validKey, err := rest.usersRepository.ValidApiKey(key)
		if err != nil || !validKey.Bool {
			fmt.Printf("Found 0 results for API Key: %s\n", APIKey)
			c.JSON(http.StatusUnauthorized, gin.H{"status": 401, "message": "Authentication failed"})
			c.Abort()
			return
		}
	}
}

func CorsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST, HEAD, PATCH, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
