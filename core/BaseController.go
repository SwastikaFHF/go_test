package core

import (
	// "html/template"
	"io/ioutil"
	"log"
	"net/http"
)

type Controller struct {
	Ctx      *Context
	TplNames string
	Data     map[interface{}]interface{}
}

type DealHttpInterface interface {
	Init(context *Context)
	Post()
	Get()
	HandleTpl(context *Context)
}

func (this *Controller) HandleTpl(ctx *Context) {
	if this.TplNames != "" {
		// t, _ := template.New("html").ParseFiles(this.TplNames)
		// // data := make(map[string]interface{})
		// // data["Website"] = "Website"
		// // data["Email"] = "Email"
		// t.ExecuteTemplate(os.Stdout, "html", this.Data)

		content, err := ioutil.ReadFile(this.TplNames)
		if err != nil {
			log.Println(err.Error())
			return
		} else {
			ctx.WriteByte(content)
		}
	}
}

func (this *Controller) Init(ctx *Context) {
	this.Ctx = ctx
	this.Data = make(map[interface{}]interface{})
}

func (this *Controller) Get() {
	http.Error(this.Ctx.ResponseWriter, "Method Not Allowed", 405)
}

func (this *Controller) Post() {
	http.Error(this.Ctx.ResponseWriter, "Method Not Allowed", 405)
}
