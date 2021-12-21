package commands

import (
	"errors"
	"os"
)

const (
	root = 0
)

func AskRoot() error {
	if os.Getegid() == root {
		return nil
	}

	return errors.New("run the application with root privileges")
}
