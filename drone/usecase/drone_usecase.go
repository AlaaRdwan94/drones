package usecase

import (
	"task/drone"
	"task/entity"
	"task/model"
	transformer "task/transformer/drone"
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

func (d *DroneUsecase) ChangeStatus(serial string , statusNum int) (*model.DroneData,error) {
	var drone *entity.Drone
	drone.SerialNum = serial
	drone.Status = statusNum 
	drone , err := d.DroneRepo.UpdateDroneStatus(drone)
	if err != nil {
		return nil , err
	}
	status , err := d.DroneRepo.GetDronStatusByStatusNum(statusNum)
	if err != nil {
		return nil , err
	}
	return transformer.GetDroneDataTransform(drone,status.StatusName) , nil
}

//change the status of the drone helper functions
func (d *DroneUsecase) SetDroneToIDLE(serial string) (*model.DroneData,error) {
	return d.ChangeStatus(serial,1)
}
func (d *DroneUsecase) SetDroneToLOADING(serial string) (*model.DroneData,error) {
	return d.ChangeStatus(serial,2)
}
func (d *DroneUsecase) SetDroneToLOADED(serial string) (*model.DroneData,error) {
	return d.ChangeStatus(serial,3)
}
func (d *DroneUsecase) SetDroneToDELIVERING(serial string) (*model.DroneData,error) {
	return d.ChangeStatus(serial,4)
}
func (d *DroneUsecase) SetDroneToDELIVERED(serial string) (*model.DroneData,error) {
	return d.ChangeStatus(serial,5)
}
func (d *DroneUsecase) SetDroneToRETURNING(serial string) (*model.DroneData,error) {
	return d.ChangeStatus(serial,6)
}

func NewDroneUsecase(repo drone.Repositoy)  drone.Usecase{
   return &DroneUsecase{
	DroneRepo: repo,
   }
}