package main

import (
	"avro2pg/database"
	"database/sql"
	"encoding/json"
	"fmt"
	"runtime/debug"
	"time"
)

func saveSmartPhoneCharger(originStr string) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered", r)
			debug.PrintStack()
		}
	}()

	loc, _ := time.LoadLocation("Asia/Seoul")
	smartphoneCharger := database.SmartphoneCharger{}
	json.Unmarshal([]byte(originStr), &smartphoneCharger)
	if smartphoneCharger.TimeStart != "" {
		dateTime, err := time.ParseInLocation("20060102150405", smartphoneCharger.TimeStart, loc)
		if err != nil {
			fmt.Println(err.Error())
		}
		smartphoneCharger.Time_use_start = sql.NullTime{Time: dateTime, Valid: true}
		smartphoneCharger.Time_use_end = sql.NullTime{Time: dateTime, Valid: false}
	} else {
		dateTime, err := time.ParseInLocation("20060102150405", smartphoneCharger.TimeEnd, loc)
		if err != nil {
			fmt.Println(err.Error())
		}
		smartphoneCharger.Time_use_end = sql.NullTime{Time: dateTime, Valid: true}
		smartphoneCharger.Time_use_start = sql.NullTime{Time: dateTime, Valid: false}
	}

	fmt.Printf("smartphonCharger save %v\n", smartphoneCharger)

	database.DB.Create(&smartphoneCharger)

}

type SmartphoneChargerStatusJsonSub struct {
	TrayNo     string
	TrayStatus string
}

type SmartphoneChargerStatusJson struct {
	KioskCode     string                           `json:"KioskCode"`
	AllTrayStatus []SmartphoneChargerStatusJsonSub `json:"AllTrayStatus"`
}

func saveSmartPhoneChargerStatus(originStr string) {
	// loc, _ := time.LoadLocation("Asia/Seoul")
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered", r)
			debug.PrintStack()
		}
	}()

	smartphoneChargerStatusJson := SmartphoneChargerStatusJson{}
	json.Unmarshal([]byte(originStr), &smartphoneChargerStatusJson)

	smartPhoneChargerStatus := database.SmartphoneChargerStatus{}
	smartPhoneChargerStatus.KioskCode = smartphoneChargerStatusJson.KioskCode
	lastStatus := ""
	for i := 0; i < len(smartphoneChargerStatusJson.AllTrayStatus); i++ {
		olp := smartphoneChargerStatusJson.AllTrayStatus[i]
		if olp.TrayStatus == "E" {
			lastStatus += ",0"
		} else {
			lastStatus += ",1"
		}

	}
	lastStatus = lastStatus[1:]

	smartPhoneChargerStatus.AllTrayStatus = lastStatus
	smartPhoneChargerStatus.CreatedAt = time.Now()
	database.DB.Save(&smartPhoneChargerStatus)
}
