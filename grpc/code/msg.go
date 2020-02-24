package code

import (
	"fmt"
	"golang.org/x/net/context"
)

type Msg struct {

}

func (m *Msg) Send(ctx context.Context, request *SendRequest) (response *SendResponse, err error) {
	fmt.Println(request.Msg,request.Tag,request.Type)
	response = &SendResponse{
		Result:true,
		Tag:request.Tag,
		Msg:request.Msg,
	}
	return response, nil
}
