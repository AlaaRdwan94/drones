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
	status , err := d.DroneRepo.GetDronStatusByStatusNum(1)
	return transformer.GetDroneDataTransform(drone,status.StatusName) , nil
}


func NewDroneUsecase(repo drone.Repositoy)  drone.Usecase{
   return &DroneUsecase{
	DroneRepo: repo,
   }
}