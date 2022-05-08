/*
Package commands keeps all cli prompt tools commands.
*/
package commands

import (
	"bufio"
	"errors"
	"os"
	"strings"

	"github.com/EmmettCorp/delorean/pkg/logger"
)

const (
	root      = 0
	deviceIdx = 0
	pathIdx   = 1
)

// CheckIfRoot checks if the application is runned with root privileges.
func CheckIfRoot() error {
	if os.Getegid() == root {
		return nil
	}

	return errors.New("run the application with root privileges")
}

// GetRootDevice returns the path to root device.
func GetRootDevice() (string, error) {
	fp, err := os.Open("/proc/self/mounts")
	if err != nil {
		return "", err
	}
	defer logger.Client.CloseOrLog(fp)

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
