package main

// Goland: import的时候按目录导入，但是使用的时候需要以目录中的package的名字用，一个目录里面的所有go文件都必须使用同一个package名字
// 这里用pb重命名的"github.com/xiongwei9/Gogogo/rpc/gRPC/proto"的helloworld package
import (
	pb "github.com/xiongwei9/Gogogo/rpc/gRPC/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

const (
	port = ":50051"
)

type server struct {
}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	log.Printf("request name: %s", in.Name)
	return &pb.HelloResponse{Message: "hello, " + in.Name}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})

	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
