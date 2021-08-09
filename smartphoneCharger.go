package main

import (
	"avro2pg/database"
	"database/sql"
	"encoding/json"
	"fmt"
	"time"
)

func saveSmartPhoneCharger(originStr string) {

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

type SmartphoneChargerStatusJson struct {
	KioskCode     string          `json:"KioskCode"`
	AllTrayStatus json.RawMessage `json:"AllTrayStatus"`
}

func saveSmartPhoneChargerStatus(originStr string) {
	// loc, _ := time.LoadLocation("Asia/Seoul")

	smartphoneChargerStatusJson := SmartphoneChargerStatusJson{}
	json.Unmarshal([]byte(originStr), &smartphoneChargerStatusJson)

	smartPhoneChargerStatus := database.SmartphoneChargerStatus{}
	smartPhoneChargerStatus.KioskCode = smartphoneChargerStatusJson.KioskCode
	smartPhoneChargerStatus.AllTrayStatus = string(smartphoneChargerStatusJson.AllTrayStatus)
	smartPhoneChargerStatus.CreatedAt = time.Now()
	database.DB.Create(&smartPhoneChargerStatus)
}
