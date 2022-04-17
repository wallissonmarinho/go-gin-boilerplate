package transport

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/wallissonmarinho/go-gin-boilerplate/internal/domain"
)

func (s *server) HealthCheckHandler(c *gin.Context) {

	resp, err := s.endpoint.Health(c, nil)
	if err != nil {
		logrus.Error(err)

	}

	c.JSON(int(resp.(domain.CustomerResponse).Code.Int64), resp.(domain.CustomerResponse).Response)
}
