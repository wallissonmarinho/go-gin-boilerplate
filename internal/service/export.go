package service

import (
	"github.com/go-kit/log"
	"github.com/jmoiron/sqlx"
	"github.com/wallissonmarinho/go-gin-boilerplate/internal/client"
)

type ServiceFactory interface {
}

type serviceFactory struct {
}

func NewServiceFactory(db *sqlx.DB, logger log.Logger, clients client.ClientService) ServiceFactory {
	return &serviceFactory{}
}
