package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/go-kata/kinit"
	"github.com/go-kata/kinitx"

	"github.com/go-kata/examples/cmd/example-di/greeter"
	"github.com/go-kata/examples/cmd/example-di/logger"
	"github.com/go-kata/examples/cmd/example-di/system"
)

var _ = kinit.MustDeclare(func() {
	kinitx.MustProvide((*Config)(nil))
	kinitx.MustApply((*Config).Load)
})

type Config struct {
	Logger  *logger.Config
	Greeter *greeter.Config
}

func (c *Config) Load(sys *system.System) error {
	path := filepath.Join(sys.ExecutableDir, sys.Command+".json")
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			fmt.Printf("WARN: no configuration found, defaults used\n")
			return nil
		}
		return err
	}
	fmt.Printf("INFO: configuration found at %s\n", path)
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, c)
}
