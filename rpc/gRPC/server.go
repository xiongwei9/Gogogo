package main

import (
	"github.com/xiongwei9/Gogogo/rpc/gRPC/lib"
	"log"
)

func main() {
	err := lib.StartRpcServer()
	if err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
