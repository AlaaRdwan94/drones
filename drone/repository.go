package drone

import "task/entity"

type Repositoy interface {
	CreateDrone(*entity.Drone) (*entity.Drone ,error)
}
