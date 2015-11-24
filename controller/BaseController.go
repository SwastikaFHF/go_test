package controller

import (
	"io"
	"log"
	"net/http"
	"time"
)

type Controller struct {
	Ctx Context
	HttpMethod
}

type Context struct {
	Rep http.ResponseWriter
	Req *http.Request
}

func (ctx *Context) WriteString(str string) {
	io.WriteString(ctx.Rep, str)
}

var mux map[string]DealHttpInterface

func init() {
	mux = make(map[string]DealHttpInterface)
}

//设置路由
func Router(url string, controller DealHttpInterface) {
	log.Println("map")
	mux[url] = controller
}

//开始运行监听
func Run() {
	log.Println("run")
	server := http.Server{
		Addr:        ":8080",
		Handler:     &Controller{},
		ReadTimeout: 5 * time.Second,
	}
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("server is running now!!!")
	}
}

func (controller *Controller) ServeHTTP(rep http.ResponseWriter, res *http.Request) {
	controller.Ctx.Rep = rep
	controller.Ctx.Req = res
	if _, ok := mux[res.URL.String()]; ok {
		controller.DealHttp(rep, res)
		log.Println("DealHttp")
		return
	}
	controller.Ctx.WriteString("Not found the right controller")
}

type DealHttpInterface interface {
	DealHttp(res http.ResponseWriter, req *http.Request)
}

func (this *Controller) DealHttp(res http.ResponseWriter, req *http.Request) {
	log.Println("DealHttp111")
	log.Println(req.Method)
	if req.Method == "GET" {
		this.HttpMethod.Get()
	} else if req.Method == "POST" {
		this.HttpMethod.Post()
	} else {
		this.Ctx.WriteString("No matched http type")
	}

}

type HttpMethod interface {
	Post()
	Get()
}
