package database

import "gorm.io/gorm"

type DroneStation struct {
	gorm.Model
	AwsHumidity      string `json:"aws_humidity"`
	AwsRain          string `json:"aws_rain"`
	AwsTemperature   string `json:"aws_temperature"`
	AwsWindDirection string `json:"aws_wind_direction"`
	AwsWindSpeed     string `json:"aws_wind_speed"`
	Msg              string `json:"msg"`
	Result           string `json:"result"`
	StClose          string `json:"st_close"`
	StOpen           string `json:"st_open"`
	StOperation      string `json:"st_operation"`
	StationID        string `json:"station_id"`
}

func (u *DroneStation) TableName() string {
	// custom table name, this is default
	return "s_army.drone_station"
}
