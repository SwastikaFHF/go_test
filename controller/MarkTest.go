package controller

import (
	"github.com/russross/blackfriday"
	"go_test/core"
	"io/ioutil"
)

type MarkTest struct {
	core.Controller
}

func (this *MarkTest) Get() {
	buf, err := ioutil.ReadFile("view/首页.md")
	if err != nil {
	}
	output := blackfriday.MarkdownCommon(buf)
	this.Ctx.WriteByte(output)
}
