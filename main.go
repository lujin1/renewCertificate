package main

import (
	"flag"
	"fmt"
	"renewCertificate/pkg/cert"
)

var namespace = flag.String("namespace","","namespace")
var certificate = flag.String("certificate","","notary-cert,harbor-cert")

func main() {
	//certificate := "notary-cert,harbor-cert"
	//namesapce := "harbor"
	flag.Parse()
	fmt.Println("namespace:", *namespace)
	fmt.Println("certificate:", *certificate)
	if *namespace != "" && *certificate != "" {
		cert.Renew(*certificate, *namespace)
	} else {
		fmt.Println("ERROR: namespace or certificate is null ")
	}
}
