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

// GetSubvolumeID deletes existing snapshot by path.
func getSubvolumeID(ph string) (string, error) {
	cmd := exec.Command("btrfs", "subvolume", "show", ph)
	var cmdErr bytes.Buffer
	cmd.Stderr = &cmdErr
	output, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("can't get subvolume id %s: %s", cmd.String(), cmdErr.String())
	}

	name := "Name:"
	var value string

	scanner := bufio.NewScanner(bytes.NewReader(output))
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		if len(fields) < 2 {
			continue
		}

		if fields[0] != name {
			continue
		}

		value = fields[1]
	}

	if scanner.Err() != nil {
		return "", scanner.Err()
	}

	return value, nil
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
	cmd := exec.Command("btrfs", "subvolume", "list", "-s", volume.MountPoint)
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
			Path:        path.Join(volume.MountPoint, fields[len(fields)-1]),
			VolumeLabel: volume.Label,
			VolumePoint: volume.MountPoint,
		}
		sn.SetLabel()
		sn.SetType()
		snaps = append(snaps, sn)
	}

	return snaps, nil
}

// BtrfsSupported checks if kernel supports btrfs.
func BtrfsSupported() (bool, error) {
	var wordIDx int
	fp, err := os.Open("/proc/filesystems")
	if err != nil {
		return false, err
	}
	defer fp.Close()

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
