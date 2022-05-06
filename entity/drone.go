package entity

type Drone struct {
	SerialNum   string `gorm:"type:varchar(100);primary_key"`
	DroneModel   string
	Weight      float32
	BatteryCap  float32
	Status      int
	status      *Status `gorm:"foreignKey:Status;constraint:OnUpdate:CASCADE,OnDelete:0;"`
}

func (d *Drone) TableName() string {
	return "drones"
}

func (d *Drone) OverWeighted() bool {
	return d.Weight >= 500
}

func (d *Drone) BatteryLow() bool {
	return d.BatteryCap <= 25
}
