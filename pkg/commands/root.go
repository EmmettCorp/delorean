/*
Package commands keeps all cli prompt tools commands.
*/
package commands

import (
	"bufio"
	"errors"
	"os"
	"strings"
)

const (
	root      = 0
	deviceIdx = 0
	pathIdx   = 1
)

func CheckIfRoot() error {
	if os.Getegid() == root {
		return nil
	}

	return errors.New("run the application with root privileges")
}

func GetRootDevice() (string, error) {
	fp, err := os.Open("/proc/self/mounts")
	if err != nil {
		return "", err
	}
	defer fp.Close()

	var device string

	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		if fields[pathIdx] != "/" {
			continue
		}

		device = fields[deviceIdx]

		break
	}

	if scanner.Err() != nil {
		return "", scanner.Err()
	}

	return device, nil
}
