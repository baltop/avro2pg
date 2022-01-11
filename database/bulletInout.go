package database

import "gorm.io/gorm"

// Product struct
type BulletInout struct {
	gorm.Model
	UserName     string `json:"user_name"`
	NodeName     string `json:"node_name"`
	EvtNm        string `json:"evt_nm"`
	CarNumber    string `json:"car_number"`
	CarCategory  string `json:"car_category"`
	LocationName string `json:"location_name"`
	LocationType string `json:"location_type"`
	InDateEse    string `json:"in_date_ese"`
	OutDateEse   string `json:"out_date_ese"`
}

func (u *BulletInout) TableName() string {
	// custom table name, this is default
	return "s_army.bulletinout"
}
