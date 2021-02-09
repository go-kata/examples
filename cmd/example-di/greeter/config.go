package greeter

import "github.com/go-kata/kinit/kinitx"

func init() { kinitx.MustProvide(NewConfig) }

type Config struct {
	Format string
}

func NewConfig() *Config {
	return &Config{
		Format: "Hello, %s!\n",
	}
}
