package domain

import "gopkg.in/guregu/null.v4"

type AsaasCustomers struct {
	Object     null.String     `json:"object"`
	HasMore    null.Bool       `json:"hasMore"`
	TotalCount null.Int        `json:"totalCount"`
	Limit      null.Int        `json:"limit"`
	Offset     null.Int        `json:"offset"`
	Data       []AsaasCustomer `json:"data"`
}

type AsaasCustomer struct {
	Object            null.String `json:"object"`
	ID                null.String `json:"id"`
	DateCreated       null.String `json:"dateCreated"`
	Name              null.String `json:"name"`
	Email             null.String `json:"email"`
	Phone             null.String `json:"phone"`
	MobilePhone       null.String `json:"mobilePhone"`
	Address           null.String `json:"address"`
	AddressNumber     null.String `json:"addressNumber"`
	Complement        null.String `json:"complement"`
	Province          null.String `json:"province"`
	PostalCode        null.String `json:"postalCode"`
	CpfCnpj           null.String `json:"cpfCnpj"`
	PersonType        null.String `json:"personType"`
	AdditionalEmails  null.String `json:"additionalEmails"`
	ExternalReference null.String `json:"externalReference"`
	Observations      null.String `json:"observations"`
	State             null.String `json:"state"`
	Country           null.String `json:"country"`
}

type AsaasPayments struct {
	Object     null.String    `json:"object"`
	HasMore    null.Bool      `json:"hasMore"`
	TotalCount null.Int       `json:"totalCount"`
	Limit      null.Int       `json:"limit"`
	Offset     null.Int       `json:"offset"`
	Data       []AsaasPayment `json:"data"`
}

type AsaasPayment struct {
	Object          null.String `json:"object"`
	ID              null.String `json:"id"`
	DateCreated     null.String `json:"dateCreated"`
	CustomerID      null.String `json:"customer"`
	Value           null.Float  `json:"value"`
	NetValue        null.Float  `json:"netValue"`
	Description     null.String `json:"description"`
	BillingType     null.String `json:"billingType"`
	Status          null.String `json:"status"`
	DueDate         null.String `json:"dueDate"`
	OriginalDueDate null.String `json:"originalDueDate"`
	InvoiceUrl      null.String `json:"invoiceUrl"`
	InvoiceNumber   null.String `json:"invoiceNumber"`
	Deleted         null.Bool   `json:"deleted"`
	Anticipated     null.Bool   `json:"anticipated"`
	NossoNumero     null.String `json:"nossoNumero"`
	BankSlipUrl     null.String `json:"bankSlipUrl"`
}

type AsaasLinhaDigitavel struct {
	IdentificationField null.String `json:"identificationField"`
	NossoNumero         null.String `json:"nossoNumero"`
}
