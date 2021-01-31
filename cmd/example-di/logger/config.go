package logger

import (
	"github.com/go-kata/kinit"
	"github.com/go-kata/kinitx"

	"github.com/go-kata/examples/cmd/example-di/system"
)

var _ = kinit.MustHook(func() { kinitx.MustProvide(NewConfig) })

type Config struct {
	Prefix       string
	DateTime     bool
	Microseconds bool
	UTC          bool
	File         bool
}

func NewConfig(sys *system.System) *Config {
	return &Config{
		Prefix:   sys.Command + " ",
		DateTime: true,
	}
}
