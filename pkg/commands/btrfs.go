/*
Package commands keeps all the os commands.
*/
package commands

import (
	"bufio"
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

const (
	labelIdx      = 0
	LabelValueIdx = 1
	uidIdx        = 3
)

// Volume represents btrfs volume.
type Volume struct {
	Label string
	UID   string
}

// CreateSnapshot creates a new snapshot.
func CreateSnapshot(sv, path string) error {
	return exec.Command("btrfs", "subvolume", "snapshot", sv, path).Run()
}

// DeleteSnapshot deletes existing snapshot by path.
func DeleteSnapshot(path string) error {
	return exec.Command("btrfs", "subvolume", "delete", path).Run()
}

// GetVolumes returns all the btrfs volumes in current filesystem.
func GetVolumes() ([]Volume, error) {
	cmd := exec.Command("btrfs", "filesystem", "show")
	output, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("can't get output: %v", err)
	}

	scanner := bufio.NewScanner(bytes.NewReader(output))

	volumes := []Volume{}

	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		if len(fields) < uidIdx+1 {
			continue
		}
		if fields[labelIdx] != "Label:" {
			continue
		}

		label := strings.Trim(fields[LabelValueIdx], "'") // label value prints with quotes. like 'label'

		volumes = append(volumes, Volume{
			Label: label,
			UID:   fields[uidIdx],
		})
	}

	return volumes, nil
}
