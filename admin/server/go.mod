module mmrath.com/gobase/admin

go 1.12

require (
	github.com/onsi/ginkgo v1.8.0 // indirect
	github.com/onsi/gomega v1.5.0 // indirect
	github.com/spf13/cobra v0.0.5
	mmrath.com/gobase/common/errors v0.0.0-00010101000000-000000000000
	mmrath.com/gobase/model v0.0.0-00010101000000-000000000000
)

replace mmrath.com/gobase/common/auth => ../../common/auth

replace mmrath.com/gobase/common/crypto => ../../common/crypto

replace mmrath.com/gobase/common/email => ../../common/email

replace mmrath.com/gobase/common/errors => ../../common/errors

replace mmrath.com/gobase/common/log => ../../common/log

replace mmrath.com/gobase/model => ../../model
