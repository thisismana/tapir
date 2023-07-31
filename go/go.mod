module github.com/stroeer/tapir/go

go 1.17

require (
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.16.2
	github.com/stroeer/go-tapir v0.0.0-20210901143243-2da66482ca62
	google.golang.org/genproto v0.0.0-20230706204954-ccb25ca9f130 // indirect
	google.golang.org/grpc v1.57.0 // indirect
	google.golang.org/grpc/cmd/protoc-gen-go-grpc v1.1.0
	google.golang.org/protobuf v1.31.0
	gopkg.in/yaml.v2 v2.4.0 // indirect
)

require (
	github.com/golang/glog v1.1.0 // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	golang.org/x/net v0.12.0 // indirect
	golang.org/x/sys v0.10.0 // indirect
	golang.org/x/text v0.11.0 // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20230726155614-23370e0ffb3e // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20230726155614-23370e0ffb3e // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	sigs.k8s.io/yaml v1.3.0 // indirect
)

// replace go-tapir dependency with generated module
replace github.com/stroeer/go-tapir => ../gen/go
