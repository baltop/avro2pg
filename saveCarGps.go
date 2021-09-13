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
	// carGps := database.CarGps{}
	json.Unmarshal([]byte(originStr), &carGpsList)
	// for val := range carGpsList.PacketLogList {
	// 	carGps = val
	// 	database.DB.Create(&carGps)
	// }
	database.DB.Create(&carGpsList.PacketLogList[0])
}

type CarGpsStatusList struct {
	PacketLogList []database.CarGpsStatus `json:"packetLogList"`
}

func saveCarGpsStatus(originStr string) {

	carGpsStatusList := CarGpsStatusList{}
	// carGps := database.CarGps{}
	json.Unmarshal([]byte(originStr), &carGpsStatusList)
	// for val := range carGpsList.PacketLogList {
	// 	carGps = val
	// 	database.DB.Create(&carGps)
	// }
	database.DB.Save(&carGpsStatusList.PacketLogList[0])
}
