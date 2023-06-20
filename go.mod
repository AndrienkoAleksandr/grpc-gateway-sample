module example.com/m

go 1.18

require github.com/grpc-ecosystem/grpc-gateway/v2 v2.16.0

require (
	github.com/AndrienkoAleksandr/go v0.0.19
	github.com/golang/protobuf v1.5.3 // indirect
	golang.org/x/net v0.11.1-0.20230613203745-f5464ddb689c // indirect
	golang.org/x/sys v0.9.0 // indirect
	golang.org/x/text v0.10.1-0.20230613190012-2df65d769a9e // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20230530153820-e85fd2cbaebc // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20230530153820-e85fd2cbaebc // indirect
	google.golang.org/grpc v1.55.0 // indirect
	google.golang.org/protobuf v1.30.0 // indirect
// github.com/AndrienkoAleksandr/net v0.10.1
)

// replace golang.org/x/net => github.com/AndrienkoAleksandr/net v0.10.1
exclude golang.org/x/net v0.1.0

exclude golang.org/x/net v0.2.0

exclude golang.org/x/net v0.3.0

exclude golang.org/x/net v0.4.0

exclude golang.org/x/net v0.5.0

exclude golang.org/x/net v0.6.0

exclude golang.org/x/net v0.7.0

exclude golang.org/x/net v0.8.0

exclude golang.org/x/net v0.9.0

exclude golang.org/x/net v0.10.0

// exclude golang.org/x/net v0.11.0
