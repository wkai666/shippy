module shippy

go 1.14

require (
	github.com/golang/protobuf v1.5.2
	github.com/micro/go-micro/v2 v2.9.1
	github.com/micro/go-plugins/broker/nsq/v2 v2.9.1
	google.golang.org/grpc v1.42.0
	google.golang.org/protobuf v1.27.1
)

replace google.golang.org/grpc v1.42.0 => google.golang.org/grpc v1.26.0
