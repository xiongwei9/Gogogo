package main

import (
	"github.com/xiongwei9/Gogogo/rpc/gRPC/lib"
	"log"
)

func main() {
	err := lib.StartRpcServerRegister()
	if err != nil {
		log.Fatalf("StartRpcServerRegister error")
	}
}
