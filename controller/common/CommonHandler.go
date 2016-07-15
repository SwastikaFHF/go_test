package common

import (
	"go_personal_blog/core"
	"go_personal_blog/model"
)

type Common struct {
	core.Controller
}

type Request struct {
	model.Request
}

type Response struct {
	model.Response
	body ResponseBody
}
type ResponseBody struct {
	url string
}

func (this *Common) Get() {
	var respStr Response
	respStr.Header.RspCode, respStr.Header.RspMsg, respStr.body.url = "0000", "获取数据成功", "htttp://www.baidu.com"
	this.Json = respStr
}
