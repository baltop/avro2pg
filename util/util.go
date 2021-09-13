package util

import (
	"strconv"
	"strings"
	"time"

	uuid "github.com/satori/go.uuid"
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

func UUID24Digit() string {
	u2 := uuid.NewV4()
	t := strings.ReplaceAll(u2.String(), "-", "")
	return t[0:24]
}
