package main

import (
    "context"
    "gopkg.in/mgo.v2"
    pb "shippy/vessel-service/proto/vessel"
)

type handler struct {
    session *mgo.Session
}

func (h *handler) GetRepo() Repository {
    return &VesselRepository{h.session.Clone()}
}

func (h *handler) Create(ctx context.Context, req *pb.Vessel, resp *pb.Response) error {
    defer h.GetRepo().Close()

    if err := h.GetRepo().Create(req); err != nil {
        return err
    }

    resp.Vessel = req
    resp.Created = true
    return nil
}

func (h *handler) FindAvailable(ctx context.Context, req *pb.Specification, res *pb.Response) error {
    defer h.GetRepo().Close()

    vessel, err := h.GetRepo().FindAvailable(req)
    if err != nil {
        return err
    }

    res.Vessel = vessel
    return nil
}

type service struct {
    repo Repository
}

// FindAvailable 实现服务端
//func (s *service) FindAvailable(ctx context.Context, spec *pb.Specification, resp *pb.Response) error {
//    // 调用内部方法查找
//    v, err := s.repo.FindAvailable(spec)
//    if err != nil {
//        return err
//    }
//    resp.Vessel = v
//    return nil
//}
