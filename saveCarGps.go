package main

import (
	"avro2pg/database"
	"encoding/json"
	"fmt"
)

type CarGpsList struct {
	PacketLogList []database.CarGps `json:"packetLogList"`
}

func saveCarGps(originStr string) {

	carGpsList := CarGpsList{}
	json.Unmarshal([]byte(originStr), &carGpsList)

	fmt.Printf("%v \n", carGpsList)
}
