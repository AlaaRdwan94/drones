package drone

import "task/entity"

type Repositoy interface {
	CreateDrone(*entity.Drone) (*entity.Drone ,error)
	GetDronStatusByStatusNum( int) (*entity.Status , error)
	UpdateDroneStatus(Drone *entity.Drone) (*entity.Drone , error)
}
