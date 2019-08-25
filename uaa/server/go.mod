module mmrath.com/gobase/uaa

go 1.12

require (
	github.com/go-chi/chi v4.0.2+incompatible
	github.com/go-chi/cors v1.0.0
	github.com/go-chi/render v1.0.1
	github.com/spf13/cobra v0.0.5
	mmrath.com/gobase/common/config v0.0.0
	mmrath.com/gobase/model v0.0.0
)

replace mmrath.com/gobase/common/auth => ../../common/auth

replace mmrath.com/gobase/common/config => ../../common/config

replace mmrath.com/gobase/common/crypto => ../../common/crypto

replace mmrath.com/gobase/common/email => ../../common/email

replace mmrath.com/gobase/common/errors => ../../common/errors

replace mmrath.com/gobase/common/log => ../../common/log

replace mmrath.com/gobase/model => ../../model
