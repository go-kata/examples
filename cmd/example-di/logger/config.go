package logger

import (
	"github.com/go-kata/kinit/kinitx"

	"github.com/go-kata/examples/cmd/example-di/system"
)

func init() { kinitx.MustProvide(NewConfig) }

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
