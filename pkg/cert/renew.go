package cert

import (
	"fmt"
	"renewCertificate/pkg/utils"
	"strings"
)

func Renew(certificate string,namespace string) (err error) {
	certificates := strings.Split(certificate, ",")
	for cert := range certificates {
		fmt.Println("-------------------------------------------------------------")
		fmt.Println(certificates[cert])
		certificate := certificates[cert]
		err := utils.GetTlsCrt(certificate,namespace)
		if err != nil {
			fmt.Println(err)
			return err
		}
		timeRemaining, err := GetTimeRemaining(certificate)
		if err != nil {
			fmt.Println(err)
			return err
		}
		if timeRemaining != -100 && timeRemaining <= 3 {
			command := `kubectl cert-manager renew --namespace ` + namespace + ` ` + certificate
			output, err := utils.Cmd(command)
			if err != nil {
				fmt.Println(err)
				return err
			}
			fmt.Println(string(output))
		}
	}
	return nil
}
