package database

type GunState struct {
	EquipID   string `gorm:"primarykey" json:"equipId"`
	GunsState string `json:"gunsState"`
	LockState string `json:"lockState"`
	DateTime  string `json:"dateTime"`
	Name      string `json:"name"`
	Rank      string `json:"rank"`
	Company   string `json:"company"`
	Dormitory string `json:"dormitory"`
	GunNo     string `json:"gunNo"`
	GunKind   string `json:"gunKind"`
}

func (u *GunState) TableName() string {
	// custom table name, this is default
	return "s_army.gunstate"
}
