module github.com/mmrath/gobase/uaa-server

go 1.12

require (
	github.com/go-chi/chi v4.0.2+incompatible
	github.com/go-chi/cors v1.0.0
	github.com/go-chi/render v1.0.1
	github.com/mmrath/gobase/common v0.0.0
	github.com/mmrath/gobase/model v0.0.0
	github.com/ory/hydra v1.0.8
	github.com/ory/x v0.0.76
	github.com/rs/zerolog v1.15.0
	github.com/spf13/cobra v0.0.5
)

replace github.com/mmrath/gobase/common => ../../common

replace github.com/mmrath/gobase/model => ../../model
