package main

import (
	"github.com/micro/go-micro/v2"
	_ "github.com/micro/go-micro/v2/registry/etcd"
	_ "github.com/micro/go-plugins/broker/nsq/v2"
	"log"
	"os"
	pb "shippy/consignment-service/proto/consignment"
	vesselPb "shippy/vessel-service/proto/vessel"
)

//const DEFAULT_HOST = "172.21.0.1:27017"	// docker MongoDB
const DEFAULT_HOST = "127.0.0.1:27017"		// 本地 MongoDB

func main() {

	dbHost := os.Getenv("DB_HOST")
	if dbHost == "" {
		dbHost = DEFAULT_HOST
	}

	session, err := CreateSession(dbHost)
	// 创建于 MongoDB 的主会话，需在退出 main 的时候手动释放连接
	defer session.Close()
	if err != nil {
		log.Fatalf("create session error: %v\n", err)
	}

	server := micro.NewService(
		// 必须和 consignment.proto 中的 package 一致
		micro.Name("go.micro.srv.consignment"),
		micro.Version("latest"),
	)

	// 解析命令行参数
	server.Init()

	// 作为 vessel-service 客户端
	vClient := vesselPb.NewVesselService("go.micro.srv.vessel", server.Client())
	pb.RegisterShippingServiceHandler(server.Server(), &handler{session, vClient})

	if err := server.Run(); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
