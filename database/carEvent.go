package database

import "gorm.io/gorm"

type CarEvent struct {
	gorm.Model
	CarsInOperationID int     `gorm:"index" json:"carsInOperationID"`
	CarLicenseNo      string  `json:"carLicenseNo"`
	NotifyType        int     `json:"notifyType"`
	Longitude         float64 `json:"longitude"`
	Latitude          float64 `json:"latitude"`
	ReceivedDateTime  int64   `json:"receivedDateTime"`
}

func (u *CarEvent) TableName() string {
	// custom table name, this is default
	return "s_army.carevent"
}

// {	carsInOperationID : 1,
// 	carLicenseNo : xxx,
// 	notifyType : 1 (1. skip_TP 2. wrong_destination 3. late_arrival / 4:사고+인명 / 5:사고+차대차 / 6:고장+기동불가 / 7:고장+기동가능),
// 	longitude : 127.000,
// 	latitude : 37,000,
// 	receivedDateTime : 1265275107687
// }
