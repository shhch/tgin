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

// command为需要执行的命令，可为脚本，可为shell命令，返回执行的输出
// 该方法为阻塞调用，调用完成后返回结果
func RunCmd(command string) (string, error ) {
	cmd := exec.Command("/bin/sh", "-c", command)
	// Stdout与Stderr用于显示正常输出与错误输出
	// cmd.Stdout = os.Stdout
	// cmd.Stderr = os.Stderr
	output, err := cmd.Output()
	return string(output), err
}
