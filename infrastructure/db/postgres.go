package db

import (
	"fmt"
	"task/entity"
	"task/infrastructure/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Postgres struct {
	gormDB *gorm.DB
}

func NewPostgres() *Postgres {
	return &Postgres{}
}

func (mysql *Postgres) Open() *DB {
	viper:=config.NewViper()
	DBMS := "postgres"
	dbConnection := viper.Database.Connection
	db, err := gorm.Open(DBMS, dbConnection)
	if err != nil {
		panic(err.Error())
	}
	db.LogMode(true)
	err = db.DB().Ping()
	if err != nil {
		panic(err)
	}
	m := db.AutoMigrate(&entity.Status{},&entity.Drone{},&entity.Medication{},&entity.DroneMedication{})

	if m != nil && m.Error != nil {
		//We have an error
		fmt.Printf(m.Error.Error())
	}

	return &DB{GormDB: db}
}
