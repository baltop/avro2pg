package event

import (
	"avro2pg/config"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"strings"
	"time"

	uuid "github.com/satori/go.uuid"
)

type Item struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type statEvetItem []Item

// type StatEvetCntn []interface{}

// 나중에 json으로 변경해서 body content를 만듬.
type BodyJson struct {
	StatEvet struct {
		StatEvetGdCd     int    `json:"statEvetGdCd"`
		OutbPosNm        string `json:"outbPosNm"`
		StatEvetNm       string `json:"statEvetNm"`
		StatEvetClrDtm   string `json:"statEvetClrDtm"`
		StatEvetActnRslt string `json:"statEvetActnRslt"`

		StatEvetCntn          string       `json:"statEvetCntn"`
		SutbScopRads          string       `json:"outbScopRads"`
		OutbPos               string       `json:"outbPos"`
		StatEvetOutbDtm       string       `json:"statEvetOutbDtm"`
		StatEvetActnCntn      string       `json:"statEvetActnCntn"`
		ProcSt                string       `json:"procSt"`
		USvcOutbId            string       `json:"uSvcOutbId"`
		StatEvetItem          statEvetItem `json:"statEvetItem"`
		StatEvetActnMn        string       `json:"statEvetActnMn"`
		CpxRelEvetOutbSeqnCnt string       `json:"cpxRelEvetOutbSeqnCnt"`
		CpxRelEvetOutbSeqn    string       `json:"cpxRelEvetOutbSeqn"`
		OutbPosCnt            string       `json:"outbPosCnt"`
		StatEvetItemCnt       string       `json:"statEvetItemCnt"`
		StatEvetActnDtm       string       `json:"statEvetActnDtm"`
		StatEvetId            string       `json:"statEvetId"`
		OutbMainGb            string       `json:"outbMainGb"`
	} `json:"StatEvet"`
}

func SendMessage(originStr string) {
	//  originStr 예시 :  {"carsInOperationID":251,"carLicenseNo":"452","notifyType":2,"longitude":37.742485238266084,"latitude":126.92642211914064}
	//  carEvent := database.CarEvent{}
	//  json.Unmarshal([]byte(originStr), &carEvent)
	//  - notifyType 컬럼 : 이벤트 종류 (경로이탈:1, 도착지이탈:2, 지연도착:3, 사고+인명:4, 사고+차대차:5, 고장+기동X:6, 고장+기동O:7)
	//  - 사고 관련하여 APP에서 호출하고 있는 /api/web/sos/req API 호출시 파라미터 추가 필요 (latitude, longitude => 현재위치 좌표)

	// mrs 서버의 ip와 temporary port. port는 mrs의 system.properties에서 정함.
	conn, err := net.Dial("tcp", config.Config("MrsServer"))
	if nil != err {
		log.Fatalf("failed to connect to server")
	}
	defer conn.Close()

	// 현재시간을 스트링으로 만듬. golang 특유의 포맷 문자열에 주의. 20060102...는 자바로 치면 yyyyMMdd 에 해당.
	currentDateTimeString := strings.ReplaceAll(time.Now().Format("20060102150405.999"), ".", "")

	u2 := uuid.NewV4()
	fmt.Printf("UUIDv4a: %s\n", u2)
	fmt.Printf("%-24s", u2)

	// 헤더 a + 바디 문자열 길이 + 헤더B로 헤더를 구성함. 추후 변경값이 있으면 printf로 변경할 것.
	// 웹에서 mrs에 접속해서 헤더관리 메뉴에 가면 정의되어 있음.
	// headerA와 headerB로 나눈 이유는 사이에 바이너리 데이터가 들어감. body length의 int 값을
	// little endian 바이너리로 치환하여 중간에 삽입.
	headerA := "SMT" +
		fmt.Sprintf("%8v", config.Config("siteCD")) + // fmt.Sprintf("%8v", config.Config("siteCD"))
		"A1" +
		"      SIM" +
		" " // message exchange pattern

	t := strings.ReplaceAll(u2.String(), "-", "")
	fmt.Println("test [", t[:24])
	headerB := "001" + // message type cd baa4989a74da4e0aa1db5272
		t[:24] + // trace id
		currentDateTimeString // 20210517150516142

	// body content를 만듬.
	bodyJson := BodyJson{}
	bodyJson.StatEvet.OutbPosNm = originStr
	bodyJson.StatEvet.StatEvetGdCd = 99
	// bodyJson.StatEvet.StatEvetClrDtm = currentDateTimeString
	bodyJson.StatEvet.USvcOutbId = t[:24]
	bodyJson.StatEvet.StatEvetNm = "vms-event"
	bodyJson.StatEvet.StatEvetId = config.Config("outCarEventId") // config.Config("outCarEventId")
	bodyJson.StatEvet.StatEvetItem = []Item{{"DATA", originStr}}
	bodyJson.StatEvet.StatEvetItemCnt = "1"
	bodyJson.StatEvet.OutbPosCnt = "0"
	bodyJson.StatEvet.ProcSt = "1"
	bodyJson.StatEvet.CpxRelEvetOutbSeqnCnt = "0"
	bodyJson.StatEvet.StatEvetCntn = originStr
	bodyJson.StatEvet.StatEvetOutbDtm = currentDateTimeString
	//	bodyJson.StatEvet.OutbPos = []

	// var statEvetCntn StatEvetCntn
	// statEvetCntn = []string{"sleep", "comma", "flag"}

	// bodyJson.StatEvet.StatEvetCntn = append(bodyJson.StatEvet.StatEvetCntn, "sleep", "comma", "flag")

	// 상단의 BodyJson struct를 json으로 변환하여 바이트 배열로 받음.
	bodyByte, err := json.Marshal(bodyJson)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(bodyByte))

	// body := []byte(`{
	// 	"StatEvet": {
	// 		"statEvetGdCd": "",
	// 		"outbPosNm": "",
	// 		"statEvetNm": "",
	// 		"statEvetClrDtm": "",
	// 		"statEvetActnRslt": "",
	// 		"statEvetCntn": " []",
	// 		"outbScopRads": "",
	// 		"outbPos": [],
	// 		"statEvetOutbDtm": "20210511164324439",
	// 		"statEvetActnCntn": "",
	// 		"procSt": "",
	// 		"uSvcOutbId": "S_ej9ncQEz4cdaGj3MKfxRQU",
	// 		"statEvetItem": [],
	// 		"statEvetActnMn": "",
	// 		"cpxRelEvetOutbSeqnCnt": 0,
	// 		"cpxRelEvetOutbSeqn": [],
	// 		"outbPosCnt": 0,
	// 		"statEvetItemCnt": 0,
	// 		"statEvetActnDtm": "",
	// 		"statEvetId": "",
	// 		"outbMainGb": ""
	// 	}
	// }`)

	// 바디 byte array의 길이
	bs := make([]byte, 4)
	binary.LittleEndian.PutUint32(bs, uint32(len(bodyByte)))

	// 위에서 구한 body 의 lenght를 headerA + bodyLength + headerB로 합치면 heaer 가 된다.
	header := append([]byte(headerA), append(bs, headerB...)...)

	//	fmt.Println(string(header))
	conn.Write(header)
	conn.Write(bodyByte)

}
