package database

type WatchBox struct {
	ChargerNo string `gorm:"primarykey" json:"ChargerNo"`
	CheckDate string `json:"CheckDate"`
	Status    string `json:"Status"`
}

func (u *WatchBox) TableName() string {
	// custom table name, this is default
	return "public.watch_box"
}
