/*
Package btrfs keeps all the btrfs commands.
*/
package btrfs

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path"
	"sort"
	"strings"
	"time"

	"github.com/EmmettCorp/delorean/pkg/domain"
)

type sortableSnapshots []domain.Snapshot

func (ss sortableSnapshots) Len() int           { return len(ss) }
func (ss sortableSnapshots) Swap(i, j int)      { ss[i], ss[j] = ss[j], ss[i] }
func (ss sortableSnapshots) Less(i, j int) bool { return ss[i].Timestamp > ss[j].Timestamp }

// CreateSnapshot creates a new snapshot.
func CreateSnapshot(sv, ph string) error {
	cmd := exec.Command("btrfs", "subvolume", "snapshot", "-r", sv, path.Join(ph, time.Now().Format(domain.SnapshotFormat)))
	var cmdErr bytes.Buffer
	cmd.Stderr = &cmdErr
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("can't execute %s: %s", cmd.String(), cmdErr.String())
	}

	return nil
}

// Restore creates a new snapshot.
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

// SnapshotsList returns the snapshots list for all active subvolumes with desc sort.
func SnapshotsList(volumes []domain.Volume) ([]domain.Snapshot, error) {
	snaps := sortableSnapshots{}
	for _, v := range volumes {
		if !v.Active || !v.Device.Mounted {
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
	snaps := []domain.Snapshot{}
	sn, err := getSnapshots(volume.SnapshotsPath)
	if err != nil {
		return nil, err
	}

	for i := range sn {
		sn := domain.Snapshot{
			Path:        sn[i],
			VolumeLabel: volume.Label,
			VolumeID:    volume.ID,
		}
		sn.SetLabel()
		sn.SetType()
		sn.SetTimestamp()
		snaps = append(snaps, sn)
	}

	return snaps, nil
}

// SupportedByKernel checks if kernel supports btrfs.
func SupportedByKernel() (bool, error) {
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

func GetVolumeID(ph string) (string, error) {
	cmd := exec.Command("btrfs", "subvolume", "show", ph)
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}

	scanner := bufio.NewScanner(bytes.NewReader(output))
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		if len(fields) < 3 {
			continue
		}

		if fields[0] == "Subvolume" && fields[1] == "ID:" {
			return fields[len(fields)-1], nil
		}
	}

	return "", fmt.Errorf("can't find volume id from path %s", ph)
}

func getSnapshots(ph string) ([]string, error) {
	dirs, err := osReadDir(ph)
	if err != nil {
		return nil, err
	}
	snaps := []string{}
	for i := range dirs {
		sp := path.Join(ph, dirs[i])
		snap, err := osReadDir(sp)
		if err != nil {
			return nil, err
		}

		for j := range snap {
			snaps = append(snaps, path.Join(sp, snap[j]))
		}

	}

	return snaps, nil
}

func osReadDir(root string) ([]string, error) {
	var files []string
	f, err := os.Open(root)
	if err != nil {
		return files, err
	}
	fileInfo, err := f.Readdir(-1)
	f.Close()
	if err != nil {
		return files, err
	}

	for _, file := range fileInfo {
		files = append(files, file.Name())
	}
	return files, nil
}
