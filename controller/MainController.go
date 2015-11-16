package controller

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func Index(res http.ResponseWriter, req *http.Request) {
	url := req.URL.String()
	if strings.EqualFold("/", url) {
		contents, err := ioutil.ReadFile("./static/index.html")
		if err != nil {
			fmt.Fprintf(res, "404")
			return
		}
		fmt.Fprintf(res, "%s\n", contents)
	} else {
		fmt.Fprint(res, "404 not found:\n\n")
	}
}
