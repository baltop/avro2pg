package database

import "gorm.io/gorm"

type SensorGasAddr struct {
	City string `gorm:"null" json:"city"`
	Ku   string `gorm:"null" json:"ku"`
}

// Product struct
type SensorGas struct {
	gorm.Model
	NumberOfKey int64         `gorm:"not null" json:"numberOfKey"`
	Name        string        `gorm:"not null" json:"name"`
	A           float64       `gorm:"not null" json:"a"`
	Addr        SensorGasAddr `gorm:"embedded;embeddedPrefix:addr_" json:"addr"`
}

func (u *SensorGas) TableName() string {
	// custom table name, this is default
	return "s_army.sensorgas"
}
