package main

import (
	"fmt"
	"log"
	"os"

	"github.com/go-kata/kerror"
	"github.com/go-kata/kinit"
	"github.com/go-kata/kinit/kinitx"

	"github.com/go-kata/examples/cmd/example-di/command/hello"
)

func init() { kinitx.MustConsider(Func) }

func Func(_ *Config, logger *log.Logger) (kinit.Functor, error) {
	logger.Print("entering main")
	defer logger.Print("exiting main")
	if len(os.Args) > 1 {
		switch os.Args[1] {
		default:
			return nil, kerror.Newf(kerror.ECustom, "invalid command: %s", os.Args[1])
		case "hello":
			return kinitx.NewFunctor(hello.Func)
		}
	}
	fmt.Println()
	fmt.Println("hello <name>\t\tSay hello")
	fmt.Println()
	return nil, nil
}
