module mmrath.com/gobase/admin

go 1.12

require (
	github.com/go-chi/chi v4.0.2+incompatible
	github.com/go-chi/render v1.0.1
	github.com/go-playground/locales v0.12.1 // indirect
	github.com/go-playground/universal-translator v0.16.0 // indirect
	github.com/leodido/go-urn v1.1.0 // indirect
	github.com/rs/zerolog v1.15.0
	gopkg.in/go-playground/assert.v1 v1.2.1 // indirect
	gopkg.in/go-playground/validator.v9 v9.29.1
	gopkg.in/yaml.v2 v2.2.2 // indirect
	mmrath.com/gobase/common/errors v0.0.0
	mmrath.com/gobase/model v0.0.0-00010101000000-000000000000
)

replace mmrath.com/gobase/common/auth => ../../common/auth

replace mmrath.com/gobase/common/crypto => ../../common/crypto

replace mmrath.com/gobase/common/email => ../../common/email

replace mmrath.com/gobase/common/errors => ../../common/errors

replace mmrath.com/gobase/common/log => ../../common/log

replace mmrath.com/gobase/model => ../../model
