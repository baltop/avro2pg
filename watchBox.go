package main

import (
	"avro2pg/database"
	"encoding/json"
	"fmt"
	"strings"
)

type WatchBox struct {
	ChargerNo string `json:"ChargerNo"`
	CheckDate string `json:"CheckDate"`
	Status    []int  `json:"Status"`
}

func saveWatchBoxStatus(originStr string) {
	// loc, _ := time.LoadLocation("Asia/Seoul")

	watchBox := WatchBox{}
	json.Unmarshal([]byte(originStr), &watchBox)

	watchBoxDb := database.WatchBox{}
	watchBoxDb.ChargerNo = watchBox.ChargerNo
	watchBoxDb.CheckDate = watchBox.CheckDate
	watchBoxDb.Status = strings.Trim(strings.Replace(fmt.Sprint(watchBox.Status), " ", ",", -1), "[]")

	database.DB.Save(&watchBoxDb)
}
