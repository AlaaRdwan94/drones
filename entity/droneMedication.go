package entity

import "github.com/jinzhu/gorm"

type DroneMedication struct {
	gorm.Model
	DroneSerial  string
	MedicationID uint
}

func (d *DroneMedication) TableName() string {
	return "droneMedications"
}