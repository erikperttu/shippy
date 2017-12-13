package main

import pb "github.com/erikperttu/shippy/vessel-service/proto/vessel"

func (repo *VesselRepository) Create(vessel *pb.Vessel) error {
	return repo.collection().Insert(vessel)
}
