module github.com/mmrath/gobase/uaa-server

go 1.12

require (
	github.com/go-chi/chi v4.0.2+incompatible
	github.com/go-chi/cors v1.0.0
	github.com/go-chi/render v1.0.1
	github.com/go-openapi/analysis v0.19.6 // indirect
	github.com/go-openapi/jsonreference v0.19.3 // indirect
	github.com/go-openapi/loads v0.19.4 // indirect
	github.com/go-openapi/runtime v0.19.7 // indirect
	github.com/go-openapi/spec v0.19.4 // indirect
	github.com/go-openapi/validate v0.19.4 // indirect
	github.com/go-ozzo/ozzo-validation v3.6.0+incompatible
	github.com/go-pg/pg v8.0.6+incompatible // indirect
	github.com/google/uuid v1.1.1
	github.com/google/wire v0.3.0
	github.com/mailru/easyjson v0.7.0 // indirect
	github.com/mmrath/gobase/common v0.0.0
	github.com/mmrath/gobase/model v0.0.0
	github.com/ory/hydra v1.0.8
	github.com/pelletier/go-toml v1.6.0 // indirect
	github.com/rs/zerolog v1.16.0
	github.com/spf13/cobra v0.0.5
	github.com/spf13/pflag v1.0.5 // indirect
	go.mongodb.org/mongo-driver v1.1.2 // indirect
	golang.org/x/crypto v0.0.0-20191029031824-8986dd9e96cf // indirect
	golang.org/x/net v0.0.0-20191028085509-fe3aa8a45271 // indirect
	golang.org/x/sys v0.0.0-20191028164358-195ce5e7f934 // indirect
)

replace github.com/mmrath/gobase/common => ../../common

replace github.com/mmrath/gobase/model => ../../model
