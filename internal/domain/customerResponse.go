package domain

import "gopkg.in/guregu/null.v4"

type CustomerResponse struct {
	Code     null.Int    `json:"code"`
	Response interface{} `json:"response"`
}
