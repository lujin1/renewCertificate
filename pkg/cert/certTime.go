package cert

import (
	"fmt"
	"renewCertificate/pkg/utils"
	"strings"
	"time"
)

func GetTimeRemaining(certfile string)(float64, error) {
	command := `openssl x509 -in ` + certfile +` -noout -enddate`
	output, err := utils.Cmd(command)
	if err != nil {
		return -100, err
	}
	notAfter := strings.Split(string(output), "=")
	//fmt.Println("endTime: " + notAfter[1])
	notAfterSplit := strings.Split(notAfter[1], " ")
	monthStr := notAfterSplit[0]
	dayStr := notAfterSplit[1]
	timeStr := notAfterSplit[2]
	yearStr := notAfterSplit[3]
	dateStr := utils.TimeTransform(yearStr, monthStr, dayStr, timeStr)
	//fmt.Println("endTime: " + dateStr)
	endTime,err :=time.Parse("2006-01-02 15:04:05",dateStr)
	if err != nil {
		return -100, err
		fmt.Println("解析错误:",err)
	}
	fmt.Println("endTime:",endTime)
	n := time.Now().UTC()
	fmt.Println("now:", n)
	dd:=endTime.Sub(n)
	timeRemaining := dd.Hours()/24
	fmt.Println("timeRemaining:",timeRemaining,"days")
	return timeRemaining, nil
}
