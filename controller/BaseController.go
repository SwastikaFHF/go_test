package controller

import (
	"io"
	"log"
	"net/http"
	"time"
)

type Controller struct {
	Ctx Context
}

type Context struct {
	Rep http.ResponseWriter
	Req *http.Request
}

func (ctx *Context) WriteString(str string) {
	io.WriteString(ctx.Rep, str)
}

var mux map[string]*Controller

func init() {
	mux = make(map[string]*Controller)
}

//设置路由
func Router(url string, controller *Controller) {
	mux[url] = controller
}

//开始运行监听
func Run() {
	server := http.Server{
		Addr:        ":8080",
		Handler:     &Controller{},
		ReadTimeout: 5 * time.Second,
	}
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

func (controller *Controller) ServeHTTP(rep http.ResponseWriter, res *http.Request) {
	controller.Ctx.Rep = rep
	controller.Ctx.Req = res
	if h, ok := mux[res.URL.String()]; ok {
		h.DealHttp(rep, res)
		return
	}
	controller.Ctx.WriteString("Not found the right controller")
}

func (this *Controller) DealHttp(res http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		this.Get()
	} else if req.Method == "POST" {
		this.Post()
	} else {
		this.Ctx.WriteString("No matched http type")
	}

}

func (this *Controller) Post() {

}

func (this *Controller) Get() {

}
