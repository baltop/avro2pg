package main

import (
	"avro2pg/database"
	"encoding/json"
)

type CarGpsList struct {
	PacketLogList []database.CarGps `json:"packetLogList"`
}

func saveCarGps(originStr string) {

	carGpsList := CarGpsList{}
	json.Unmarshal([]byte(originStr), &carGpsList)
	for val := range carGpsList.PacketLogList {
		carGps := val
		database.DB.Create(&carGps)
	}

}
