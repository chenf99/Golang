package service

import (
	"sync"
	"fmt"
	"github.com/codegangsta/negroni"
	"net/http"
)

var mu sync.Mutex
var count int

//返回一个Negroni实例
func NewServer() *negroni.Negroni{
	//negroni.Classic()返回一个实例，默认添加了一些中间件
	n := negroni.Classic()

	//定义一个路由来处理不同的URL
	mux := http.NewServeMux()
	//访问"/"下的处理
	mux.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		mu.Lock()
		count++
		mu.Unlock()
		fmt.Fprintf(rw, "Welcome to the home page!\n")
	})
	//访问"/count"下的处理，显示总共访问了多少次该服务器
	mux.HandleFunc("/count", func(rw http.ResponseWriter, r *http.Request) {
		mu.Lock()
		fmt.Fprintf(rw, "Total requests: %d\n", count)
		mu.Unlock()
	})
	//添加中间件,UseHandler添加的是http.Handler类型的中间件
	n.UseHandler(mux)

	//添加中间件，UseFunc添加的是Negroni Handler中间件
	//添加http.Handler中间件会默认调用下一个中间件，即printInfo
	n.UseFunc(printInfo)

	//在printInfo中间件中没有通过next(rw, r)来执行下一个中间件
	//因此printRequest不会执行
	//可以尝试去掉那一行注释再访问服务器查看输出
	n.UseFunc(printRequest)
	return n
}

func printInfo(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	fmt.Fprintf(rw, "This is a test server.\n")
	//next(rw, r)
}

func printRequest(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	fmt.Fprintf(rw, "URL.Path = %q\n", r.URL.Path)
}