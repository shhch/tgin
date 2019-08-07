package tool

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"os/exec"
	"strings"
	"bytes"
	"encoding/json"
	"net/http"
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
	// 无需执行Run或者Start，OutPut中已经封装完成
	output, err := cmd.Output()
	return string(output), err
}

func HttpReq(method string, url string, params map[string]interface{}){

	byteBuf, err := json.Marshal(params)
	if err != nil {
		fmt.Println(err.Error())
	}
	headerStr := "application/json"

	req, err := http.NewRequest(method,url,bytes.NewBuffer(byteBuf))
	if err != nil {
		fmt.Println(err.Error())
	}

	req.Header.Set("Content-Type", headerStr)

	client := &http.Client{}
	resp, err := client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		fmt.Println(err.Error())
	}

	buffer := new(bytes.Buffer)
	buffer.ReadFrom(resp.Body)

	var response map[string]interface{}
	err = json.Unmarshal(buffer.Bytes(), &response)

	fmt.Println(&response)

}
