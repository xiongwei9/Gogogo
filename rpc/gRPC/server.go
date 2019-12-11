package main

import (
	"log"

	"github.com/xiongwei9/Gogogo/rpc/gRPC/lib"
)

func main() {
	err := lib.StartRpcServer()
	if err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
