package logger

import (
	"log"
	"os"

	"github.com/go-kata/kinit"
	"github.com/go-kata/kinitx"
)

var _ = kinit.MustHook(func() { kinitx.MustProvide(New) })

func New(config *Config) *log.Logger {
	var flags int
	if config.File {
		flags |= log.Llongfile
	}
	if config.DateTime {
		flags |= log.Ldate | log.Ltime
	}
	if config.Microseconds {
		flags |= log.Lmicroseconds
	}
	if config.UTC {
		flags |= log.LUTC
	}
	return log.New(os.Stderr, config.Prefix, flags)
}
