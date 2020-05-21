package utils

import (
	"encoding/base64"
	"io/ioutil"
	"strings"
)

func GetTlsCrt(secret string, namespace string)(err error) {
	command :=`kubectl get secret ` + secret + ` -n ` + namespace + ` -o json|jq '.data."tls.crt"'`
	output, err := Cmd(command)
	if err != nil {
		return err
	}
	tlsCrtBase64 := strings.Replace(string(output),"\"","",-1)
	decodeBytes, err := base64.StdEncoding.DecodeString(tlsCrtBase64)
	if err != nil {
		return err
	}
	var data = []byte(decodeBytes)
	err2 := ioutil.WriteFile(secret, data, 0666)
	if err != nil {
		return err2
	}
	return nil
}
