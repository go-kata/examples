package greeter

import (
	"github.com/go-kata/kinit"
	"github.com/go-kata/kinitx"
)

var _ = kinit.MustDeclare(func() { kinitx.MustProvide(NewConfig) })

type Config struct {
	Format string
}

func NewConfig() *Config {
	return &Config{
		Format: "Hello, %s!\n",
	}
}
