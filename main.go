package main

import (
	"fmt"
	"task/infrastructure/config"
	"task/infrastructure/db"
	"github.com/gin-gonic/gin"
)

func main()  {
	viper:=config.NewViper()
	var newDB db.Database
	newDB = db.NewPostgres()
	newRD := db.NewRedis()
	RdClient ,err := newRD.NewClient()
	obj := newDB.Open()
	defer obj.Close()
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	v1 := router.Group("/api/v1")
	InitializeRouts(obj, v1,RdClient)
	err = router.Run(viper.Server.Port)

	if err != nil {
		fmt.Println(err)
	}
}

