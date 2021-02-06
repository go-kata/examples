package hello

import (
	"log"
	"os"

	"github.com/go-kata/kinit"
	"github.com/go-kata/kinitx"

	"github.com/go-kata/examples/cmd/example-di/greeter"
)

func Func() (kinit.Functor, error) {
	return kinitx.NewFunctor(func(logger *log.Logger, greet greeter.Greeter) {
		logger.Print("entering hello")
		defer logger.Print("exiting hello")
		var name string
		if len(os.Args) > 2 {
			name = os.Args[2]
		} else {
			name = "anonymous"
		}
		greet(name)
	})
}
