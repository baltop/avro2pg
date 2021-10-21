package database

import (
	"database/sql"

	"gorm.io/gorm"
)

// Product struct
type Incar struct {
	gorm.Model
	Command string `gorm:"not null" json:"command"`
	Data    Data   `gorm:"embedded" json:"data"`
}

type Data struct {
	CarNumber   string       `gorm:"not null" json:"car_number"`
	DateTime    string       `gorm:"-" json:"date_time"`
	InDateTime  sql.NullTime `gorm:"null" json:"in_date_time,omitempty"`
	OutDateTime sql.NullTime `gorm:"null" json:"out_date_time,omitempty"`
	Kind        string       `gorm:"not null" json:"kind"`
	Lprid       string       `gorm:"not null" json:"lprid"`
	Owner_name  string       `gorm:"null" json:"owner_name"`
	Group_name  string       `gorm:"null" json:"group_name"`
	Dept_name   string       `gorm:"null" json:"dept_name"`
}

func (u *Incar) TableName() string {
	// custom table name, this is default
	return "s_army.incar"
}
