package main

import (
	"github.com/xiongwei9/Gogogo/rpc/demo/data"
	"log"
	"net/rpc"
)

func CallEcho(arg data.Message) (result *data.Message, err error) {
	var client *rpc.Client
	client, err = rpc.DialHTTP("tcp", ":9999") // 建立连接
	if err != nil {
		return nil, err
	}
	result = new(data.Message)
	err = client.Call("EchoService.Echo", arg, result) // 发送请求
	if err != nil {
		return nil, err
	}
	return result, err
}

func main() {
	done := make(chan struct{})

	go func() {
		result, err := CallEcho(data.Message{Name: "go", Age: 20})
		if err != nil {
			log.Fatal("error calling: ", err)
		} else {
			log.Println("call echo: ", *result)
		}
		done <- struct{}{}
	}()

	<-done
}
