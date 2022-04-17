package client

import (
	"os"

	"github.com/go-resty/resty/v2"
)

type ClientService interface {
	AsaasServiceWithRest() *resty.Request
}

type clientService struct {
}

func NewClients() *clientService {
	return &clientService{}
}

func (c *clientService) AsaasServiceWithRest() *resty.Request {
	apiToken := os.Getenv("ASAAS_API_TOKEN")
	cli := resty.New()
	cli.SetBaseURL("https://www.asaas.com/api")
	cli.SetHeader("access_token", apiToken)
	return cli.R()
}
