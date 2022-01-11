package main

import (
	"avro2pg/database"
	"encoding/json"
	"fmt"
	"runtime/debug"
)

func saveCarPlan(originStr string) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered", r)
			debug.PrintStack()
		}
	}()

	carPlan := database.CarPlan{}
	carSel := database.CarPlan{}

	json.Unmarshal([]byte(originStr), &carPlan)

	switch carPlan.Type {
	case "C":
		database.DB.Create(&carPlan)
	case "U":
		database.DB.First(&carSel, "cars_in_operation_id = ?", carPlan.CarsInOperationID)
		carPlan.ID = carSel.ID
		database.DB.Save(&carPlan)
	case "D":
		database.DB.First(&carSel, "cars_in_operation_id = ?", carPlan.CarsInOperationID)
		carPlan.ID = carSel.ID
		database.DB.Delete(&carPlan)
	case "X":
		database.DB.Delete(&carPlan, "car_license_no = ?", carPlan.CarLicenseNo)
	}

}

// {"equipId" : "10001", "eventNo": "0", "eventKind": "0", "eventValue": "1",
//  "name": "홍길동", "rank":"병장", "dateTime" : "20210620101215","sosok":"제1중대/제1관활과"}
//
func saveCarOut(originStr string) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered", r)
			debug.PrintStack()
		}
	}()

	carOut := database.CarOut{}

	json.Unmarshal([]byte(originStr), &carOut)

	database.DB.Create(&carOut)
}

func saveCarEvent(originStr string) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered", r)
			debug.PrintStack()
		}
	}()

	carEvent := database.CarEvent{}

	json.Unmarshal([]byte(originStr), &carEvent)

	database.DB.Create(&carEvent)
}
