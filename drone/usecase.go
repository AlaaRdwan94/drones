package drone

import "task/model"

type Usecase interface {
	Register(*model.DroneData) (*model.DroneData ,error)
	SetDroneToIDLE(serial string) (*model.DroneData,error)
	SetDroneToLOADING(serial string) (*model.DroneData,error)
	SetDroneToLOADED(serial string) (*model.DroneData,error)
	SetDroneToDELIVERING(serial string) (*model.DroneData,error)
	SetDroneToDELIVERED(serial string) (*model.DroneData,error) 
	SetDroneToRETURNING(serial string) (*model.DroneData,error) 
}
