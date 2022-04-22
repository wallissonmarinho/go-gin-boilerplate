package service

import (
	"github.com/go-kit/log"
	"github.com/jmoiron/sqlx"
)

type ServiceFactory interface {
}

type serviceFactory struct {
}

func NewServiceFactory(db *sqlx.DB, logger log.Logger) ServiceFactory {
	return &serviceFactory{}
}
