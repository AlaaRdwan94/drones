package entity

import "github.com/jinzhu/gorm"

type Medication struct {
	gorm.Model
	Name   string
	Weight float32
	Code   string
	Img    string
}

func (m *Medication) ValidName() bool {
	for _, v := range m.Name {
		if (v >= 'a' && v <='z' ) || (v >='A' && v <='Z') || (v >= '0' && v <= '9') || v == '_' || v == '-' {
			return true
		}
	}
	return false
}

func (m *Medication) ValidCode() bool {
	for _, v := range m.Code {
		if (v >= 'a' && v <='z' ) || (v >='A' && v <='Z') || (v >= '0' && v <= '9') {
			return true
		}
	}
	return false
}

func (m *Medication) TableName() string {
	return "medications"
}