package main

import (
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"tgin/grpc/code"
)

const (
	addr = "127.0.0.1:8080"
)

// PerRPCCredentials是一个接口，需要实现它的函数
type customCredential struct{}

// 实现GetRequestMetadata
func (c customCredential) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{
		"appid":  "101010",
		"appkey": "i am key",
	}, nil
}

// 实现RequireTransportSecurity
func (c customCredential) RequireTransportSecurity() bool {
	return false
}

func main() {

	// 客户端添加metadata认证
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	opts = append(opts, grpc.WithBlock())
	// 使用自定义认证
	opts = append(opts, grpc.WithPerRPCCredentials(new(customCredential)))

	cn, _ := grpc.Dial(addr, opts...)
	defer cn.Close()
	client := code.NewSendClient(cn)

	var req code.SendRequest
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
	fmt.Println(rep.Tag, rep.Msg, rep.Result)
}
