package main

import (
	"log"
	"os"

	"github.com/go-kata/kerror"
	"github.com/go-kata/kinit"
	"github.com/go-kata/kinitx"

	"github.com/go-kata/examples/cmd/example-di/command/hello"
)

func main() { kinitx.MustInvoke(Main) }

func Main(_ *Config, logger *log.Logger) (kinit.Executor, error) {
	logger.Println("entering main executor")
	defer logger.Println("exiting main executor")
	if len(os.Args) > 1 {
		switch os.Args[1] {
		default:
			return nil, kerror.Newf(kerror.EInvalid, "invalid command: %s", os.Args[1])
		case "hello":
			return hello.Exec(), nil
		}
	}
	Help()
	return nil, nil
}
