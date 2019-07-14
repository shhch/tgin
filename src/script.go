package src

import (
	"fmt"
	"os/exec"
)

func checkBuild() {
	fmt.Println("adasdas")
}

func runCmd(command string, arg string) (bool, string) {
	cmd := exec.Command(command, arg)
	output, err := cmd.Output()
	if err != nil {
		return false, string(output)
	}
	if output != nil{
		return false, string(output)
	}
	return true, string(output)
}
