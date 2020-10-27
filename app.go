package go_envelope

import (
	"go-envelope/infra"
	"go-envelope/infra/base"
)

func init() {
	infra.Register(&base.PropsStarter{})
}
