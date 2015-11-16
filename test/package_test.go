package test

import (
	"encoding/json"
	"fmt"
	"go_test/model"
	"testing"
)

type Request struct {
	model.Request
}

type Response struct {
}

func TestDomain(t *testing.T) {
	jsonStr := `{"header": {"digitalSign":"ceee"}}`

	//json to jsonString
	if b, err := json.Marshal(jsonStr); err == nil {
		fmt.Println(string(b))
	}

	//json to struct
	// var config Man
	// if err := json.Unmarshal([]byte(jsonStr), &config); err == nil {
	// 	fmt.Println(strings(config))
	// 	fmt.Println(config.Name)
	// }
}
