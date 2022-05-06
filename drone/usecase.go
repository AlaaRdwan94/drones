package drone

import "task/model"

type Usecase interface {
	Register(*model.DroneData) (*model.DroneData ,error)
}
