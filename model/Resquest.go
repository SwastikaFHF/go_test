package model

import (
	"go_personal_blog/model/request"
)

type Request struct {
	Header     request.RequestHeader `json:"header"`
	ClientInfo request.ClientInfo    `json:"clientInfo"`
}
