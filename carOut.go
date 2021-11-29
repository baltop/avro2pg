package main

import (
	"avro2pg/database"
	"encoding/json"
	"fmt"
)

func saveCarPlan(originStr string) {
	defer func() {
		msg := recover()
		fmt.Println(msg)
	}()

	carPlan := database.CarPlan{}

	json.Unmarshal([]byte(originStr), &carPlan)

	database.DB.Create(&carPlan)
}

// {"equipId" : "10001", "eventNo": "0", "eventKind": "0", "eventValue": "1",
//  "name": "홍길동", "rank":"병장", "dateTime" : "20210620101215","sosok":"제1중대/제1관활과"}
//
func saveCarOut(originStr string) {
	defer func() {
		msg := recover()
		fmt.Println(msg)
	}()

	carOut := database.CarOut{}

	json.Unmarshal([]byte(originStr), &carOut)

	database.DB.Create(&carOut)
}
