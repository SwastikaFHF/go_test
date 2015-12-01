package model

import (
	"go_test/model/response"
)

type Response struct {
	Header response.ResponseHeader `json:"header"`
}
