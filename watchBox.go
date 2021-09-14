package main

import (
	"avro2pg/database"
	"encoding/json"
)

//type WatchBox struct {
//	ChargerNo string `json:"ChargerNo"`
//	CheckDate string `json:"CheckDate"`
//	Status    []int  `json:"Status"`
//}

func saveWatchBoxStatus(originStr string) {
	// loc, _ := time.LoadLocation("Asia/Seoul")

	watchBox := database.WatchBox{}
	json.Unmarshal([]byte(originStr), &watchBox)

	database.DB.Save(&watchBox)
}
