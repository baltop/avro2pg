package main

import (
	"avro2pg/database"
	"encoding/json"
	"fmt"
	"runtime/debug"
)

func saveInoutmanage(originStr string) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered", r)
			debug.PrintStack()
		}
	}()
	inOutManage := database.InoutManage{}

	json.Unmarshal([]byte(originStr), &inOutManage)

	database.DB.Save(&inOutManage)
}

func saveBulletInout(originStr string) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered", r)
			debug.PrintStack()
		}
	}()
	BulletInOut := database.BulletInout{}

	json.Unmarshal([]byte(originStr), &BulletInOut)

	database.DB.Save(&BulletInOut)
}
