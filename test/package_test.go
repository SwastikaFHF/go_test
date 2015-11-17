package test

import (
	"encoding/json"
	"fmt"
	"go_test/model"
	// "strings"
	"testing"
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

func TestDomain(t *testing.T) {
	// jsonStr := `{"header": {"digitalSign":"ceee"}}`

	//json to jsonString
	// if b, err := json.Marshal(jsonStr); err == nil {
	// 	fmt.Println(string(b))
	// }

	//json to struct
	// var config Request = make(Request("{}"))
	// if err := json.Unmarshal([]byte(jsonStr), &config); err == nil {
	// 	fmt.Println(config.Header.DigitalSign)
	// }

	var response Response
	response.Header.RspCode, response.Header.RspMsg = "12", "ccc"
	// response := Response{"header": {"rspCode": "ceee"}}
	// fmt.Println(response.Header.RspCode)
	if b, err := json.Marshal(response); err == nil {
		fmt.Println(string(b))
	}
}
