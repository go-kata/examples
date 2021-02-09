package system

import (
	"os"
	"path/filepath"

	"github.com/go-kata/kinit/kinitx"
)

func init() { kinitx.MustProvide(New) }

type System struct {
	CurrentDir    string
	ExecutableDir string
	Command       string
}

func New() (*System, error) {
	s := &System{}
	var err error
	if s.CurrentDir, err = os.Getwd(); err != nil {
		return nil, err
	}
	executablePath, err := os.Executable()
	if err != nil {
		return nil, err
	}
	s.ExecutableDir = filepath.Dir(executablePath)
	executableName := filepath.Base(executablePath)
	s.Command = executableName[0 : len(executableName)-len(filepath.Ext(executableName))]
	return s, nil
}
