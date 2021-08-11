package main

import (
	"avro2pg/database"
	"encoding/json"
)

func saveGunState(originStr string) {

	gunState := database.GunState{}

	json.Unmarshal([]byte(originStr), &gunState)

	database.DB.Create(&gunState)
}

func saveGunEvent(originStr string) {

	gunEvent := database.GunEvent{}

	json.Unmarshal([]byte(originStr), &gunEvent)

	database.DB.Create(&gunEvent)
}
