package drone

import "task/entity"

type Repositoy interface {
	CreateDrone(*entity.Drone) (*entity.Drone ,error)
	GetDronStatusByStatusNum( int) (*entity.Status , error)
	UpdateDroneStatus(Drone *entity.Drone) (*entity.Drone , error)
	UpdateDroneWeight(Drone *entity.Drone) (*entity.Drone, error)
	CreateMedication(DroneMedication *entity.DroneMedication) (*entity.DroneMedication, error)
	GetDroneMedications(serial string) (*[]entity.Medication, error)
	GetDroneBySerialNum(num string) (*entity.Drone, error)
	GetSingleMedication(id uint) (*entity.Medication, error)
}
