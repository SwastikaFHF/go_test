package core

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
)

type Controller struct {
	Ctx      *Context
	TplNames string
	Data     map[interface{}]interface{}
	Json     interface{}
}

type DealHttpInterface interface {
	Init(context *Context)
	Post()
	Get()
	HandleTpl()
}

func (this *Controller) HandleTpl() {
	if this.TplNames != "" {
		t, TempErr := template.ParseFiles(this.TplNames)
		if TempErr == nil {
			err := t.Execute(this.Ctx.ResponseWriter, this.Data)
			if err != nil {
				log.Println(err)
			}
		} else {
			log.Println(TempErr)
		}
	} else if this.Json != nil {
		if b, err := json.Marshal(this.Json); err == nil {
			this.Ctx.WriteByte(b)
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
