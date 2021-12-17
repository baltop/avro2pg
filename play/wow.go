package main

import (
	"avro2pg/config"
	"avro2pg/event"
	"fmt"
	"strings"

	uuid "github.com/satori/go.uuid"
)

func main() {
	config.TestLoad()

	a := fmt.Sprintf("%8v", config.Config("siteCD"))
	fmt.Println("sss" + a + "ttt" + config.Config("siteCD"))
	u2 := uuid.NewV4()
	fmt.Printf("UUIDv4: %s\n", u2)
	t := strings.ReplaceAll(u2.String(), "-", "")
	fmt.Println(t[0:24] + "pp")

	event.SendMessage("경로 이탈")
}
