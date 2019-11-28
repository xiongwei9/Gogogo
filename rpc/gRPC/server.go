package main

// Goland: import的时候按目录导入，但是使用的时候需要以目录中的package的名字用，一个目录里面的所有go文件都必须使用同一个package名字
// 这里用pb重命名的"github.com/xiongwei9/Gogogo/rpc/gRPC/proto"的helloworld package
import (
	"fmt"
	"github.com/xiongwei9/Gogogo/rpc/gRPC/constant"
	pb "github.com/xiongwei9/Gogogo/rpc/gRPC/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"strings"
)

type server struct {
}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	log.Printf("recv client message: %s", in.Message)
	return &pb.HelloResponse{Message: "message: " + in.Message}, nil
}
func (s *server) SayHelloServerStream(request *pb.HelloRequest, srv pb.Greeter_SayHelloServerStreamServer) error {
	msg := request.Message
	for i := 0; i < 3; i++ {
		err := srv.Send(&pb.HelloResponse{Message: msg})
		if err != nil {
			return err
		}
	}
	return nil
}
func (s *server) SayHelloClientStream(srv pb.Greeter_SayHelloClientStreamServer) error {
	var msgList []string
	for {
		req, err := srv.Recv()
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			log.Fatalf("recv client message error: %v", err)
		}
		log.Printf("recv client message: %v", req.Message)
		msgList = append(msgList, req.Message)
	}

	err := srv.SendAndClose(&pb.HelloResponse{Message: strings.Join(msgList, "\n")})
	return err
}
func (s *server) SayHelloAllStream(srv pb.Greeter_SayHelloAllStreamServer) error {
	ch := make(chan struct{})
	go func() {
		for {
			req, err := srv.Recv()
			if err != nil {
				if err.Error() == "EOF" {
					ch <- struct{}{}
					break
				}
				log.Fatalf("recv client message error: %v", err)
			}
			log.Printf("recv client message: %v", req.Message)
		}
	}()

	for i := 0; i < 3; i++ {
		msg := fmt.Sprintf("hello from server stream %d", i)
		log.Printf("send stream message: %s", msg)
		err := srv.Send(&pb.HelloResponse{Message: msg})
		if err != nil {
			return err
		}
	}

	<-ch
	return nil
}

func main() {
	lis, err := net.Listen("tcp", ":"+constant.Port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	gRpcServer := grpc.NewServer()
	pb.RegisterGreeterServer(gRpcServer, &server{})

	reflection.Register(gRpcServer)
	if err := gRpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
