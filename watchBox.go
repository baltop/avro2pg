package main

import (
	"avro2pg/database"
	"encoding/json"
	"fmt"
	"runtime/debug"
)

// type WatchBox struct {
// 	ChargerNo string `json:"charge_no"`
// 	CheckDate string `json:"check_date"`
// 	Status    []int  `json:"status"`
// 	Company   string `json:"company"`
// 	RoomName  string `json:"room_name"`
// }

func saveWatchBoxStatus(originStr string) {
	// loc, _ := time.LoadLocation("Asia/Seoul")

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered", r)
			debug.PrintStack()
		}
	}()

	watchBox := database.WatchBox{}
	json.Unmarshal([]byte(originStr), &watchBox)

	// watchBoxDb := database.WatchBox{}
	// watchBoxDb.ChargerNo = watchBox.ChargerNo
	// watchBoxDb.CheckDate = watchBox.CheckDate
	// watchBoxDb.Status = strings.Trim(strings.Replace(fmt.Sprint(watchBox.Status), " ", ",", -1), "[]")
	// watchBoxDb.Company = watchBox.Company
	// watchBoxDb.RoomName = watchBox.RoomName

	database.DB.Save(&watchBox)
}
