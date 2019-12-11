package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/xiongwei9/Gogogo/rpc/gRPC/constant"
	pb "github.com/xiongwei9/Gogogo/rpc/gRPC/proto"
	"google.golang.org/grpc"
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

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	c := pb.NewGreeterClient(conn)
	r, err := c.SayHelloAllStream(ctx, grpc.EmptyCallOption{})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	for i := 0; i < 10; i++ {
		err := r.Send(&pb.HelloRequest{Message: fmt.Sprintf("Hello, gRPC stream %d", i)})
		if err != nil {
			log.Printf("send error: %v", err)
			break
		}
	}
	err = r.CloseSend()
	if err != nil {
		log.Printf("close send error: %v", err)
	}

	for {
		res, err := r.Recv()
		if err != nil {
			if err.Error() == "EOF" {
				log.Println("EOF")
				break
			}
			log.Fatalf("recv client message error: %v", err)
		}
		log.Printf("recv server message: %v", res.Message)
	}
}
