package utils

import (
	"fmt"
	"os"
	"os/exec"
)

func Cmd(command string) ([]byte, error){
	fmt.Println("cmd:", command)
	cmd := exec.Command("/bin/bash", "-c", command)
	//cmd := exec.Command("kubectl", "get", "secret", "notary-cert", "-o", "json")
	//cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	output,err := cmd.Output()
	if err != nil {
		cmd.Stderr = os.Stderr
		fmt.Println("Execute Command failed:" + err.Error())
		return nil,err
	}
	return output,nil
}
