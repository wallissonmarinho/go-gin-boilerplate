package domain

import "gopkg.in/guregu/null.v4"

type ZintechRequest struct {
	ID        null.Int        `json:"id"`
	ClienteID null.Int        `json:"clienteId"`
	Text      null.String     `json:"text"`
	Contact   ZintechContacts `json:"contact"`
	Data      interface{}     `json:"data"`
}

type ZintechResponse struct {
	Type        null.String         `json:"type"`
	Text        null.String         `json:"text"`
	Attachments []ZintechAttachment `json:"attachments"`
}

type ZintechContacts struct {
	UID      null.Int        `json:"uid"`
	Type     null.String     `json:"type"`
	Key      null.String     `json:"key"`
	Name     null.String     `json:"name"`
	Fields   ZintechFields   `json:"fields"`
	Callback ZintechCallback `json:"callback"`
}

type ZintechFields struct {
	CPF     null.String `json:"cpf"`
	CNPJ    null.String `json:"cnpj"`
	Celular null.String `json:"celular"`
}

type ZintechAttachment struct {
	Position null.String `json:"position"`
	Type     null.String `json:"type"`
	Name     null.String `json:"name"`
	Url      null.String `json:"url"`
}

type ZintechCallback struct {
	Endpoint null.String `json:"endpoint"`
	Data     interface{} `json:"data"`
}
