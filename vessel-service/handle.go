package main

import (
    "context"
    "gopkg.in/mgo.v2"
    "log"
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

    log.Println("vessel find available")

    res.Vessel = vessel
    return nil
}
