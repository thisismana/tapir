module github.com/stroeer/tapir/go

go 1.17

require (
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.14.0
	github.com/stroeer/go-tapir v0.0.0-20210901143243-2da66482ca62
	google.golang.org/genproto v0.0.0-20221114212237-e4508ebdbee1 // indirect
	google.golang.org/grpc v1.50.1 // indirect
	google.golang.org/grpc/cmd/protoc-gen-go-grpc v1.1.0
	google.golang.org/protobuf v1.28.1
	gopkg.in/yaml.v2 v2.4.0 // indirect
)

require (
	github.com/golang/glog v1.0.0 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	golang.org/x/net v0.2.0 // indirect
	golang.org/x/sys v0.2.0 // indirect
	golang.org/x/text v0.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	sigs.k8s.io/yaml v1.3.0 // indirect
)

// replace go-tapir dependency with generated module
replace github.com/stroeer/go-tapir => ../gen/go
