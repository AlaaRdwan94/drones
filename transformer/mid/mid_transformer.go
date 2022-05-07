package mid

import (
	"task/entity"
	"task/model"
)



func GetDroneDataTransform(m *entity.Medication, status string) *model.MedicationData {
	return &model.MedicationData{
		ID:     m.ID,
		Name:   m.Name,
		Weight: m.Weight,
		Code:   m.Code,
		Img:    m.Img,
	}
}