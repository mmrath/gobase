module mmrath.com/gobase/uaa/api

go 1.12

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/go-session/session v3.1.2+incompatible
	gopkg.in/oauth2.v3 v3.10.0
	mmrath.com/gobase/common/crypto v0.0.0
	mmrath.com/gobase/common/errors v0.0.0
	mmrath.com/gobase/model v0.0.0-00010101000000-000000000000
)

replace mmrath.com/gobase/common/errors => ../../common/errors

replace mmrath.com/gobase/common/crypto => ../../common/crypto

replace mmrath.com/gobase/model => ../../model
