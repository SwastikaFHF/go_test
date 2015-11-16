package common

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type Request struct {
	RequestHeader RequestHeader `json:"header"`
}

type Response struct {
}

type RequestHeader struct {
}

func env(res http.ResponseWriter, req *http.Request) {
	fmt.Fprint(res, "System Environment:\n\n")
	env := os.Environ()
	for _, e := range env {
		fmt.Fprintln(res, e)
	}

}
