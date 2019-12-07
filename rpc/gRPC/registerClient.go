package main

import (
	"context"
	"fmt"
	"github.com/xiongwei9/Gogogo/rpc/gRPC/constant"
	pb "github.com/xiongwei9/Gogogo/rpc/gRPC/proto"
	"github.com/xiongwei9/Gogogo/rpc/gRPC/registerCenter"
	"google.golang.org/grpc"
	"log"
)

func main() {
	schema, err := registerCenter.GenerateAndRegisterEtcdResolver(constant.EtcdAddress, "HelloService")
	if err != nil {
		log.Fatalf("init etcd resolver error: %v", err)
	}

	address := fmt.Sprintf("%s:///HelloService", schema)
	conn, err := grpc.Dial(address, grpc.WithInsecure()) // 无SSL认证
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer func() {
		err := conn.Close()
		if err != nil {
			log.Fatalf("close connection error: %v", err)
		}
	}()

	c := pb.NewGreeterClient(conn)
	message := "hello, gRPC"

	r, err := c.SayHello(context.Background(), &pb.HelloRequest{Message: message})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.Message)
}
