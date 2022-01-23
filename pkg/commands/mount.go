package commands

import (
	"bytes"
	"fmt"
	"os/exec"
)

// Mount mounts device to point.
func Mount(device, point string) error {
	cmd := exec.Command("mount", device, point)
	var cmdErr bytes.Buffer
	cmd.Stderr = &cmdErr
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("can't execute %s: %s", cmd.String(), cmdErr.String())
	}

	return nil
}

// Umount unmounts device.
func Umount(point string) error {
	cmd := exec.Command("umount", point)
	var cmdErr bytes.Buffer
	cmd.Stderr = &cmdErr
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("can't execute %s: %s", cmd.String(), cmdErr.String())
	}

	return nil
}
