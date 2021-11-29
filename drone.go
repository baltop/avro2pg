package main

import (
	"avro2pg/database"
	"encoding/json"
	"fmt"
	"runtime/debug"
	"strings"
)

func saveDroneStation(originStr string) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered", r)
			debug.PrintStack()
		}
	}()

	// loc, _ := time.LoadLocation("Asia/Seoul")
	droneStation := database.DroneStation{}
	json.Unmarshal([]byte(originStr), &droneStation)

	fmt.Printf("DroneStation save %v\n", droneStation)

	database.DB.Create(&droneStation)
}

func saveDroneFlight(originStr string) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered", r)
			debug.PrintStack()
		}
	}()
	droneFlight := database.DroneFlight{}
	err := json.Unmarshal([]byte(originStr), &droneFlight)
	if err != nil {
		fmt.Printf("DrongFlight msg format err %v\n", err)
		return
	}
	asciiSubstring := droneFlight.AvPos[1 : len(droneFlight.AvPos)-1]
	droneFlight.AvPos = strings.ReplaceAll(asciiSubstring, ",", "")

	fmt.Printf("DroneFlight save %v\n", droneFlight)

	database.DB.Create(&droneFlight)
}
