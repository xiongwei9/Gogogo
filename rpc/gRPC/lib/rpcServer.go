package lib

import (
	"context"
	"fmt"
	"github.com/xiongwei9/Gogogo/rpc/gRPC/constant"
	pb "github.com/xiongwei9/Gogogo/rpc/gRPC/proto"
	"github.com/xiongwei9/Gogogo/rpc/gRPC/registerCenter"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"strings"
	"time"
)

/**
 * 使用注册中心的gRpc服务器
 */
func StartRpcServerRegister() error {
	lis, err := net.Listen("tcp", ":"+constant.Port)
	if err != nil {
		return err
	}

	etcdRegister := registerCenter.NewRegisterImpl([]string{constant.EtcdAddress}, time.Second*3)
	go func() {
		for {
			err := etcdRegister.Register(registerCenter.ServiceDescInfo{
				ServiceName:  "HelloService",
				Host:         constant.Host,
				Port:         constant.PortInt,
				IntervalTime: time.Duration(10),
			})
			if err != nil {
				log.Fatalf("register service error")
			}
			time.Sleep(time.Second * 5)
		}
	}()

	gRpcServer := grpc.NewServer()
	pb.RegisterGreeterServer(gRpcServer, &server{})

	// 开启服务端
	reflection.Register(gRpcServer)
	if err := gRpcServer.Serve(lis); err != nil {
		return err
	}
	return nil
}

func StartRpcServer() error {
	lis, err := net.Listen("tcp", ":"+constant.Port)
	if err != nil {
		return err
	}

	/*************** SSL认证 start ***************/
	/*
	 * SSL crt file config
	 * Country: CN
	 * Province: Guangdong
	 * Locality: shenzhen
	 * Organization: Gogogo
	 * Organizational Unit: IT
	 * Common Name: Gogogo
	 * Email Address: zhuxiongwei@foxmail.com
	 */
	//crtPath := "./ssl/server.crt"
	//keyPath := "./ssl/server.key"
	//_, filename, _, _ := runtime.Caller(1)
	//
	//crtFile := path.Join(path.Dir(filename), crtPath)
	//keyFile := path.Join(path.Dir(filename), keyPath)
	//creds, err := credentials.NewServerTLSFromFile(crtFile, keyFile)
	//if err != nil {
	//	return err
	//}

	//gRpcServer := grpc.NewServer(grpc.Creds(creds))
	/*************** SSL认证 start ***************/

	gRpcServer := grpc.NewServer()
	pb.RegisterGreeterServer(gRpcServer, &server{})

	reflection.Register(gRpcServer)
	if err := gRpcServer.Serve(lis); err != nil {
		return err
	}
	return nil
}

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
