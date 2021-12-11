package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/etcdv3"
	opentracing2 "github.com/micro/go-plugins/wrapper/trace/opentracing"
	userproto "github.com/octopuszy/go-micro-user/proto/user"
	util "github.com/octopuszy/micro-util"
	"github.com/opentracing/opentracing-go"
	"log"
)

func main() {
	registre := etcdv3.NewRegistry(func(op *registry.Options) {
		op.Addrs = []string{"http://127.0.0.1:2379"}
	})

	// 链路追踪
	trancer, i, err := util.NewTrancer("test.client", "localhost:6831")
	if err != nil {
		return
	}
	defer i.Close()
	opentracing.SetGlobalTracer(trancer)

	ser := micro.NewService(
		micro.Name("test.client"),
		micro.Registry(registre),
		micro.WrapClient(opentracing2.NewClientWrapper(opentracing.GlobalTracer())),
	)
	ser.Init()

	// 注册服务
	user := userproto.NewUserService("test.server.user", ser.Client())

	// 请求注册接口
	//register := user_proto.RegisterReq{UserName: "zhang", Email: "1@qq.com", Password: "123"}
	//rsp, err := user.Register(context.TODO(), &register)
	//if err != nil {
	//	log.Fatal(err)
	//	return
	//}

	// 请求登录接口
	login := userproto.LoginReq{UserName: "zhang",Password: "1234"}
	rsp, err := user.Login(context.TODO(), &login)
	if err != nil {
		log.Fatal(err)
		return
	}

	marshal, _ := json.Marshal(map[string]interface{}{
		"ErrorNo:" : rsp.ErrorNo,
		"Message:" : rsp.Message,
	})
	fmt.Println(string(marshal))
}
