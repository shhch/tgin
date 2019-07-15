package tool

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"os/exec"
	"strings"
)

func GetParam(c gin.Context)  {
	dataType := strings.ToUpper(c.ContentType())
	switch dataType {
	case "Json":
		fmt.Println("json")
	case "Form-Data":
		fmt.Println("form-data")
	default:
		fmt.Println("Nothing")
	}
	req := c.Request.Header.Get("Method")
	switch req {
	case "POST":
		fmt.Println("POST")
	default:
		fmt.Println(req)
	}
}

func RunCmd(command string, arg string) (bool, string) {
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
