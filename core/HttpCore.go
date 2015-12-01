package core

import (
	"log"
	"net/http"
	"path"
	"strings"
	"time"
)

var mux map[string]DealHttpInterface
var staticDir map[string]string

func init() {
	mux = make(map[string]DealHttpInterface)
	staticDir = make(map[string]string)
}

//设置路由
func Router(url string, controller DealHttpInterface) {
	mux[url] = controller
}

//设置静态文件目录
func SetStaticPath(urlPath string, staticPath string) {
	staticDir[urlPath] = staticPath
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
	//开始处理静态文件
	requestPath := path.Clean(res.URL.Path)
	sli := strings.Split(requestPath, "/")
	prefix := "/" + sli[1]
	if localdir, ok := staticDir[prefix]; ok {
		file := path.Join(localdir, requestPath[len(prefix):])
		http.ServeFile(rep, res, file)
		return
	}

	//开始处理http请求
	if h, ok := mux[res.URL.String()]; ok {
		ctx := &Context{
			ResponseWriter: rep,
			Request:        res,
		}
		h.Init(ctx)
		switch res.Method {
		case "GET":
			h.Get()
			h.HandleTpl()
		case "POST":
			h.Post()
			h.HandleTpl()
		default:
			rep.Write([]byte("No matched http type"))
		}

	} else {
		rep.Write([]byte("Not found the right controller"))
	}
}
