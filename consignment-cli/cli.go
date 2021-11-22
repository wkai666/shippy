package main

import (
    "context"
    "encoding/json"
    "errors"
    "google.golang.org/grpc"
    "io/ioutil"
    "log"
    "os"
    pb "shippy/consignment-service/proto/consignment"
)

const (
    ADDRESS = "127.0.0.1:50051"
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
    // 连接到 rpc 服务器
    conn, err := grpc.Dial(ADDRESS, grpc.WithInsecure())
    if err != nil {
        log.Fatalf("connent error: %v", err)
    }

    defer conn.Close()

    // 初始化 rpc 客户端
    client := pb.NewShippingServiceClient(conn)

    // 命令行中获取货物信息
    infoFile := DEFAULT_INFO_FILE
    if len(os.Args) > 1 {
        infoFile = os.Args[1]
    }

    // 解析货物信息
    consignment, err := parseFile(infoFile)
    if err != nil {
        log.Fatalf("parse file info error: %v", err)
    }

    // 调用 rpc，将货物存储到自己的仓库中
    resp, err := client.CreateConsignment(context.Background(), consignment)
    if err != nil {
        log.Fatalf("create consignment error: %v", err)
    }

    // 货物是否托运成功
    log.Printf("created: %t", resp.Created)

    // 批量货物托运
    resp, err = client.GetConsignments(context.Background(), &pb.GetRequest{})
    if err != nil {
        log.Fatalf("failed to list consignments: %v", err)
    }

    for _, c := range resp.Consignments {
        log.Printf("%+v", c)
    }
}

