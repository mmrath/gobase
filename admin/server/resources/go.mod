module github.com/mmrath/gobase/admin

go 1.12

require (
	github.com/go-chi/chi v4.0.2+incompatible
	github.com/go-chi/cors v1.0.0
	github.com/go-chi/render v1.0.1
	github.com/go-playground/locales v0.12.1 // indirect
	github.com/go-playground/universal-translator v0.16.0 // indirect
	github.com/leodido/go-urn v1.1.0 // indirect
	github.com/rs/zerolog v1.15.0
	github.com/spf13/cast v1.3.0
	github.com/spf13/cobra v0.0.5
	gopkg.in/go-playground/assert.v1 v1.2.1 // indirect
	gopkg.in/go-playground/validator.v9 v9.29.1
	github.com/mmrath/gobase/common/config v0.0.0
	github.com/mmrath/gobase/common/errors v0.0.0
	github.com/mmrath/gobase/model v0.0.0
)

replace github.com/mmrath/gobase/common/auth => ../../common/auth

replace github.com/mmrath/gobase/common/config => ../../common/config

replace github.com/mmrath/gobase/common/crypto => ../../common/crypto

replace github.com/mmrath/gobase/common/email => ../../common/email

replace github.com/mmrath/gobase/common/errors => ../../common/errors

replace github.com/mmrath/gobase/common/log => ../../common/log

replace github.com/mmrath/gobase/model => ../../model
