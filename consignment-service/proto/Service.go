package proto

import (
	"context"
	pb "shippy/consignment-service/proto/consignment"
)

// Service 定义微服务
type Service struct {
	Repo Repository
}

// Service 实现 consignment.pb.go 中的 ShippingServiceServer 接口
// 使 service 作为 gRPC 的服务端

// CreateConsignment 托运新货物
func (s *Service) CreateConsignment(ctx context.Context, req *pb.Consignment) (*pb.Response, error) {
	// 接受承运的货物
	consignment, err := s.Repo.Create(req)
	if err != nil {
		return nil, err
	}

	resp := &pb.Response{Created: true, Consignment: consignment}
	return resp, nil
}

// GetConsignments 获取目前所有托运的货物
func (s *Service) GetConsignments(ctx context.Context, req *pb.GetRequest) (*pb.Response, error) {
	allConsignments := s.Repo.GetAll()
	resp := &pb.Response{Consignments: allConsignments}
	return resp, nil
}
