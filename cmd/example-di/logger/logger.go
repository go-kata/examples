package logger

import (
	"log"
	"os"

	"github.com/go-kata/kdone"
	"github.com/go-kata/kinit/kinitx"

	"github.com/go-kata/examples/cmd/example-di/system"
)

func init() { kinitx.MustProvide(New) }

func New(config *Config, sys *system.System) (*log.Logger, kdone.Destructor, error) {
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
	l := log.New(os.Stderr, config.Prefix, flags)
	l.Printf("%s is up", sys.Command)
	dtor := kdone.DestructorFunc(func() error {
		l.Printf("%s is down", sys.Command)
		return nil
	})
	return l, dtor, nil
}
