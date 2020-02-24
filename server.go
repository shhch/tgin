package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	log "github.com/sirupsen/logrus"
	"net"
	"os"
	msg "tgin/grpc"
	"tgin/tool"
)

func init()  {
	// set log output format
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)

	// set log level
	log.SetLevel(log.DebugLevel)
}

func main() {
	runGrpc()
}

func runGrpc()  {
	addr := "127.0.0.1:8080"
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listem" + addr, err)
	}
	s := grpc.NewServer()
	msg.RegisterSendServer(s, &msg.Msg{})
	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("error", err)
	}
	fmt.Println("grpc server is listen: %s", addr)
	log.Println("grpc server is listen: %s", addr)
}

func runGin()  {

	addr := "127.0.0.1:8080"

	r := gin.Default()
	r.GET("/checkbuild", func(c *gin.Context) {
		ret, err := tool.CheckBuild()
		if err != nil {
			c.JSON(200,gin.H{
				"msg":err.Error(),
			})
		}
		c.JSON(200,gin.H{
			"result":ret,
		})
	})
	r.GET("/build", func(c *gin.Context) {
		tool.P.Password = c.Param("password")
		tool.P.Client = c.Param("client")
		tool.P.Ip = c.Param("ip")
		tool.P.Swoole = c.Param("swoole")
		tool.P.Platform = c.Param("platform")
		tool.P.Datrixlog = c.Param("datrixlog")
		tool.P.Web = c.Param("web")

		result := tool.Build()
		c.JSON(200, result)

	})

	err := r.Run(addr)
	if err != nil {
		fmt.Println("Error :" + err.Error())
	}else {
		fmt.Println("服务启动，监听：" + addr)
	}
}

