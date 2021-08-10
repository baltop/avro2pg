package database

import "gorm.io/gorm"

type GunState struct {
	gorm.Model
	EquipID   string `json:"equipId"`
	GunsState string `json:"gunsState"`
	LockState string `json:"lockState"`
	DateTime  string `json:"dateTime"`
	Name      string `json:"name"`
	Rank      string `json:"rank"`
}

func (u *GunState) TableName() string {
	// custom table name, this is default
	return "public.gunstate"
}
