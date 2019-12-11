package main

import (
	"context"
	"log"

	"github.com/xiongwei9/Gogogo/rpc/gRPC/constant"
	pb "github.com/xiongwei9/Gogogo/rpc/gRPC/proto"
	"google.golang.org/grpc"
)

func main() {
	/*************** SSL认证 start ***************/
	//crtPath := "./ssl/server.crt"
	//_, filename, _, _ := runtime.Caller(1)
	//crtFile := path.Join(path.Dir(filename), crtPath)

	//crtFile := "/Users/xiongwei.zhu/programmer/go_workplace/Gogogo/rpc/gRPC/ssl/server.crt"
	//creds, err := credentials.NewClientTLSFromFile(crtFile, "Gogogo")
	//if err != nil {
	//	log.Fatalf("credentail error: %v", err)
	//}
	//
	//conn, err := grpc.Dial(constant.Address, grpc.WithTransportCredentials(creds)) // SSL认证
	/*************** SSL认证 end ***************/

	conn, err := grpc.Dial(constant.Address, grpc.WithInsecure()) // 无SSL认证
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
