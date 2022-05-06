package drone

import (
	"task/entity"
	"task/model"
)

func CreateDroneDBTransform(drone *model.DroneData , status int) *entity.Drone {
	return &entity.Drone{
		SerialNum:   drone.Serial,
		DroneModel:   drone.Model,
		Weight:      drone.Weight,
		BatteryCap:  drone.BatteryCapacity,
		Status:      status,
	}
}

func GetDroneDataTransform(d *entity.Drone ,status string) *model.DroneData {
	return &model.DroneData{
		Serial:          d.SerialNum,
		Model:           d.DroneModel,
		Weight:          d.Weight,
		BatteryCapacity: d.BatteryCap,
		Status:          status,
	}
}
