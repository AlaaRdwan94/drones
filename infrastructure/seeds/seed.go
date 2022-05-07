package seeds

import (
	"task/entity"
	"task/infrastructure/db"
	"task/infrastructure/worker"
	"task/model"
)

func Seed() {
	var newDB db.Database
	newDB = db.NewPostgres()
	db := newDB.Open()
	defer db.Close()
    //seed database only one time
	countstatus := 0
	db.GormDB.Model(&entity.Status{}).Count(&countstatus)
	if countstatus == 0 {
	SeedWithStatus(db)
	}

	countMedications := 0 
	db.GormDB.Model(&entity.Drone{}).Count(&countMedications)
	if countMedications == 0 {
		SeedWithMedications(db)
	}

    //initialize map with prev data in database
	var drones []entity.Drone
	var drone entity.Drone
	db.GormDB.Select("serial_num").Table(drone.TableName()).Find(&drones)
	model.InitSerialMap(drones)

	//update drone battery
	UpdateDroneBattery(db)
}

func SeedWithMedications(db *db.DB) {
	meds := []entity.Medication{
		{
			Name:   "Med_1",
			Weight: 10,
			Code:   "Med110",
			Img:    "https://cdn.picpng.com/pills/tablets-pills-drugs-medication-77410.png",
		},
		{
			Name:   "Med_2",
			Weight: 100,
			Code:   "Med2100",
			Img:    "https://cdn.picpng.com/pills/tablets-pills-drugs-medication-77410.png",
		},
		{
			Name:   "Med_3",
			Weight: 50,
			Code:   "Med350",
			Img:    "https://cdn.picpng.com/pills/tablets-pills-drugs-medication-77410.png",
		},
		{
			Name:   "Med_4",
			Weight: 10,
			Code:   "Med410",
			Img:    "https://cdn.picpng.com/pills/tablets-pills-drugs-medication-77410.png",
		},
		{
			Name:   "Med_5",
			Weight: 210,
			Code:   "Med5210",
			Img:    "https://cdn.picpng.com/pills/tablets-pills-drugs-medication-77410.png",
		},
		{
			Name:   "Med_6",
			Weight: 210,
			Code:   "Med6210",
			Img:    "https://cdn.picpng.com/pills/tablets-pills-drugs-medication-77410.png",
		},
		{
			Name:   "Med_7",
			Weight: 20,
			Code:   "Med720",
			Img:    "https://cdn.picpng.com/pills/tablets-pills-drugs-medication-77410.png",
		},
		{
			Name:   "Med_8",
			Weight: 40,
			Code:   "Med840",
			Img:    "https://cdn.picpng.com/pills/tablets-pills-drugs-medication-77410.png",
		},
	}
	for _, v := range meds {
		if v.ValidName() && v.ValidCode() {
			db.GormDB.Create(&v)
		}
	}
}

func SeedWithStatus(db *db.DB) {
	status := []entity.Status{
		{
			StatusNum:  1,
			StatusName: "IDLE",
		},
		{
			StatusNum:  2,
			StatusName: "LOADING",
		},
		{
			StatusNum:  3,
			StatusName: "LOADED",
		},
		{
			StatusNum:  4,
			StatusName: "DELIVERING",
		},
		{
			StatusNum:  5,
			StatusName: "DELIVERED",
		},
		{
			StatusNum:  6,
			StatusName: "RETURNING",
		},
	}
	for _, v := range status {
		db.GormDB.Create(&v)
	}
}

//update battery if status is loading or returning
func UpdateDroneBattery(db *db.DB) {
	// reduse 10% from the battery in status delevering and returning every 1 min
	if err := db.GormDB.Model(&entity.Drone{}).Where("status = 4 OR status = 6").Update("battery_cap","battery_cap - 10").Error ; err != nil {
		worker.Historylog("error update battery for all drones...")
	}
	drones := []entity.Drone{}
	if err := db.GormDB.Where("status = 4 OR status = 6").Find(&drones).Error ; err != nil {
		worker.Historylog("error loading drones for all drones...")
	}
	for _, v := range drones {
		worker.Historylog("Drone with Serial:",v.SerialNum,"Has Battery:",v.BatteryCap,"%")
	}
} 