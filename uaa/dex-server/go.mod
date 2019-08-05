module mmrath.com/gobase/uaa

go 1.12

require (
	github.com/dexidp/dex v0.0.0-20190803115620-526e07836656
	github.com/ghodss/yaml v1.0.0
	github.com/grpc-ecosystem/go-grpc-prometheus v1.2.0
	github.com/kylelemons/godebug v1.1.0
	github.com/prometheus/client_golang v1.1.0
	github.com/spf13/cobra v0.0.5
	golang.org/x/crypto v0.0.0-20190701094942-4def268fd1a4
	golang.org/x/sys v0.0.0-20190804053845-51ab0e2deafa // indirect
	google.golang.org/genproto v0.0.0-20190801165951-fa694d86fc64 // indirect
	google.golang.org/grpc v1.22.1
	mmrath.com/gobase/common/log v0.0.0
)

replace mmrath.com/gobase/common/log => ../../common/log
