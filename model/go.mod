module mmrath.com/gobase/model

go 1.12

require (
	github.com/go-ozzo/ozzo-validation v3.5.0+incompatible
	github.com/go-pg/pg v8.0.5+incompatible
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/kr/pretty v0.1.0 // indirect
	github.com/onsi/ginkgo v1.8.0 // indirect
	github.com/onsi/gomega v1.5.0 // indirect
	gopkg.in/check.v1 v1.0.0-20180628173108-788fd7840127 // indirect
	mellium.im/sasl v0.2.1 // indirect
	mmrath.com/gobase/common/errors v0.0.0
)

replace mmrath.com/gobase/common/errors => ../common/errors
