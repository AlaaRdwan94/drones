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

//CreateMedication for a drone
func (d *DroneRepo) CreateMedication(DroneMedication *entity.DroneMedication) (*entity.DroneMedication, error) {
	if err := d.db.Create(&DroneMedication).Error; err != nil {
		return nil, err
	}
	return DroneMedication, nil
}

//FIXME: get drone medications with less database calls 
func (d *DroneRepo) GetDroneMedications(serial string) (*[]entity.Medication, error) {
	DroneMedications := []entity.DroneMedication{}
	if err := d.db.Where("drone_serial = ?", serial).Find(&DroneMedications).Error; err != nil {
		return nil, err
	}
	Medications := []entity.Medication{}
	for _, v := range DroneMedications {
		m , err := d.GetSingleMedication(v.MedicationID)
		if err != nil {
			return nil, err
		}
		Medications = append(Medications, *m)
	}
	return &Medications, nil
}

//get all loading drones
func (d *DroneRepo) GetLoadingDrone() (*[]entity.Drone, error) {
	Drones := []entity.Drone{}
	if err := d.db.Where("status = 2").Find(&Drones).Error; err != nil {
		return nil, err
	}
	
	return &Drones, nil
}

//get from medications table by ID
func (d *DroneRepo) GetSingleMedication(id uint) (*entity.Medication, error) {
	var Medication entity.Medication
	d.db.First(&Medication,id)
	return &Medication, nil
}

//get single drone from drones table by serial number
func (d *DroneRepo) GetDroneBySerialNum(serial string) (*entity.Drone, error) {
	var Drone entity.Drone
	d.db.Model(entity.Drone{SerialNum: serial}).Find(&Drone)
	return &Drone, nil
}

// CreateDrone implements drone.Repositoy
func (d *DroneRepo) CreateDrone(drone *entity.Drone) (*entity.Drone, error) {
	//count number of drones
	count := 0
	d.db.Model(&entity.Drone{}).Count(&count)
	if count == 10 {
		return nil, errors.New("exceeded number of drones per fleet")
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
func (d *DroneRepo) UpdateDroneStatus(Drone *entity.Drone) (*entity.Drone, error) {
	if err := d.db.Model(Drone).Update("status", Drone.Status).Error; err != nil {
		return nil, err
	}
	return Drone, nil
}

//update drone weight and change the status to loading
//NOTE: for simplicity we asume that the status for the drone should be loading only when there is updates in it's weight
func (d *DroneRepo) UpdateDroneWeight(Drone *entity.Drone) (*entity.Drone, error) {
	if err := d.db.Model(Drone).Update("weight", Drone.Weight).Update("status",2).Error; err != nil {
		return nil, err
	}
	return Drone, nil
}

//check the status for a drone
func (d *DroneRepo) GetDronStatusByStatusNum(num int) (*entity.Status, error) {
	var Status entity.Status
	d.db.Model(entity.Status{StatusNum: num}).First(&Status ,num)
	return &Status, nil
}

func NewDroneRepo(db *gorm.DB, rd *redis.Client) drone.Repositoy {
	return &DroneRepo{
		db: db,
		rd: rd,
	}
}
