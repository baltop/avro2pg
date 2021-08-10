package database

import "gorm.io/gorm"

type GunEvent struct {
	gorm.Model
	EquipID     string `json:"equipId"`
	GunsState   string `json:"gunsState"`
	LockState   string `json:"lockState"`
	GunsEvent   string `json:"gunsEvent"`
	LockEvent   string `json:"lockEvent"`
	EventReason string `json:"eventReason"`
	DateTime    string `json:"dateTime"`
}

func (u *GunEvent) TableName() string {
	// custom table name, this is default
	return "public.gunevent"
}
