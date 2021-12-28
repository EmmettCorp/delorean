/*
Package commands keeps all the os commands.
*/
package commands

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/EmmettCorp/delorean/pkg/domain"
)

const (
	deviceIdx      = 0
	pathIdx        = 1
	typeIdx        = 2
	snapshotFormat = "2006-01-02_15-04-05"
)

// CreateSnapshot creates a new snapshot.
func CreateSnapshot(sv, path string) error {
	return exec.Command("btrfs", "subvolume", "snapshot", "-r",
		sv, fmt.Sprintf("%s/%s", path, time.Now().Format(snapshotFormat))).Run()
}

// DeleteSnapshot deletes existing snapshot by path.
func DeleteSnapshot(path string) error {
	return exec.Command("btrfs", "subvolume", "delete", path).Run()
}

// SnapshotsList returns the snapshots list for all active subvolumes with desc sort.
func SnapshotsList(volumes []domain.Volume) ([]domain.Snapshot, error) {
	snaps := []domain.Snapshot{}
	for _, v := range volumes {
		if !v.Active {
			continue
		}
		sn, err := snapshotsListByVolume(v)
		if err != nil {
			return nil, err
		}
		snaps = append(snaps, sn...)
	}

	return snaps, nil
}

func snapshotsListByVolume(volume domain.Volume) ([]domain.Snapshot, error) {
	cmd := exec.Command("btrfs", "subvolume", "list", "-s", "--sort=-gen", volume.Point)
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	snaps := []domain.Snapshot{}

	scanner := bufio.NewScanner(bytes.NewReader(output))
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		if len(fields) == 0 {
			continue
		}
		sn := domain.Snapshot{
			Path:   fmt.Sprintf("%s/%s", volume.Point, fields[len(fields)-1]),
			Volume: volume.Label,
		}
		sn.SetLabel()
		sn.SetType()
		snaps = append(snaps, sn)
	}

	return snaps, nil
}

// GetVolumes returns all the btrfs volumes in current filesystem.
func GetVolumes() ([]domain.Volume, error) {
	volumes := []domain.Volume{}

	fp, err := os.Open("/proc/self/mounts")
	if err != nil {
		return nil, err
	}
	defer fp.Close()

	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		if fields[typeIdx] != "btrfs" {
			continue
		}

		point := fields[pathIdx]

		label := getVolumeLabelByPath(point)
		if label == "" {
			label = point
		}

		volumes = append(volumes, domain.Volume{
			Label:  label,
			Point:  point,
			Device: fields[deviceIdx],
		})
	}

	if scanner.Err() != nil {
		return nil, scanner.Err()
	}

	return volumes, nil
}

func getVolumeLabelByPath(p string) string {
	cmd := exec.Command("btrfs", "filesystem", "label", p)
	output, err := cmd.Output()
	if err != nil {
		return ""
	}

	return strings.TrimSpace(string(output))
}
