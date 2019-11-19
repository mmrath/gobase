module github.com/mmrath/gobase/uaa-server

go 1.12

require (
	github.com/andybalholm/cascadia v1.1.0 // indirect
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
	github.com/mattn/go-runewidth v0.0.6 // indirect
	github.com/mmrath/gobase/common v0.0.0
	github.com/mmrath/gobase/model v0.0.0
	github.com/olekukonko/tablewriter v0.0.2 // indirect
	github.com/ory/hydra v1.0.9
	github.com/pelletier/go-toml v1.6.0 // indirect
	github.com/rs/zerolog v1.16.0
	github.com/spf13/cobra v0.0.5
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/spf13/viper v1.5.0 // indirect
	go.mongodb.org/mongo-driver v1.1.3 // indirect
	golang.org/x/crypto v0.0.0-20191106202628-ed6320f186d4 // indirect
	golang.org/x/net v0.0.0-20191105084925-a882066a44e0 // indirect
	golang.org/x/sys v0.0.0-20191105231009-c1f44814a5cd // indirect
	gopkg.in/yaml.v2 v2.2.5 // indirect
)

replace github.com/mmrath/gobase/common => ../../common

replace github.com/mmrath/gobase/model => ../../model
