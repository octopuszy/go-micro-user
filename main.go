package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/etcdv3"
	opentracing2 "github.com/micro/go-plugins/wrapper/trace/opentracing"
	user_repository "github.com/octopuszy/go-micro-user/domain/repository/user"
	user_service "github.com/octopuszy/go-micro-user/domain/service/user"
	"github.com/octopuszy/go-micro-user/handler"
	user_proto "github.com/octopuszy/go-micro-user/proto/user"
	util "github.com/octopuszy/micro-util"
	"github.com/opentracing/opentracing-go"
	"log"
)

func main() {
	// 注册中心
	registre := etcdv3.NewRegistry(func(op *registry.Options) {
		op.Addrs = []string{"http://127.0.0.1:2379"}
	})

	// 链路追踪
	trancer, i, err := util.NewTrancer("test.server.user", "localhost:6831")
	if err != nil {
		return
	}
	defer i.Close()
	opentracing.SetGlobalTracer(trancer)

	srv := micro.NewService(
		micro.Name("test.server.user"),
		micro.Version("latest"),
		micro.Registry(registre),
		micro.WrapHandler(opentracing2.NewHandlerWrapper(opentracing.GlobalTracer())),
	)
	srv.Init()

	// 创建数据库连接
	db, err := gorm.Open("mysql","root:123456@tcp(127.0.0.1:3306)/zhangyu")
	if err != nil {
		log.Fatal(err);
		return 
	}
	defer db.Close()

	// 设置禁用表名复数形式
	db.SingularTable(true)

	// 创建 repository
	repository := user_repository.NewUserRepository(db)

	// 创建 service
	service := user_service.NewUserService(repository)

	// 创建 handler
	user_handler := handler.User{UserService: service}

	// 注册服务
	err = user_proto.RegisterUserHandler(srv.Server(), &user_handler)
	if err != nil {
		log.Fatal(err);
		return
	}

	if err := srv.Run(); err != nil {
		log.Fatal(err);
	}
}
