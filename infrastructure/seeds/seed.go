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
//NODE:we assume that the drone uses it's battery only when its status is delivering or returning
//NOTE:we assume that the drone is full battery when it's status is IDLE
func UpdateDroneBattery(db *db.DB) {
	// reduce 10% from the battery in status delevering and returning every 1 min
	if err := db.GormDB.Raw("UPDATE drones SET battery_cap = battery_cap - 10 WHERE status = 4 OR status = 6").Scan(&[]entity.Drone{}).Error; err != nil {
		worker.Historylog("error update battery for all drones...")
	}
	drones := []entity.Drone{}
	if err := db.GormDB.Where("status = 4 OR status = 6").Find(&drones).Error; err != nil {
		worker.Historylog("error loading drones for all drones...")
	}
	for _, v := range drones {
		worker.Historylog("Drone with Serial:", v.SerialNum, "Has Battery:", v.BatteryCap, "%")
	}

	//recharge the battery if the status is idle
	if err := db.GormDB.Raw("UPDATE drones SET battery_cap = 100 WHERE status = 1").Scan(&entity.Drone{}).Error; err != nil {
		worker.Historylog("error update battery for all drones...")
	}
}

//update loaded drons to DELIVERING, will run it every 2 mins
//NODE:we create this cron as simulation for drone delivering functionality
func UpdateDroneToDelevering() {
	var newDB db.Database
	newDB = db.NewPostgres()
	db := newDB.Open()
	defer db.Close()

	if err := db.GormDB.Model(&entity.Drone{}).Where("status = 3 ").Update("status", 4).Error; err != nil {
		worker.Historylog("error update status for all drones with status 3 [loaded]...")
	}
}

//update delivering to delivered will run every 5 min
//NODE:we create this cron as simulation for drone delivering functionality
//NOTE:we assume that while the drone is delievered then its weight is reset to 0 and
//NOTE:remove all medications loaded to this drone
func UpdateDroneToDelevered() {
	var newDB db.Database
	newDB = db.NewPostgres()
	db := newDB.Open()
	defer db.Close()

	if err := db.GormDB.Model(&entity.Drone{}).Where("status = 4 ").Update("status", 5).Update("weight", 0).Error; err != nil {
		worker.Historylog("error update status for all drones with status 4 [delivering]...")
	}
	if err := db.GormDB.Model(&entity.Drone{}).Where("status = 5 ").Update("weight", 0).Error; err != nil {
		worker.Historylog("error reset weight for all drones with status 5 [delivered]...")
	}
	deleteFromDroneMedications(db)
}

//delete from droneMedications where status is delivered
func deleteFromDroneMedications(db *db.DB) {
	DroneMedication := entity.DroneMedication{}
	if err := db.GormDB.Raw(`
	DELETE FROM "droneMedications" WHERE drone_serial IN ( SELECT serial_num FROM drones WHERE status = 5)`).Scan(&DroneMedication).Error; err != nil {
		worker.Historylog("error delete from droneMedications...")
	}
	
}

//update delivering to delivered will run every 3 min
func UpdateDroneToReturing() {
	var newDB db.Database
	newDB = db.NewPostgres()
	db := newDB.Open()
	defer db.Close()

	if err := db.GormDB.Model(&entity.Drone{}).Where("status = 5 ").Update("status", 6).Error; err != nil {
		worker.Historylog("error update status for all drones with status 5 [delivered]...")
	}
}
