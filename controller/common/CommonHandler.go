package common

import (
	"encoding/json"
	"fmt"
	"go_test/model"
	"net/http"
)

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

func GetLoadImage(res http.ResponseWriter, req *http.Request) {
	var respStr Response
	respStr.Header.RspCode, respStr.Header.RspMsg, respStr.body.url = "0000", "获取数据成功", "htttp://www.baidu.com"
	if b, err := json.Marshal(respStr); err == nil {
		fmt.Fprint(res, string(b))
	}
	req.Form
}
