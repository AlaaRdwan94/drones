package drone

import (
	"task/entity"
	"task/model"
)

func CreateDroneDBTransform(drone *model.DroneData, status int) *entity.Drone {
	return &entity.Drone{
		SerialNum:  drone.Serial,
		DroneModel: drone.Model,
		Weight:     drone.Weight,
		BatteryCap: drone.BatteryCapacity,
		Status:     status,
	}
}

func GetDroneDataTransform(d *entity.Drone, status string) *model.DroneData {
	return &model.DroneData{
		Serial:          d.SerialNum,
		Model:           d.DroneModel,
		Weight:          d.Weight,
		BatteryCapacity: d.BatteryCap,
		Status:          status,
	}
}


func GetDroneDataWithMedicationTransform(d *entity.Drone, m *[]entity.Medication, status string) *model.DroneData {
	return &model.DroneData{
		Token:           "",
		Serial:          d.SerialNum,
		Model:           d.DroneModel,
		Weight:          d.Weight,
		BatteryCapacity: d.BatteryCap,
		Status:          status,
		MedicationData:  getMedications(*m),
	}
}
func GetSingleDroneDataWithMedicationTransform(d *entity.Drone, m *entity.Medication, status string) *model.DroneData {
	arr := []model.MedicationData{}
	if m != nil {
		 mData := model.MedicationData{
		 	ID:     m.ID,
		 	Name:   m.Name,
		 	Weight: m.Weight,
		 	Code:   m.Code,
		 	Img:    m.Img,
		 }
		arr = append(arr, mData)
	}
	return &model.DroneData{
		Token:           "",
		Serial:          d.SerialNum,
		Model:           d.DroneModel,
		Weight:          d.Weight,
		BatteryCapacity: d.BatteryCap,
		Status:          status,
		MedicationData:  &arr,
	}
}
func getMedications(medArr []entity.Medication) (*[]model.MedicationData){
	arr := []model.MedicationData{}
	for _, v := range medArr {
		m := model.MedicationData{
			ID:     v.ID,
			Name:   v.Name,
			Weight: v.Weight,
			Code:   v.Code,
			Img:    v.Img,
		}
		arr = append(arr,m)
	}
	return &arr
}

//transform the droneMedication entity to Medication data
func GetDroneMedications(medArr []entity.Medication) (*[]model.MedicationData){
	arr := []model.MedicationData{}
	for _, v := range medArr {
		m := model.MedicationData{
			ID:     v.ID,
			Name:   v.Name,
			Weight: v.Weight,
			Code:   v.Code,
			Img:    v.Img,
		}
		arr = append(arr,m)
	}
	return &arr
}