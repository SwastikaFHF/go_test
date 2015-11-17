package model

import (
	"go_test/model/request"
)

type Request struct {
	Header     request.RequestHeader `json:"header"`
	ClientInfo request.ClientInfo    `json:"clientInfo"`
}
