package usecase

import (
	"errors"
	"task/drone"
	"task/entity"
	"task/model"
	transformer "task/transformer/drone"
)

type DroneUsecase struct {
	DroneRepo drone.Repositoy
}

// GetLoadingDrone implements drone.Usecase
func (d *DroneUsecase) GetLoadingDrone() (*[]model.DroneData, error) {
	drones, err := d.DroneRepo.GetLoadingDrone()
	if err != nil {
		return nil, err
	}
	arr := []model.DroneData{}
	for _, v := range *drones {
		medications , err := d.DroneRepo.GetDroneMedications(v.SerialNum)
		if err != nil {
			return nil, err
		}
		data := transformer.GetDroneDataWithMedicationTransform(&v,medications, v.SerialNum)
		arr = append(arr, *data)
	}
	return &arr , nil
}

// GetDroneMedications implements drone.Usecase
func (d *DroneUsecase) GetDroneMedications(serial string) (*[]model.MedicationData, error) {
	mids, err := d.DroneRepo.GetDroneMedications(serial)
	if err != nil {
		return nil, err
	}
	return transformer.GetDroneMedications(*mids), nil
}

// AddMedication implements drone.Usecase
func (d *DroneUsecase) AddMedication(serial string, id uint) (*model.DroneData, error) {

	drone, err := d.DroneRepo.GetDroneBySerialNum(serial)
	if err != nil {
		return nil, err
	}

	medication, err := d.DroneRepo.GetSingleMedication(id)
	if err != nil {
		return nil, err
	}
	if drone.Status == 6 {
		drone.Status = 1
		_, err = d.SetDroneToIDLE(drone.SerialNum) //reset the status and asume that return is finished
		if err != nil {
			return nil, err
		}
	}
	if drone.BatteryCap <= 25 {
		return nil, errors.New("low Battery")
	}
	if drone.Status != 1 && drone.Status != 2 {
		return nil, errors.New("cannot load to this drone ...")
	}
	drone.Weight = drone.Weight + medication.Weight
	if drone.Weight < 500 {

		drone, err = d.DroneRepo.UpdateDroneWeight(drone)
		if err != nil {
			return nil, err
		}
	} else {
		_, err = d.SetDroneToLOADED(serial)
		if err != nil {
			return nil, err
		}
		return nil, errors.New("exceed the max weight for this drone")
	}
	obj := entity.DroneMedication{
		DroneSerial:  serial,
		MedicationID: id,
	}
	_, err = d.DroneRepo.CreateMedication(&obj)
	if err != nil {
		return nil, err
	}
	status, err := d.DroneRepo.GetDronStatusByStatusNum(drone.Status)
	if err != nil {
		return nil, err
	}
	return transformer.GetSingleDroneDataWithMedicationTransform(drone, medication, status.StatusName), nil
}

func (d *DroneUsecase) Register(model *model.DroneData) (*model.DroneData, error) {
	droneDB := transformer.CreateDroneDBTransform(model, 1)
	drone, err := d.DroneRepo.CreateDrone(droneDB)
	if err != nil {
		return nil, err
	}
	status, err := d.DroneRepo.GetDronStatusByStatusNum(1)
	return transformer.GetDroneDataTransform(drone, status.StatusName), nil
}

func (d *DroneUsecase) ChangeStatus(serial string, statusNum int) (*model.DroneData, error) {
	drone := entity.Drone{
		SerialNum: serial,
		Status:    statusNum,
	}

	new, err := d.DroneRepo.UpdateDroneStatus(&drone)
	if err != nil {
		return nil, err
	}
	status, err := d.DroneRepo.GetDronStatusByStatusNum(statusNum)
	if err != nil {
		return nil, err
	}
	return transformer.GetDroneDataTransform(new, status.StatusName), nil
}

//change the status of the drone helper functions
func (d *DroneUsecase) SetDroneToIDLE(serial string) (*model.DroneData, error) {
	return d.ChangeStatus(serial, 1)
}
func (d *DroneUsecase) SetDroneToLOADING(serial string) (*model.DroneData, error) {
	return d.ChangeStatus(serial, 2)
}
func (d *DroneUsecase) SetDroneToLOADED(serial string) (*model.DroneData, error) {
	return d.ChangeStatus(serial, 3)
}
func (d *DroneUsecase) SetDroneToDELIVERING(serial string) (*model.DroneData, error) {
	return d.ChangeStatus(serial, 4)
}
func (d *DroneUsecase) SetDroneToDELIVERED(serial string) (*model.DroneData, error) {
	return d.ChangeStatus(serial, 5)
}
func (d *DroneUsecase) SetDroneToRETURNING(serial string) (*model.DroneData, error) {
	return d.ChangeStatus(serial, 6)
}

func NewDroneUsecase(repo drone.Repositoy) drone.Usecase {
	return &DroneUsecase{
		DroneRepo: repo,
	}
}
