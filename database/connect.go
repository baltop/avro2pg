package database

import (
	"avro2pg/config"

	"fmt"
	"log"
	"strconv"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// DB gorm connector
var DB *gorm.DB

// ConnectDB connect to db
func ConnectDB() {

	var err error
	p := config.Config("DB_PORT")
	port, err := strconv.ParseUint(p, 10, 32)
	if err != nil {
		panic(err)
	}

	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Seoul",
		config.Config("DB_HOST"),
		port,
		config.Config("DB_USER"),
		config.Config("DB_PASSWORD"),
		config.Config("DB_NAME"),
	)
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		PrepareStmt:            true,
		SkipDefaultTransaction: true,
		Logger:                 logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Println(err)
		panic("failed to connect database")
	}

	sqldb, e := DB.DB()
	if e != nil {
		panic("sqldb error")
	}
	sqldb.SetMaxOpenConns(100)
	// gorm.DefaultTableNameHandler = func(db *gorm.DB, tableName string) string {
	// 	return "public." + tableName
	// }
	fmt.Println("Connection Opened to Database")
	DB.AutoMigrate(&SensorGas{})
	DB.AutoMigrate(&Incar{})
	DB.AutoMigrate(&SmartphoneCharger{})

	DB.AutoMigrate(&SmartphoneChargerStatus{})

	DB.AutoMigrate(&GunState{})
	DB.AutoMigrate(&GunEvent{})
	DB.AutoMigrate(&DroneFlight{})
	DB.AutoMigrate(&DroneStation{})
	DB.AutoMigrate(&WatchBox{})
	DB.AutoMigrate(&CarOut{})
	DB.AutoMigrate(&CarPlan{})
	fmt.Println("Database Migrated")

}
