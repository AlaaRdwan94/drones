package model

type MedicationData struct {
	ID     uint    `json:"id"`
	Name   string  `json:"name"`
	Weight float32 `json:"weight"`
	Code   string  `json:"code"`
	Img    string  `json:"img"`
}

type LoadDroneRequest struct {
	ID     uint   `json:"id"`
	Serial string `json:"serial"`
}
