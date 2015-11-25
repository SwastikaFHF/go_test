package core

import (
	"log"
	"net/http"
	"time"
)

type Controller struct {
}

var mux map[string]DealHttpInterface

func init() {
	mux = make(map[string]DealHttpInterface)
}

//设置路由
func Router(url string, controller DealHttpInterface) {
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
	}
}

func (this *Controller) ServeHTTP(rep http.ResponseWriter, res *http.Request) {
	if h, ok := mux[res.URL.String()]; ok {
		switch res.Method {
		case "GET":
			h.Get(rep, res)
		case "POST":
			h.Post(rep, res)
		default:
			rep.Write([]byte("No matched http type"))
		}

	} else {
		rep.Write([]byte("Not found the right controller"))
	}
}

type DealHttpInterface interface {
	Post(rep http.ResponseWriter, res *http.Request)
	Get(rep http.ResponseWriter, res *http.Request)
}

func (this *Controller) Get(rep http.ResponseWriter, res *http.Request) {

}

func (this *Controller) Post(rep http.ResponseWriter, res *http.Request) {

}
