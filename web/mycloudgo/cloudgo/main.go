package main

import (
	"os"
	flag "github.com/spf13/pflag"
	"github.com/chenf99/Golang/web/mycloudgo/service"
)

const PORT string = "8080"

func main() {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = PORT
	}

	//在启动服务器时可以输入端口号
	pPort := flag.StringP("port", "p", PORT, "PORT for httpd listening")
	flag.Parse()
	if len(*pPort) != 0 {
		port = *pPort
	}

	server := service.NewServer()
	//运行negroni的run方法
	//等价于http.ListenAndServe
	//但是会从环境变量PORT中获取服务监听的端口号
	server.Run(":" + port)
}