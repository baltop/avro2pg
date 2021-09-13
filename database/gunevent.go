package database

import "gorm.io/gorm"

type GunEvent struct {
	gorm.Model
	EquipID    string `json:"equipId"`
	EventNo    string `json:"eventNo"`
	EventKind  string `json:"eventKind"`
	EventValue string `json:"eventValue"`
	DateTime   string `json:"dateTime"`
	Name       string `json:"name"`
	Rank       string `json:"rank"`
}

func (u *GunEvent) TableName() string {
	// custom table name, this is default
	return "public.gunevent"
}
