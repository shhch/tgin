package tool

import (
	"bytes"
	"go/types"
	"log"
	"net"
	"strings"
)

type BuildParam struct {
	Password string
	Ip string
	Swoole string
	Platform string
	Web string
	Client string
	Datrixlog string
}

var P = &BuildParam{}

func CheckBuild() (string, error) {
	command := "ps aux|grep /data/git-develop-build/|grep -v grep|awk '{print $2}'"
	output, err := RunCmd(command)
	return output, err
}

func Build() (ret map[string]interface{}) {
	ret = make(map[string]interface{})
	cmdBuffer := bytes.Buffer{}
	if P.Password != "buildpw" {
		ret["success"] = false
		ret["msg"] = "密码错误"
		return
	}
	if P.Swoole != "" {
		cmdBuffer.WriteString(" -s ")
	}
	if P.Platform != "" {
		cmdBuffer.WriteString(" -p ")
	}
	if P.Datrixlog != "" {
		cmdBuffer.WriteString(" -l ")
	}
	if P.Client != "" {
		cmdBuffer.WriteString(" -c ")
	}
	ip := P.Ip
	if ip == "" {
		ret["success"] = false
		ret["msg"] = "ip不能为空"
		return
	}else {
		ipArr := strings.Split(ip, ",")
		for _, val := range ipArr{
			if addr := net.ParseIP(val); addr == nil {
				ret["success"] = false
				ret["msg"] = "ip" + val + "有误"
				return
			}
		}
		cmdBuffer.WriteString(" -i  " + ip)
	}
	if P.Web != "" {
		cmdBuffer.WriteString(" -w " + P.Web)
	}

	cmd := "/data/git-develop-build/build.sh " + cmdBuffer.String()
	output, err := RunCmd(cmd)
	if err != nil {
		ret["success"] = false
		ret["msg"] = "ip不能为空"
	}
	return
}
