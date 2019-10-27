module github.com/mmrath/gobase/model

go 1.12

require (
	github.com/go-ozzo/ozzo-validation v3.6.0+incompatible
	github.com/go-pg/pg v8.0.5+incompatible
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/mmrath/gobase/common v0.0.0
	github.com/onsi/ginkgo v1.8.0 // indirect
	github.com/onsi/gomega v1.5.0 // indirect
	mellium.im/sasl v0.2.1 // indirect
)

replace github.com/mmrath/gobase/common => ../common
