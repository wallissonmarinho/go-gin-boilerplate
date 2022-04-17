package endpoint

import (
	"context"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/wallissonmarinho/go-gin-boilerplate/internal/domain"
	"github.com/wallissonmarinho/go-gin-boilerplate/internal/service"
	"gopkg.in/guregu/null.v4"
)

// makeHealthEndpoint return if service up
func makeHealthEndpoint(s service.ServiceFactory, logger log.Logger) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {

		_ = level.Error(logger).Log("message", "ok")

		return domain.CustomerResponse{
			Code:     null.IntFrom(http.StatusOK),
			Response: "Service UP",
		}, nil
	}
}
