package main

import (
	"avro2pg/database"
	"encoding/json"
	"fmt"
)

func saveGunState(originStr string) {
	defer func() {
		msg := recover()
		fmt.Println(msg)
	}()

	gunState := database.GunState{}

	json.Unmarshal([]byte(originStr), &gunState)

	database.DB.Save(&gunState)
}

// {"equipId" : "10001", "eventNo": "0", "eventKind": "0", "eventValue": "1",
//  "name": "홍길동", "rank":"병장", "dateTime" : "20210620101215","sosok":"제1중대/제1관활과"}
//
func saveGunEvent(originStr string) {
	defer func() {
		msg := recover()
		fmt.Println(msg)
	}()

	fmt.Println("guneventorigin   : ", originStr)
	gunEvent := database.GunEvent{}

	json.Unmarshal([]byte(originStr), &gunEvent)

	database.DB.Create(&gunEvent)
}
