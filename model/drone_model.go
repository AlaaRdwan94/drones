package model

import (
	"math/rand"
	"strconv"
	"task/entity"
)

type DroneData struct {
	Token           string  `json:"token"`
	Serial          string  `json:"serial_number"`
	Model           string  `json:"model"`
	Weight          float32 `json:"weight"`
	BatteryCapacity float32 `json:"battery_capacity"`
	Status          string  `json:"status"`
	MedicationData *[]MedicationData `json:"medications"`
}

func (d *DroneData) GenerateSerial() {
	d.Serial = generate()
}

func generate() string {
	
	RandomAlphabetAndNumeric := 10
	
	rund := rand.Intn(1000000000)
	str := strconv.Itoa(rund)
	if len(str) < RandomAlphabetAndNumeric {
		count := RandomAlphabetAndNumeric - len(str)
		for i:=0 ; i < count ; i++ {
			str += "0"
		}
	}
	str = "DRONS" + "_" + str
	if _, found := serials[str] ; found == true {
		return generate()
	}
	serials[str] = str
	return str
}



//insure that we have a unique serials by adding them in a global map
var serials map[string]string

func InitSerialMap(drones []entity.Drone)  {
	serials = make(map[string]string,0)
	for _, v := range drones {
		serials[v.SerialNum] = v.SerialNum
	}
}

func GetSerials() map[string]string {
	return serials
}