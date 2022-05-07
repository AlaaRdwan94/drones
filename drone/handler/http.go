package handler

import (
	"fmt"
	"log"
	"net/http"
	"task/drone"
	"task/middleware/auth"
	"task/model"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
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
		c.JSON(http.StatusInternalServerError,err.Error())
		return
	}

	token, err := auth.CreateToken(returned.Serial)
	if err != nil {
		c.JSON(http.StatusUnauthorized,err.Error())
		return
	}
	returned.Token = token
	c.JSON(http.StatusCreated,returned)
}


//load loads the medication to a drone
func (d *DroneHandler) Load(c *gin.Context) {

	c.Writer.Header().Set("Content-Type", "application/json")
	req := model.LoadDroneRequest{}
	if err := c.ShouldBindBodyWith(&req, binding.JSON); err != nil {
		log.Printf("%+v", err)
	}
	
	returned , err := d.Dusecase.AddMedication(req.Serial ,req.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError,err.Error())
		return
	}
	c.JSON(http.StatusCreated,returned)
}

//request for get by serial number APIs
type request struct {
	Serial string `json:"serial"`
}

//Get drone medications
func (d *DroneHandler) GetDroneMedications(c *gin.Context) {

	c.Writer.Header().Set("Content-Type", "application/json")
	req := request{}
	if err := c.ShouldBindBodyWith(&req, binding.JSON); err != nil {
		log.Printf("%+v", err)
	}
	
	returned , err := d.Dusecase.GetDroneMedications(req.Serial)
	if err != nil || returned == nil{
		c.JSON(http.StatusInternalServerError,err.Error())
		return
	}
	c.JSON(http.StatusAccepted,returned)
}

//Check avaliable drones for loading
func (d *DroneHandler) CheckLoadingDrones(c *gin.Context) {
	
	returned , err := d.Dusecase.GetLoadingDrone()
	if err != nil || returned == nil{
		c.JSON(http.StatusInternalServerError,err.Error())
		return
	}
	c.JSON(http.StatusAccepted,returned)
}

//Get drone medications
func (d *DroneHandler) GetDroneBattery(c *gin.Context) {

	c.Writer.Header().Set("Content-Type", "application/json")
	type request struct {
		Serial string `json:"serial"`
	}
	req := request{}
	if err := c.ShouldBindBodyWith(&req, binding.JSON); err != nil {
		log.Printf("%+v", err)
	}
	
	returned , err := d.Dusecase.GetDroneBatteryCap(req.Serial)
	if err != nil {
		c.JSON(http.StatusInternalServerError,err.Error())
		return
	}
	type response struct {
		BatteryCap string `json:"battery-capacity"`
	}
	res := response{}
	res.BatteryCap = fmt.Sprintf("%v",returned ) + "%"
	c.JSON(http.StatusAccepted,res)
}

func NewDroneHandler(e *gin.RouterGroup, dus drone.Usecase)  {
	handler := &DroneHandler{Dusecase: dus}
	e.POST("/register-drone",handler.RegisterDrone)
	e.POST("/load",handler.Load)
	e.GET("/medications",handler.GetDroneMedications)
	e.GET("/loading-drones",handler.CheckLoadingDrones)
	e.GET("/drone-battery",handler.GetDroneBattery)

}