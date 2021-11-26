package main

import (
	"context"
	"github.com/micro/go-micro/v2"
	"log"
	pb "shippy/consignment-service/proto/consignment"
	_ "github.com/micro/go-micro/v2/registry/etcd"
	_ "github.com/micro/go-plugins/broker/nsq/v2"
)

// IRepository 仓库接口
type IRepository interface {
	Create(consignment *pb.Consignment) (*pb.Consignment, error) 	// 存放新货物
	GetAll() []*pb.Consignment										// 获取仓库中存放的所有货物
}

// Repository 存放多批货物，实现 IRepository 接口
type Repository struct {
	Consignments []*pb.Consignment
}

// Create 创建货物
func (repo *Repository) Create(consignment *pb.Consignment) (*pb.Consignment, error) {
	repo.Consignments = append(repo.Consignments, consignment)
	return consignment, nil
}

// GetAll 获取货物
func (repo *Repository) GetAll() []*pb.Consignment {
	return repo.Consignments
}

// Service 定义微服务
type Service struct {
	Repo Repository
}

// Service 实现 consignment.pb.go 中的 ShippingServiceServer 接口
// 使 service 作为 gRPC 的服务端

// CreateConsignment 托运新货物
func (s *Service) CreateConsignment(ctx context.Context, req *pb.Consignment, resp *pb.Response)  error {
	// 接受承运的货物
	consignment, err := s.Repo.Create(req)
	if err != nil {
		return err
	}

	resp = &pb.Response{Created: true, Consignment: consignment}
	return nil
}

// GetConsignments 获取目前所有托运的货物
func (s *Service) GetConsignments(ctx context.Context, req *pb.GetRequest, resp *pb.Response) error {
	allConsignments := s.Repo.GetAll()
	resp = &pb.Response{Consignments: allConsignments}
	return nil
}


func main() {
	server := micro.NewService(
		// 必须和 consignment.proto 中的 package 一致
		micro.Name("go.micro.srv.consignment"),
		micro.Version("latest"),
	)

	// 解析命令行参数
	server.Init()
	repo := Repository{}
	pb.RegisterShippingServiceHandler(server.Server(), &Service{repo})

	if err := server.Run(); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
