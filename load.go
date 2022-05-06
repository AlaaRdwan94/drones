package main

import (
	"task/infrastructure/db"
	"task/infrastructure/worker/api_worker"
	_droneHandler "task/drone/handler"
	_droneRepo "task/drone/repository"
	_droneUsecase "task/drone/usecase"
	"github.com/gin-gonic/gin"
	"gopkg.in/redis.v3"
)

func InitializeRouts(db *db.DB, router *gin.RouterGroup, client *redis.Client) {
	// drone model
	droneRepository := _droneRepo.NewDroneRepo(db.GormDB,client)
	droneUsecase:= _droneUsecase.NewDroneUsecase(droneRepository)
	_droneHandler.NewDroneHandler(router,droneUsecase)
	api_worker.RunCron()
}