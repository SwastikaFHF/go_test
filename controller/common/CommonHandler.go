package common

import (
	"fmt"
	"net/http"
)

func DoPost(res http.ResponseWriter, req *http.Request) {

	fmt.Fprint(res, "CommonHandler start")
}
