package main

import (
    "context"
    "encoding/json"
    "errors"
    microclient "github.com/micro/go-micro/v2/client/grpc"
    "github.com/micro/go-micro/v2/config/cmd"
    _ "github.com/micro/go-micro/v2/registry/etcd"
    _ "github.com/micro/go-plugins/broker/nsq/v2"
    "io/ioutil"
    "log"
    pb "shippy/consignment-service/proto/consignment"
)

const (
    ADDRESS = "localhost:50051"
    DEFAULT_INFO_FILE = "consignment.json"
)

// parseFile 读取记录中的货物信息
func parseFile(fileName string) (*pb.Consignment, error) {
    data, err := ioutil.ReadFile(fileName)
    if err != nil {
        return nil, err
    }

    var consignment *pb.Consignment
    err = json.Unmarshal(data, &consignment)
    if err != nil {
        return nil, errors.New("consignment.json file content error")
    }

    return consignment, nil
}

func main()  {

    cmd.Init()

    // 初始化 rpc 客户端
    client := pb.NewShippingService("go.micro.srv.consignment", microclient.NewClient())

    container := &pb.Container{
        //Id: "938388383",
        CustomerId: "cust001",
        Origin: "Manchester, United Kingdom",
        UserId: "user001",
    }

    var containers []*pb.Container
    containers = append(containers, container)

    consignment := &pb.Consignment{
        VesselId: "vessel001",
        Weight: 550,
        Containers: containers,
        Description: "this is a consignment",
    }

    // 调用 rpc，将货物存储到自己的仓库中
    resp, err := client.CreateConsignment(context.Background(), consignment)

    if err != nil {
        log.Fatalf("could not create consignment: %v", err)
    }

    // 货物是否托运成功
    log.Printf("created: %t", resp.Created)

    // 批量货物托运
    resp, err = client.GetConsignments(context.Background(), &pb.GetRequest{})
    if err != nil {
        log.Fatalf("failed to list consignments: %v", err)
    }

    for _, c := range resp.Consignments {
        log.Printf("resp consignment is: %+v", c)
    }

}
