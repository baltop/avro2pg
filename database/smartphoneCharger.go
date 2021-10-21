package database

import (
	"database/sql"

	"gorm.io/gorm"
)

// Product struct
type SmartphoneCharger struct {
	gorm.Model
	KioskCode      string `gorm:"not null" json:"KioskCode"`
	TimeStart      string `gorm:"-" json:"TimeUseStart,omitempty"`
	TimeEnd        string `gorm:"-" json:"TimeUseEnd,omitempty"`
	Time_use_start sql.NullTime
	Time_use_end   sql.NullTime
	UseCode        string `gorm:"null" json:"UseCode"`
	UsingTray      string `gorm:"null" json:"UsingTray"`
}

func (u *SmartphoneCharger) TableName() string {
	// custom table name, this is default
	return "s_army.smartphone_charger"
}
