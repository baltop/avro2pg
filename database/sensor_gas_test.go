package database

import (
	"avrotest/config"
	"testing"
)

func TestSensorDataInsert(t *testing.T) {
	config.TestLoad()
	ConnectDB()
	sensorGasAddr := SensorGasAddr{"busan", "sasang"}
	sensorGas := SensorGas{
		NumberOfKey: 21,
		Name:        "smartLight",
		A:           120.45,
		Addr:        sensorGasAddr,
	}
	DB.Save(&sensorGas)
	sensor := SensorGas{}
	result := DB.Find(&sensor, "name = ?", "smartLight")
	t.Log(result.RowsAffected)
	t.Log(result.Error)
	t.Logf("%v\n", sensor)

	t.Log(config.Config("DB_NAME"))

	if result.Error != nil {
		t.Error("error occured")
	}
}
