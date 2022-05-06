package usecase

import (
	"task/model"
	transformer "task/transformer/drone"
	"task/drone"
)

type DroneUsecase struct {
	DroneRepo drone.Repositoy
}

func (d *DroneUsecase) Register(model *model.DroneData) (*model.DroneData,error) {
	droneDB := transformer.CreateDroneDBTransform(model,1)
	drone , err := d.DroneRepo.CreateDrone(droneDB)
	if err != nil {
		return nil , err
	}
	return transformer.GetDroneDataTransform(drone,"IDLE") , nil
}


func NewDroneUsecase(repo drone.Repositoy)  drone.Usecase{
   return &DroneUsecase{
	DroneRepo: repo,
   }
}