package proto

import pb "shippy/consignment-service/proto/consignment"

// IRepository 仓库接口
type IRepository interface {
	Create(consignment *pb.Consignment) (*pb.Consignment, error) // 存放新货物
}

// Repository 存放多批货物，实现 IRepository 接口
type Repository struct {
	consignments []*pb.Consignment
}

// Create 创建货物
func (repo *Repository) Create(consignment *pb.Consignment) (*pb.Consignment, error) {
	repo.consignments = append(repo.consignments, consignment)
	return consignment, nil
}

// GetAll 获取货物
func (repo *Repository) GetAll() []*pb.Consignment {
	return repo.consignments
}
