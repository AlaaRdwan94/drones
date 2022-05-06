package seeds

import (
	"task/entity"
	"task/infrastructure/db"
	"task/model"
)

func Seed() {
	var newDB db.Database
	newDB = db.NewPostgres()
	db := newDB.Open()
	defer db.Close()
	users := []entity.Status{
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
	for _, v := range users {
		db.GormDB.Create(&v)
	}

    //initialize map with prev data in database
	var drones []entity.Drone
	var drone entity.Drone
	db.GormDB.Select("serial_num").Table(drone.TableName()).Find(&drones)
	model.InitSerialMap(drones)
}
