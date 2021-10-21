package database

import "gorm.io/gorm"

type GunEvent struct {
	gorm.Model
	EquipID    string `json:"equipId"`
	EventNo    string `json:"eventNo"`
	EventKind  string `json:"eventKind"`
	EventValue string `json:"eventValue"`
	DateTime   string `json:"dateTime"`
	Name       string `json:"name"`
	Rank       string `json:"rank"`
	Company    string `json:"company"`
	Dormitory  string `json:"dormitory"`
}

func (u *GunEvent) TableName() string {
	// custom table name, this is default
	return "s_army.gunevent"
}

// {"equipId" : "10001", "eventNo": "0", "eventKind": "0", "eventValue": "1",
//   "name": "홍길동", "rank":"병장", "dateTime" : "20210620101215","sosok":"제1중대/제1생활관"}
// eventNo  사용자 번호 0-9  G 관리자
// eventKind   0 총기 1 락
// eventValue   0 없음 열림 1 있음 닫힘
// sosok    중대 / 생활관
