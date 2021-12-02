package main

import (
    "gopkg.in/mgo.v2"
    "gopkg.in/mgo.v2/bson"
    pb "shippy/vessel-service/proto/vessel"
)

const (
    dbName              = "shippy"
    vesselCollection    = "vessels"
)

type Repository interface {
    FindAvailable(*pb.Specification) (*pb.Vessel, error)
    Create(vessel *pb.Vessel) error
    Close()
}

type VesselRepository struct {
    session *mgo.Session
}

func (repo *VesselRepository) Create(v *pb.Vessel) error {
    return repo.collection().Insert(v)
}

func (repo *VesselRepository) FindAvailable(spec *pb.Specification) (*pb.Vessel, error) {

    var vessel *pb.Vessel

    err := repo.collection().Find(bson.M{
        "capacity": bson.M{"$gte":spec.Capacity},
        "maxweight": bson.M{"$gte":spec.MaxWeight},
    }).One(&vessel)
    if err != nil {
        return nil, err
    }

    return vessel, nil
}

func (repo *VesselRepository) Close()  {
    repo.session.Close()
}

func (repo *VesselRepository) collection() *mgo.Collection {
    return repo.session.DB(dbName).C(vesselCollection)
}
