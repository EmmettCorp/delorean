/*
Package btrfs keeps all needed btrfs commands.
*/
package btrfs

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/EmmettCorp/delorean/pkg/domain"
	"github.com/EmmettCorp/delorean/pkg/logger"
)

// CreateSnapshot creates a new snapshot.
func CreateSnapshot(subvolume string, snap domain.Snapshot) error {
	// nolint:gosec // we pass commands here from code only.
	cmd := exec.Command("btrfs", "subvolume", "snapshot", "-r", subvolume, snap.Path)
	var cmdErr bytes.Buffer
	cmd.Stderr = &cmdErr
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("can't execute %s: %s", cmd.String(), cmdErr.String())
	}

	return nil
}

// Restore creates a new snapshot.
// The idea is not to just move snapshot data to subvolume (like @ or @home) but make a snapshot for it.
func Restore(snapPath, mountPoint string) error {
	cmd := exec.Command("btrfs", "subvolume", "snapshot", snapPath, mountPoint)
	var cmdErr bytes.Buffer
	cmd.Stderr = &cmdErr
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("can't execute %s: %s", cmd.String(), cmdErr.String())
	}

	return nil
}

// DeleteSnapshot deletes existing snapshot by path.
func DeleteSnapshot(ph string) error {
	cmd := exec.Command("btrfs", "subvolume", "delete", ph)
	var cmdErr bytes.Buffer
	cmd.Stderr = &cmdErr
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("can't execute %s: %s", cmd.String(), cmdErr.String())
	}

	return nil
}

// SupportedByKernel checks if kernel supports btrfs.
func SupportedByKernel() (bool, error) {
	var wordIDx int
	fp, err := os.Open("/proc/filesystems")
	if err != nil {
		return false, err
	}
	defer logger.Client.CloseOrLog(fp)

	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		if fields[wordIDx] == "btrfs" {
			return true, nil
		}
	}

	if scanner.Err() != nil {
		return false, scanner.Err()
	}

	return false, nil
}

// GetVolumeID returns volume id by path.
func GetVolumeID(ph string) (string, error) {
	minSubvolIDFields := 3
	cmd := exec.Command("btrfs", "subvolume", "show", ph)
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}

	scanner := bufio.NewScanner(bytes.NewReader(output))
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		if len(fields) < minSubvolIDFields {
			continue
		}

		if fields[0] == "Subvolume" && fields[1] == "ID:" {
			return fields[len(fields)-1], nil
		}
	}

	return "", fmt.Errorf("can't find volume id from path %s", ph)
}
