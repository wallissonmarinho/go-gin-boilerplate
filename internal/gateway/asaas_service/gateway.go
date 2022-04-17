package asaasservice

import (
	"context"
	"errors"

	"github.com/wallissonmarinho/go-gin-boilerplate/internal/client"
	"github.com/wallissonmarinho/go-gin-boilerplate/internal/domain"
	"github.com/wallissonmarinho/go-gin-boilerplate/internal/helper"
	"gopkg.in/guregu/null.v4"
)

type Service interface {
	ObterClientePorCPFCNPJ(ctx context.Context, cpfCnpj null.String) (domain.AsaasCustomers, error)
	ObterCobrancasPendentePorClienteIDEDataInicioEDataFim(ctx context.Context, clienteID, dataInicio, dataFim null.String) (domain.AsaasPayments, error)
	ObterCobrancasVencidaPorClienteIDEDataInicioEDataFim(ctx context.Context, clienteID, dataInicio, dataFim null.String) (domain.AsaasPayments, error)
	ObterLinhaDigitavelPorBoletoID(ctx context.Context, boletoID null.String) (domain.AsaasLinhaDigitavel, error)
}

type asaasGateway struct {
	cli client.ClientService
}

func NewAsaasGateWay(cli client.ClientService) Service {
	return &asaasGateway{
		cli: cli,
	}
}

func (s *asaasGateway) ObterClientePorCPFCNPJ(ctx context.Context, cpfCnpj null.String) (domain.AsaasCustomers, error) {
	customers := domain.AsaasCustomers{}

	resp, err := s.cli.AsaasServiceWithRest().
		SetHeaders(helper.FromHeader(ctx)).
		SetResult(&customers).
		SetQueryParams(map[string]string{
			"cpfCnpj": cpfCnpj.String,
		}).
		Get("/v3/customers")

	if resp.StatusCode() != 200 {
		return domain.AsaasCustomers{}, errors.New("Não foi possível obter o cliente")
	}

	return customers, err
}

func (s *asaasGateway) ObterCobrancasPendentePorClienteIDEDataInicioEDataFim(ctx context.Context, clienteID, dataInicio, dataFim null.String) (domain.AsaasPayments, error) {
	asaasPayments := domain.AsaasPayments{}

	resp, err := s.cli.AsaasServiceWithRest().
		SetHeaders(helper.FromHeader(ctx)).
		SetResult(&asaasPayments).
		SetQueryParams(map[string]string{
			"customer":    clienteID.String,
			"status":      "PENDING",
			"dueDate[ge]": dataInicio.String,
			"dueDate[le]": dataFim.String,
		}).
		Get("/v3/payments")

	if resp.StatusCode() != 200 {
		return domain.AsaasPayments{}, errors.New("Não foi possível obter o cobrança")
	}

	return asaasPayments, err
}

func (s *asaasGateway) ObterCobrancasVencidaPorClienteIDEDataInicioEDataFim(ctx context.Context, clienteID, dataInicio, dataFim null.String) (domain.AsaasPayments, error) {
	asaasPayments := domain.AsaasPayments{}

	resp, err := s.cli.AsaasServiceWithRest().
		SetHeaders(helper.FromHeader(ctx)).
		SetResult(&asaasPayments).
		SetQueryParams(map[string]string{
			"customer":    clienteID.String,
			"status":      "OVERDUE",
			"dueDate[ge]": dataInicio.String,
			"dueDate[le]": dataFim.String,
		}).
		Get("/v3/payments")

	if resp.StatusCode() != 200 {
		return domain.AsaasPayments{}, errors.New("Não foi possível obter o cobrança")
	}

	return asaasPayments, err
}

func (s *asaasGateway) ObterLinhaDigitavelPorBoletoID(ctx context.Context, boletoID null.String) (domain.AsaasLinhaDigitavel, error) {
	asaasLinhaDigitavel := domain.AsaasLinhaDigitavel{}

	resp, err := s.cli.AsaasServiceWithRest().
		SetHeaders(helper.FromHeader(ctx)).
		SetResult(&asaasLinhaDigitavel).
		SetPathParams(map[string]string{
			"id": boletoID.String,
		}).
		Get("/v3/payments/{id}/identificationField")

	if resp.StatusCode() != 200 {
		return domain.AsaasLinhaDigitavel{}, errors.New("Não foi possível obter linha digitável")
	}

	return asaasLinhaDigitavel, err
}
