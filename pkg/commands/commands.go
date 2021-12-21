/*
Package commands keeps all the os commands.
*/
package commands

import (
	"errors"
	"os"
	"os/exec"
)

const (
	root = 0
)

func CreateSnapshot(sv, path string) error {
	return exec.Command("btrfs", "subvolume", "snapshot", sv, path).Run()
}

func DeleteSnapshot(path string) error {
	return exec.Command("btrfs", "subvolume", "delete", path).Run()
}

func AskRoot() error {
	if os.Getegid() == root {
		return nil
	}

	return errors.New("run the application with root privileges")
}
