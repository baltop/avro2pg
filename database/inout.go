package database

import "gorm.io/gorm"

// Product struct
type InoutManage struct {
	gorm.Model
	UserName     string `json:"user_name"`
	NodeName     string `json:"node_name"`
	EvtNm        string `json:"evt_nm"`
	CarNumber    string `json:"car_number"`
	CarCategory  int    `json:"car_category"`
	LocationName string `json:"location_name"`
	LocationType int    `json:"location_type"`
	InDateEse    string `json:"in_date_ese"`
	OutDateEse   string `json:"out_date_ese"`
}

func (u *InoutManage) TableName() string {
	// custom table name, this is default
	return "s_army.inoutmanage"
}
