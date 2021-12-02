package main

import (
    "github.com/micro/go-micro/v2"
    _ "github.com/micro/go-micro/v2/registry/etcd"
    _ "github.com/micro/go-plugins/broker/nsq/v2"
    "log"
    "os"
    pb "shippy/vessel-service/proto/vessel"
)

const DEFAULT_HOST = "172.21.0.1:27017"     // docker MongoDB
//const DEFAULT_HOST = "127.0.0.1:27017"		// 本地 MongoDB

func createDummyData(repo Repository)  {
    defer repo.Close()

    vessels := []*pb.Vessel{
        {
            Id: "vessel001",
            Name: "Kane's Salty secret",
            MaxWeight: 20000,
            Capacity: 500,
        },
    }

    for _, v := range vessels {
        repo.Create(v)
    }
}

func main()  {

    host := os.Getenv("DB_HOST")
    if host == "" {
        host = DEFAULT_HOST
    }

    session, err := CreateSession(host)
    defer session.Close()

    if err != nil {
        log.Fatalf("Error connecting to datastore: %v", err)
    }

    repo := &VesselRepository{session.Copy()}

    createDummyData(repo)

    server := micro.NewService(
        micro.Name("go.micro.srv.vessel"),
        micro.Version("latest"),
    )
    server.Init()

    pb.RegisterVesselServiceHandler(server.Server(), &handler{session})

    if err := server.Run(); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}
