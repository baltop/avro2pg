package main

import (
	"avro2pg/database"
	"encoding/json"
	"fmt"
)

func saveGunState(originStr string) {

	gunState := database.GunState{}

	json.Unmarshal([]byte(originStr), &gunState)

	database.DB.Save(&gunState)
}

func saveGunEvent(originStr string) {

	fmt.Println("guneventorigin   : ", originStr)
	gunEvent := database.GunEvent{}

	json.Unmarshal([]byte(originStr), &gunEvent)

	database.DB.Create(&gunEvent)
}
