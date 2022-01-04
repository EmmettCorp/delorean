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
	"path"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/EmmettCorp/delorean/pkg/domain"
)

const (
	deviceIdx      = 0
	pathIdx        = 1
	typeIdx        = 2
	snapID         = 1
	snapshotFormat = "2006-01-02_15:04:05"
)

type sortableSnapshots []domain.Snapshot

func (ss sortableSnapshots) Len() int           { return len(ss) }
func (ss sortableSnapshots) Swap(i, j int)      { ss[i], ss[j] = ss[j], ss[i] }
func (ss sortableSnapshots) Less(i, j int) bool { return ss[i].ID > ss[j].ID }

// CreateSnapshot creates a new snapshot.
func CreateSnapshot(sv, ph string) error {
	cmd := exec.Command("btrfs", "subvolume", "snapshot", sv, path.Join(ph, time.Now().Format(snapshotFormat)))
	var cmdErr bytes.Buffer
	cmd.Stderr = &cmdErr
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("can't execute %s: %s", cmd.String(), cmdErr.String())
	}

	return nil
}

// SetDefault sets subvolume as default by id.
func SetDefault(volumePath string, snapID int64) error {
	cmd := exec.Command("btrfs", "subvolume", "set-default", fmt.Sprintf("%d", snapID), volumePath)
	err := cmd.Run()
	var cmdErr bytes.Buffer
	cmd.Stderr = &cmdErr
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

// SnapshotsList returns the snapshots list for all active subvolumes with desc sort.
func SnapshotsList(volumes []domain.Volume) ([]domain.Snapshot, error) {
	snaps := sortableSnapshots{}
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

	sort.Sort(snaps)

	return snaps, nil
}

func snapshotsListByVolume(volume domain.Volume) ([]domain.Snapshot, error) {
	cmd := exec.Command("btrfs", "subvolume", "list", "-s", volume.Point)
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

		id, err := strconv.ParseInt(fields[snapID], 10, 64)
		if err != nil {
			return nil, err
		}

		sn := domain.Snapshot{
			ID:          id,
			Path:        path.Join(volume.Point, fields[len(fields)-1]),
			VolumeLabel: volume.Label,
			VolumePoint: volume.Point,
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
