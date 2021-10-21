package database

import "gorm.io/gorm"

type DroneFlight struct {
	gorm.Model
	Msg           string  `json:"msg"`
	Result        string  `json:"result"`
	DroneID       string  `json:"drone_id"`
	FlightMode    string  `json:"flight_mode"`
	AvPos         string  `json:"av_pos"`
	Speed         int     `json:"speed"`
	Roll          float64 `json:"roll"`
	Pitch         float64 `json:"pitch"`
	Heading       int     `json:"heading"`
	Altitude      int     `json:"altitude"`
	Battery       int     `json:"battery"`
	MslAlt        int     `json:"msl_alt"`
	AglAlt        int     `json:"agl_alt"`
	AvDted        int     `json:"av_dted"`
	GpsAlt        int     `json:"gps_alt"`
	StreamAddress string  `json:"streamAddress"`
}

func (u *DroneFlight) TableName() string {
	// custom table name, this is default
	return "s_army.drone_flight"
}
