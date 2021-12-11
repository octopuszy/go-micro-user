module github.com/octopuszy/go-micro-user

go 1.14

require (
	github.com/go-sql-driver/mysql v1.6.0
	github.com/golang/protobuf v1.5.2
	github.com/jinzhu/gorm v1.9.16
	github.com/micro/go-micro v1.18.0
	github.com/micro/go-plugins/registry/etcdv3 v0.0.0-20200119172437-4fe21aa238fd
	github.com/micro/go-plugins/wrapper/trace/opentracing v0.0.0-20200119172437-4fe21aa238fd // indirect
	github.com/octopuszy/micro-util v0.0.0-20211211085542-a75248e9bc57 // indirect
	github.com/opentracing/opentracing-go v1.2.0
	github.com/uber/jaeger-client-go v2.30.0+incompatible
	github.com/uber/jaeger-lib v2.4.1+incompatible // indirect
	go.etcd.io/etcd/v3 v3.5.1 // indirect
	golang.org/x/crypto v0.0.0-20211209193657-4570a0811e8b
	google.golang.org/grpc/examples v0.0.0-20211130233114-c2bccd0b1594 // indirect
	google.golang.org/protobuf v1.27.1
)

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0
