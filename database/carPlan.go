package database

import "gorm.io/gorm"

type CarPlan struct {
	gorm.Model
	CarLicenseNo     string `json:"carLicenseNo"`
	CarOppNo         string `json:"carOppNo"`
	ArmyUnit         string `json:"armyUnit"`
	CarType          string `json:"carType"`
	CarDriver        string `json:"carDriver"`
	CarPassenger     string `json:"carPassenger"`
	OppDestination   string `json:"oppDestination"`
	OppPlanStartTime string `json:"oppPlanStartTime"`
	OppPlanEndTime   string `json:"oppPlanEndTime"`
}

func (u *CarPlan) TableName() string {
	// custom table name, this is default
	return "s_army.carplan"
}

// {"carLicenseNo":"111","carOppNo":0,"armyUnit":"1대대","carType":"1톤봉고트럭","carDriver":"이홍진","carPassenger":"테스트","oppDestination":"232","oppPlanStartTime":"2021-11-30 19:25","oppPlanEndTime":"2021-11-30 19:25"}
// origin_id : 39cd1e9ed0da43c0bc7fb55e214e70e9
// device_id : 5aa8aa715144494e888f0708fd3e124c
