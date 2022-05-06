package entity

type Status struct {
	StatusNum int `gorm:"primary_key"` 
	StatusName string 
}

func (s *Status) TableName() string {
	return "status"
}