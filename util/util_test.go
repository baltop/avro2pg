package util

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestStringUnixToTime(t *testing.T) {
	oldTm := time.Now().Unix()
	fmt.Println(oldTm)
	newTm := UnixToDateTime(int(oldTm))
	fmt.Println("newTm ", newTm)
	a := strconv.Itoa(int(oldTm))
	//	tm := StringUnixToDateTime("1628628210")
	tm := StringUnixToDateTime(a)
	fmt.Println(tm)
	unitTimeInRFC3339 := tm.Format(time.RFC3339)
	fmt.Println("unix time stamp in unitTimeInRFC3339 format :->", unitTimeInRFC3339)
	b := tm.Unix()
	if oldTm != b {
		t.Error("eeeeeerror")
	}

}
