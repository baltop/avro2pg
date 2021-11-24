package database

type WatchBox struct {
	ChargerNo string `gorm:"primarykey" json:"charger_no"`
	CheckDate string `json:"check_date"`
	Status    string `json:"status"`
	Company   string `json:"company"`
	RoomName  string `json:"room_name"`
}

func (u *WatchBox) TableName() string {
	// custom table name, this is default
	return "s_army.watch_box"
}
