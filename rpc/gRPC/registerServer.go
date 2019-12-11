package main

import (
	"log"

	"github.com/xiongwei9/Gogogo/rpc/gRPC/lib"
)

func main() {
	err := lib.StartRpcServerRegister()
	if err != nil {
		log.Fatalf("StartRpcServerRegister error: %v", err)
	}
}
