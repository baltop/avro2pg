package main

import (
	"avro2pg/database"
	"encoding/json"
	"fmt"
	"strings"
)

func saveDroneStation(originStr string) {

	// loc, _ := time.LoadLocation("Asia/Seoul")
	droneStation := database.DroneStation{}
	json.Unmarshal([]byte(originStr), &droneStation)

	fmt.Printf("DroneStation save %v\n", droneStation)

	database.DB.Create(&droneStation)
}

func saveDroneFlight(originStr string) {
	droneFlight := database.DroneFlight{}
	json.Unmarshal([]byte(originStr), &droneFlight)
	asciiSubstring := droneFlight.AvPos[1 : len(droneFlight.AvPos)-1]
	droneFlight.AvPos = strings.ReplaceAll(asciiSubstring, ",", "")

	fmt.Printf("DroneFlight save %v\n", droneFlight)

	database.DB.Create(&droneFlight)
}
