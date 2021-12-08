package database

import "gorm.io/gorm"

type DroneFlight struct {
	gorm.Model
	Msg             string `json:"msg"`
	Result          string `json:"result"`
	DroneID         string `json:"drone_id"`
	FlightMode      string `json:"flight_mode"`
	AvPos           string `json:"av_pos"`
	Speed           string `json:"speed"`
	Roll            string `json:"roll"`
	Pitch           string `json:"pitch"`
	Heading         string `json:"heading"`
	Altitude        string `json:"altitude"`
	Battery         string `json:"battery"`
	MslAlt          string `json:"msl_alt"`
	AglAlt          string `json:"agl_alt"`
	AvDted          string `json:"av_dted"`
	GpsAlt          string `json:"gps_alt"`
	StreamAddress   string `json:"streamAddress"`
	CctvId          string `json:"cctv_id"`
	OperateStatus   string `json:"operate_status"`
	TargetLatitude  string `json:"target_latitude"`
	TargetLongitude string `json:"target_longitude"`
	TargetCount     string `json:"target_count"`
}

func (u *DroneFlight) TableName() string {
	// custom table name, this is default
	return "s_army.drone_flight"
}
