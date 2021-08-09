package database

import (
	"time"

	_ "gorm.io/gorm"
)

// 스마트폰 상태 gorm에서 지원안하는 데이터 타입으로 자동 table 생성에서 제외함. struct
type SmartphoneChargerStatus struct {
	KioskCode     string
	AllTrayStatus string
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     time.Time
}

func (u *SmartphoneChargerStatus) TableName() string {
	// custom table name, this is default
	return "public.smartphone_charger_status"
}
