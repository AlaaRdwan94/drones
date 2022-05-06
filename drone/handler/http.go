package handler

import (
	"task/middleware/auth"
	"task/model"
	"task/drone"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"log"
	"net/http"
)

type DroneHandler struct {
	Dusecase drone.Usecase
}

func (d *DroneHandler) RegisterDrone(c *gin.Context) {

	c.Writer.Header().Set("Content-Type", "application/json")
	drone := model.DroneData{}
	if err := c.ShouldBindBodyWith(&drone, binding.JSON); err != nil {
		log.Printf("%+v", err)
	}
	drone.GenerateSerial()
	returned , err := d.Dusecase.Register(&drone)
	if err != nil {
		c.JSON(http.StatusInternalServerError,err)
		return
	}

	token, err := auth.CreateToken(returned.Serial)
	if err != nil {
		c.JSON(http.StatusUnauthorized,err)
		return
	}
	returned.Token = token
	c.JSON(http.StatusCreated,returned)
}



func NewDroneHandler(e *gin.RouterGroup, dus drone.Usecase)  {
	handler := &DroneHandler{Dusecase: dus}
	e.POST("/register-drone",handler.RegisterDrone)
}