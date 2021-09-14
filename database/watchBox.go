package database

type WatchBox struct {
	ChargerNo string `gorm:"primarykey" json:"charger_no"`
	CheckDate string `json:"check_date"`
	Status    string `json:"status"`
}

func (u *WatchBox) TableName() string {
	// custom table name, this is default
	return "public.watch_box"
}
