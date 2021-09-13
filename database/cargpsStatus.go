package database

import (
	"time"

	"gorm.io/gorm"
)

type CarGpsStatus struct {
	ID                   uint
	CreatedAt            time.Time
	UpdatedAt            time.Time
	DeletedAt            gorm.DeletedAt `gorm:"index"`
	TripFuelEfficiency   int            `json:"tripFuelEfficiency"`
	FirmwareInfo         string         `json:"firmwareInfo"`
	DischargeVoltLevel   int            `json:"dischargeVoltLevel"`
	ManageType           string         `json:"manageType"`
	DashboardDist        int            `json:"dashboardDist"`
	EngineIdleDefault    int            `json:"engineIdleDefault"`
	GpsStatus            int            `json:"gpsStatus"`
	TppAlarm             int            `json:"tppAlarm"`
	QStopCnt             int            `json:"qStopCnt"`
	MsgCode              string         `json:"msgCode"`
	AxisActivation       int            `json:"axisActivation"`
	Longitude            int            `json:"longitude"`
	VisitPointID         int            `json:"visitPointId"`
	DeviceType           int            `json:"deviceType"`
	MaxRpmDefault        int            `json:"maxRpmDefault"`
	LongOverSpeedSec     int            `json:"longOverSpeedSec"`
	TripDrivingSec       int            `json:"tripDrivingSec"`
	StartIdleSec         int            `json:"startIdleSec"`
	KeyOnDt              int            `json:"keyOnDt"`
	QStartCnt            int            `json:"qStartCnt"`
	BatteryVolt          int            `json:"batteryVolt"`
	KeyOffDt             int            `json:"keyOffDt"`
	ReqDt                int            `json:"reqDt"`
	KeyOffReportInterval int            `json:"keyOffReportInterval"`
	KeyStatus            int            `json:"keyStatus"`
	Latitude             int            `json:"latitude"`
	GpsColdWarm          int            `json:"gpsColdWarm"`
	OverSpeedSec         int            `json:"overSpeedSec"`
	VisitPathID          int            `json:"visitPathId"`
	DeviceStatus         int            `json:"deviceStatus"`
	FuelEfficiency       int            `json:"fuelEfficiency"`
	GpsEffectiveDt       int            `json:"gpsEffectiveDt"`
	FuelUse              int            `json:"fuelUse"`
	QReduceCnt           int            `json:"qReduceCnt"`
	UnitID               int            `gorm:"primarykey" json:"unitId"`
	AccumulDist          int            `json:"accumulDist"`
	EcoLevel             string         `json:"ecoLevel"`
	ServiceID            string         `json:"serviceId"`
	KeyOnReportInterval  int            `json:"keyOnReportInterval"`
	Direction            int            `json:"direction"`
	QAccelCnt            int            `json:"qAccelCnt"`
	IdleSec              int            `json:"idleSec"`
	TripDist             int            `json:"tripDist"`
	Rpm                  int            `json:"rpm"`
	AvgSpeed             int            `json:"avgSpeed"`
	IdleFuelUse          int            `json:"idleFuelUse"`
	OverSpeedDefault     int            `json:"overSpeedDefault"`
	Tpp                  int            `json:"tpp"`
	HighRpmSec           int            `json:"highRpmSec"`
	DrivingSpeed         int            `json:"drivingSpeed"`
}

func (u *CarGpsStatus) TableName() string {
	// custom table name, this is default
	return "public.cargps_status"
}
