package greeter

import (
	"fmt"
	"log"

	"github.com/go-kata/kdone"
	"github.com/go-kata/kinit/kinitx"
)

func init() { kinitx.MustProvide(New) }

type Greeter func(string)

func New(config *Config, logger *log.Logger) (Greeter, kdone.Destructor, error) {
	logger.Print("initializing greeter")
	g := func(name string) {
		fmt.Println()
		fmt.Printf(config.Format, name)
		fmt.Println()
	}
	dtor := kdone.DestructorFunc(func() error {
		logger.Print("finalizing greeter")
		return nil
	})
	return g, dtor, nil
}
