package main

import (
	"google.golang.org/grpc"
	"log"
	"net"
	"shippy/consignment-service/proto"
	pb "shippy/consignment-service/proto/consignment"
)

const (
	PORT = ":50051"
)

func main() {
	listener, err := net.Listen("tcp", PORT)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Printf("listen on: %s/n", PORT)

	server := grpc.NewServer()
	repo := proto.Repository{}

	// 向 rpc 服务器注册微服务，此时会把我们自己实现的微服务 service 与协议中的 ShippingServiceServer 绑定
	pb.RegisterShippingServiceServer(server, &proto.Service{repo})

	if err := server.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
