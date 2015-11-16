package model

import (
	"go_test/model/request"
)

type Request struct {
	header     request.RequestHeader
	clientInfo request.ClientInfo
}
