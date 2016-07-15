package model

import (
	"go_personal_blog/model/response"
)

type Response struct {
	Header response.ResponseHeader `json:"header"`
}
