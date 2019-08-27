module mmrath.com/gobase/client

go 1.12

require (
	cloud.google.com/go v0.43.0 // indirect
	github.com/OneOfOne/xxhash v1.2.5 // indirect
	github.com/ajg/form v1.5.1 // indirect
	github.com/alecthomas/template v0.0.0-20190718012654-fb15b899a751 // indirect
	github.com/alecthomas/units v0.0.0-20190717042225-c3de453c63f4 // indirect
	github.com/coreos/bbolt v1.3.3 // indirect
	github.com/coreos/etcd v3.3.13+incompatible // indirect
	github.com/coreos/go-semver v0.3.0 // indirect
	github.com/coreos/go-systemd v0.0.0-20190719114852-fd7a80b32e1f // indirect
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/dgryski/go-sip13 v0.0.0-20190329191031-25c5027a8c7b // indirect
	github.com/fatih/structs v1.1.0 // indirect
	github.com/gavv/httpexpect v2.0.0+incompatible
	github.com/go-chi/chi v4.0.2+incompatible
	github.com/go-chi/cors v1.0.0
	github.com/go-chi/jwtauth v3.3.0+incompatible
	github.com/go-chi/render v1.0.1
	github.com/go-kit/kit v0.9.0 // indirect
	github.com/go-ozzo/ozzo-validation v3.5.0+incompatible
	github.com/golang/groupcache v0.0.0-20190702054246-869f871628b6 // indirect
	github.com/google/go-querystring v1.0.0 // indirect
	github.com/google/pprof v0.0.0-20190723021845-34ac40c74b70 // indirect
	github.com/google/uuid v1.1.1
	github.com/google/wire v0.3.0
	github.com/grpc-ecosystem/grpc-gateway v1.9.5 // indirect
	github.com/hashicorp/golang-lru v0.5.3 // indirect
	github.com/imkira/go-interpol v1.1.0 // indirect
	github.com/kisielk/errcheck v1.2.0 // indirect
	github.com/klauspost/compress v1.7.5 // indirect
	github.com/klauspost/cpuid v1.2.1 // indirect
	github.com/konsorten/go-windows-terminal-sequences v1.0.2 // indirect
	github.com/kr/pty v1.1.8 // indirect
	github.com/magiconair/properties v1.8.1 // indirect
	github.com/moul/http2curl v1.0.0 // indirect
	github.com/mwitkow/go-conntrack v0.0.0-20190716064945-2f068394615f // indirect
	github.com/pelletier/go-toml v1.4.0 // indirect
	github.com/prometheus/client_golang v1.1.0 // indirect
	github.com/prometheus/tsdb v0.10.0 // indirect
	github.com/rogpeppe/fastuuid v1.2.0 // indirect
	github.com/rs/zerolog v1.15.0
	github.com/sergi/go-diff v1.0.0 // indirect
	github.com/spaolacci/murmur3 v1.1.0 // indirect
	github.com/spf13/afero v1.2.2 // indirect
	github.com/spf13/jwalterweatherman v1.1.0 // indirect
	github.com/spf13/viper v1.4.0
	github.com/stretchr/objx v0.2.0 // indirect
	github.com/stretchr/testify v1.3.0
	github.com/ugorji/go v1.1.7 // indirect
	github.com/valyala/fasthttp v1.4.0 // indirect
	github.com/xeipuuv/gojsonpointer v0.0.0-20180127040702-4e3ac2762d5f // indirect
	github.com/xeipuuv/gojsonreference v0.0.0-20180127040603-bd5ef7bd5415 // indirect
	github.com/xeipuuv/gojsonschema v1.1.0 // indirect
	github.com/yalp/jsonpath v0.0.0-20180802001716-5cc68e5049a0 // indirect
	github.com/yudai/gojsondiff v1.0.0 // indirect
	github.com/yudai/golcs v0.0.0-20170316035057-ecda9a501e82 // indirect
	go.etcd.io/bbolt v1.3.3 // indirect
	golang.org/x/exp v0.0.0-20190731235908-ec7cb31e5a56 // indirect
	golang.org/x/image v0.0.0-20190802002840-cff245a6509b // indirect
	golang.org/x/mobile v0.0.0-20190719004257-d2bd2a29d028 // indirect
	golang.org/x/sys v0.0.0-20190804053845-51ab0e2deafa // indirect
	golang.org/x/tools v0.0.0-20190802220118-1d1727260058 // indirect
	google.golang.org/genproto v0.0.0-20190801165951-fa694d86fc64 // indirect
	google.golang.org/grpc v1.22.1 // indirect
	honnef.co/go/tools v0.0.1-2019.2.2 // indirect
	mmrath.com/gobase/common/auth v0.0.0-00010101000000-000000000000
	mmrath.com/gobase/common/config v0.0.0-00010101000000-000000000000
	mmrath.com/gobase/common/crypto v0.0.0-00010101000000-000000000000
	mmrath.com/gobase/common/email v0.0.0-00010101000000-000000000000
	mmrath.com/gobase/common/errors v0.0.0
	mmrath.com/gobase/common/log v0.0.0-00010101000000-000000000000
	mmrath.com/gobase/model v0.0.0-00010101000000-000000000000
)

replace mmrath.com/gobase/common/auth => ../common/auth

replace mmrath.com/gobase/common/config => ../common/config

replace mmrath.com/gobase/common/crypto => ../common/crypto

replace mmrath.com/gobase/common/email => ../common/email

replace mmrath.com/gobase/common/errors => ../common/errors

replace mmrath.com/gobase/common/log => ../common/log

replace mmrath.com/gobase/model => ../model
