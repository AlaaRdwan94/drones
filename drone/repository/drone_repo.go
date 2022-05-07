package repository

import (
	"encoding/json"
	"errors"
	"task/drone"
	"task/entity"
	"time"

	"github.com/jinzhu/gorm"
	"gopkg.in/redis.v3"
)

type DroneRepo struct {
	db *gorm.DB
	rd *redis.Client
}

// CreateDrone implements drone.Repositoy
func (d *DroneRepo) CreateDrone(drone *entity.Drone) (*entity.Drone, error) {
	//count number of drones 
	count := 0 
	d.db.Model(&entity.Drone{}).Count(&count)
	if count == 10 {
		return nil , errors.New("exceeded number of drones per fleet")
	}
	//create drone
	if err := d.db.Create(&drone).Error; err != nil {
		return nil, err
	}
	j, err := json.Marshal(&drone)
	if err != nil {
		return nil, err
	}
	//we create drone by serial as a key
	if err := d.rd.Set(drone.SerialNum, j, time.Hour).Err(); err != nil {
		return nil, err
	}
	return drone, nil
}

//update drone status
func (d *DroneRepo) UpdateDroneStatus(Drone *entity.Drone) (*entity.Drone , error) {
	if err := d.db.Model(Drone).Where("serial_num = ?",Drone.SerialNum).Update("status",Drone.Status).Error ; err != nil {
		return nil, err
	}
	return Drone, nil
}



func (d *DroneRepo) GetDronStatusByStatusNum(num int) (*entity.Status , error) {
	var Status entity.Status
	d.db.Model(entity.Status{StatusNum: num}).First(&Status)
	return &Status, nil
}

func NewDroneRepo(db *gorm.DB, rd *redis.Client) drone.Repositoy {
	return &DroneRepo{
		db: db,
		rd: rd,
	}
}
