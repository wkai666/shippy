package main

import (
    "context"
    "google.golang.org/grpc"
    "log"
    "net"
    pb "shippy/consignment-service/proto/consignment"
)

const (
    PORT = ":50051"
)

// IRepository 仓库接口
type IRepository interface {
    Create(consignment *pb.Consignment) (*pb.Consignment, error)    // 存放新货物
}

// Repository 存放多批货物，实现 IRepository 接口
type Repository struct {
    consignments []*pb.Consignment
}

// Create 创建货物
func (repo *Repository) Create(consignment *pb.Consignment) (*pb.Consignment, error) {
    repo.consignments = append(repo.consignments, consignment)
    return consignment,nil
}

// GetAll 获取货物
func (repo *Repository) GetAll() []*pb.Consignment {
    return repo.consignments
}

// service 定义微服务
type service struct {
    repo Repository
}

// service 实现 consignment.pb.go 中的 ShippingServiceServer 接口
// 使 service 作为 gRPC 的服务端

// CreateConsignment 托运新货物
func (s *service) CreateConsignment(ctx context.Context, req *pb.Consignment) (*pb.Response, error) {
    // 接受承运的货物
    consignment, err := s.repo.Create(req)
    if err != nil {
        return nil, err
    }

    resp := &pb.Response{Created: true, Consignment: consignment}
    return resp, nil
}

// GetConsignments 获取目前所有托运的货物
func (s *service) GetConsignments(ctx context.Context, req *pb.GetRequest) (*pb.Response, error) {
    allConsignments := s.repo.GetAll()
    resp := &pb.Response{Consignments: allConsignments}
    return resp, nil
}

func main()  {
    listener, err := net.Listen("tcp", PORT)
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }

    log.Printf("listen on: %s/n", PORT)

    server := grpc.NewServer()
    repo := Repository{}

    // 向 rpc 服务器注册微服务，此时会把我们自己实现的微服务 service 与协议中的 ShippingServiceServer 绑定
    pb.RegisterShippingServiceServer(server, &service{repo})

    if err:= server.Serve(listener); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}