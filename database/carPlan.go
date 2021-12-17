package database

import "gorm.io/gorm"

type CarPlan struct {
	gorm.Model
	CarsInOperationID         int    `gorm:"unique" json:"carsInOperationID"`
	CarLicenseNo              string `json:"carLicenseNo"`
	UserFullName              string `json:"userFullName"`
	CarResponsiblePerson      string `json:"carResponsiblePerson"`
	CarType                   string `json:"carType"`
	PersonOrCargo             int    `json:"personOrCargo"`
	RequiredDestinationName   string `json:"requiredDestinationName"`
	CarPassengers             string `json:"carPassengers"`
	RequiredDepartureDateTime int64  `json:"requiredDepartureDateTime"`
	RequiredArrivalDateTime   int64  `json:"requiredArrivalDateTime"`
	RealDrivingStartDateTime  int64  `json:"realDrivingStartDateTime"`
	RealDrivingEndDateTime    int64  `json:"realDrivingEndDateTime"`
	Route                     string `json:"route"`
	Type                      string `json:"type"`
}

func (u *CarPlan) TableName() string {
	// custom table name, this is default
	return "s_army.carplan2"
}

// {	carsInOperationID : 1,    gorm:"index"
// 	carLicenseNo : xxx,
// 	userFullName : xxx,
// 	carResponsiblePerson : xxx,
// 	carType : xxx,
// 	personOrCargo : 1, (0: 병력, 1: 화물, 2: 병력,화물),
// 	requiredDestinationName : xxx,
// 	carPassengers : xxx,
// 	requiredDepartureDateTime : 1265275107687,
// 	requiredArrivalDateTime : 1265275107687,
// 	realDrivingStartDateTime : 1265275107687,
// 	realDrivingEndDateTime : 1265275107687,
// 	route : xxx,
// 	type : C (C: 신규, U: 수정, D: 삭제)
// }

// {"carLicenseNo":"111","carOppNo":0,"armyUnit":"1대대","carType":"1톤봉고트럭","carDriver":"이홍진","carPassenger":"테스트","oppDestination":"232","oppPlanStartTime":"2021-11-30 19:25","oppPlanEndTime":"2021-11-30 19:25"}
// origin_id : 39cd1e9ed0da43c0bc7fb55e214e70e9
// device_id : 5aa8aa715144494e888f0708fd3e124c
