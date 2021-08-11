package util

import (
	"strconv"
	"time"
)

func StringUnixToDateTime(inputTime string) time.Time {
	time.Now()
	i, err := strconv.ParseInt(inputTime, 10, 64)
	if err != nil {
		panic(err)
	}
	tm := time.Unix(i, 0)
	return tm
}

func UnixToDateTime(inputTime int) time.Time {
	return time.Unix(int64(inputTime), 0)
}
