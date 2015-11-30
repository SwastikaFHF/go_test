package common

import (
	"fmt"
	"go_test/core"
	"net/http"
	"strings"
)

type Image struct {
	core.Controller
}

var staticDir map[string]string

func (this *Image) Get() {
	staticDir = make(map[string]string)
	staticDir["/img/.*"] = "../go_test/static/icon"
	sli := strings.Split(this.Ctx.Request.URL.Path, "/")
	prefix := "/" + sli[1]                     //find webdir prefix such as "/asset"
	if localdir, ok := staticDir[prefix]; ok { //assertion to find localdir
		fmt.Printf("\n******the prefix is:%s  the localdir is:%s\n\n", prefix, localdir)
		file := localdir + this.Ctx.Request.URL.Path[len(prefix):]
		http.ServeFile(this.Ctx.ResponseWriter, this.Ctx.Request, file) //return local resource as static resource
		return
	}
}
