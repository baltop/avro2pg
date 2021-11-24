package main

// main.go
import (
	"avro2pg/config"
	"avro2pg/database"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
	"time"

	"github.com/hamba/avro"
	kafka "github.com/segmentio/kafka-go"
)

var (
	schema avro.Schema
)

func init() {

	avscStr, err := ioutil.ReadFile("objectRequestPrimitive.avsc")
	if err != nil {
		panic(err)
	}

	// Create Schema Once
	schema, err = avro.Parse(string(avscStr))
	if err != nil {
		panic(err)
	}
}

func main() {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	config.Load()
	database.ConnectDB()

	// TODO  read from config
	kafkaURL := config.Config("KAFKA_HOST")
	topic := config.Config("KAFKA_TOPIC")
	groupID := "topostgres"

	kafkaReader := getKafkaReader(kafkaURL, topic, groupID)

	defer kafkaReader.Close()

	fmt.Println("start consuming ... !!")
	for {

		m, err := kafkaReader.ReadMessage(ctx)
		if err != nil {
			panic(err)
		}
		// fmt.Printf("message at topic:%v partition:%v offset:%v	%s = %s\n\n\n", m.Topic, m.Partition, m.Offset, string(m.Key), string(m.Value))

		kafkaMessage := make(map[string]interface{})
		err1 := avro.Unmarshal(schema, m.Value, &kafkaMessage)
		if err1 != nil {
			panic(err1)
		}

		if !strings.HasSuffix(kafkaMessage["to"].(string), "/origin") {
			fmt.Println("not origin")
			continue
		}

		fmt.Println("------------kafka to   [", kafkaMessage["to"].(string))
		// from 값이 device id
		fmt.Println("------------kafka from [", kafkaMessage["from"].(string))
		deviceId := kafkaMessage["from"].(string)
		content := make(map[string]interface{})

		err = json.Unmarshal([]byte(kafkaMessage["content"].(string)), &content)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(content)

		cin := content["m2m:cin"].(map[string]interface{})

		var originid string
		if config.Config("PROG_MODE") == "dev" {
			originid = "866785e4540c458985bfdc33ca2040d6"
		} else {
			lbl := cin["lbl"].([]interface{})
			originid = lbl[1].(string)
		}

		originStr := cin["con"].(string)
		fmt.Println(originStr)

		switch originid {
		case "866785e4540c458985bfdc33ca2040d6":
			go saveSensorGas(originStr)
		case "8accaef281e34f91b327d75b36f8660d":
			go saveIncar(originStr)
		case "d08fb5f50bf040b7aab8c5d00ea8a777":
			if deviceId == "3c76e4d2551a4956a1eb6162126aad3b" {
				go saveSmartPhoneCharger(originStr)
			} else if deviceId == "716b645637de42cfbed120018b7d5bd7" {
				go saveSmartPhoneChargerStatus(originStr)
			}
		case "39cd1e9ed0da43c0bc7fb55e214e70e9":
			go saveCarGps(originStr)
			go saveCarGpsStatus(originStr)
		case "61c42d3e474c47509e4adaebc08c8d47":
			if deviceId == "3a8d372a53904793bc7988c6a3f68db2" {
				go saveGunState(originStr)
			} else {
				go saveGunEvent(originStr)
			}
		case "b02a481a38f7496d97bf80c70ab7c207":
			go saveDroneStation(originStr)
		case "2f1a4cf209a34db0affab936762de55c":
			go saveDroneFlight(originStr)
		case "49d3fc4e819245dfa5005e372fa05cc4":
			fmt.Println("+++++++!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!+++++++++++++++++++++++++++++++++++++++")
			fmt.Println(originStr)
			fmt.Println("+++++++!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!+++++++++++++++++++++++++++++++++++++++")
			go saveWatchBoxStatus(originStr)
		}

	}
}

func saveIncar(originStr string) {
	defer func() {
		msg := recover()
		fmt.Println(msg)
	}()

	loc, _ := time.LoadLocation("Asia/Seoul")
	incar := database.Incar{}
	json.Unmarshal([]byte(originStr), &incar)
	dateTime, err := time.ParseInLocation("20060102150405", incar.Data.DateTime, loc)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(dateTime)

	if incar.Command == "smt_alert_incar" { // 입차시
		incar.Data.InDateTime = sql.NullTime{Time: dateTime, Valid: true}
		incar.Data.OutDateTime = sql.NullTime{Time: dateTime, Valid: false}
		database.DB.Create(&incar)
	} else { // 출차시
		selectCar := database.Incar{}
		database.DB.Model(&database.Incar{}).Select("*").Where("car_number = ?", incar.Data.CarNumber).Order("id desc").First(&selectCar)
		if selectCar.Data.InDateTime.Valid {
			selectCar.Data.OutDateTime = sql.NullTime{Time: dateTime, Valid: true}
			database.DB.Save(&selectCar)
		}
	}

}

func saveSensorGas(originStr string) {
	defer func() {
		msg := recover()
		fmt.Println(msg)
	}()

	sensor := database.SensorGas{}
	json.Unmarshal([]byte(originStr), &sensor)

	database.DB.Save(&sensor)
}

func getKafkaReader(kafkaURL, topic, groupID string) *kafka.Reader {
	brokers := strings.Split(kafkaURL, ",")
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers:     brokers,
		GroupID:     groupID,
		Topic:       topic,
		MinBytes:    10e3, // 10KB
		MaxBytes:    10e5, // 1MB
		StartOffset: kafka.LastOffset,
	})
}
