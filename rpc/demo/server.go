package main

import (
	"log"
	"net/http"
	"net/rpc"

	"github.com/xiongwei9/Gogogo/rpc/demo/data"
)

type EchoService struct {
}

func (service EchoService) Echo(arg data.Message, result *data.Message) error {
	log.Println("receive: ", arg)
	arg.Age++
	*result = arg
	return nil
}

func RegisterAndServe() {
	err := rpc.Register(&EchoService{}) // 注册EchoService定义的服务
	if err != nil {
		log.Fatal("error registering", err)
		return
	}
	rpc.HandleHTTP()                        // 使用HTTP处理请求
	err = http.ListenAndServe(":9999", nil) // 开始监听
	if err != nil {
		log.Fatal("error listening", err)
	}
}

func main() {
	RegisterAndServe()
}
