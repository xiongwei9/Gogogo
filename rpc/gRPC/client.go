package main

import (
	"github.com/xiongwei9/Gogogo/rpc/gRPC/constant"
	pb "github.com/xiongwei9/Gogogo/rpc/gRPC/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"os"
)

func main() {
	conn, err := grpc.Dial(constant.Address, grpc.WithInsecure())
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
	if len(os.Args) > 1 {
		message = os.Args[1]
	}
	r, err := c.SayHello(context.Background(), &pb.HelloRequest{Message: message})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.Message)
}
