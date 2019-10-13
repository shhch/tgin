package main

import (
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	msg "tgin/grpc"
)

const (
	addr = "127.0.0.1:8080"
)

func main()  {
	cn, _ := grpc.Dial(addr)
	defer cn.Close()
	client := msg.NewSendClient(cn)

	var req msg.SendRequest
	req.Type = "测试type"
	req.Tag = 123
	req.Msg = "测试msg"
	//&msg.SendRequest{
	//	Msg:"测试msg",
	//	Tag:123,
	//	Type:"测试type",
	//}
	rep, err := client.Send(context.Background(), &req)
	if err != nil {
		log.Fatalf("grpc Error: ", err.Error())
	}
	fmt.Println(rep.Tag,rep.Msg,rep.Result)
}


