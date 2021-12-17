package database

import "gorm.io/gorm"

type CarOut struct {
	gorm.Model
	CarLicenseNo     string `json:"carLicenseNo"`
	KeyOnDt          int64  `json:"keyOnDt"`
	KeyOffDt         int64  `json:"keyOffDt"`
	KeyStatus        string `json:"keyStatus"`
	AvgSpeed         string `json:"avgSpeed"`
	OppLatitude      string `json:"latitude"`
	OppLongitude     string `json:"longitude"`
	FuelUse          int    `json:"fuelUse"`
	Rpm              int    `json:"rpm"`
	ReceivedDateTime string `json:"receivedDateTime"`
}

func (u *CarOut) TableName() string {
	// custom table name, this is default
	return "s_army.carout2"
}

// {"carLicenseNo":"7","keyOnDt":1637911260,"keyOffDt":1637911289,"keyStatus":0,"avgSpeed":0,"oppLatitude":"37763912","oppLongitude":"126785676","fuelUse":0,"rpm":0,"receivedDateTime":"2021-11-27 19:38:12"}
// origin_id : 39cd1e9ed0da43c0bc7fb55e214e70e9
// device_id : c07b5804e7fe4b38953a8355eb5dd793
