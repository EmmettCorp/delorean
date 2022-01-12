package commands

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os/exec"
	"strings"

	"github.com/EmmettCorp/delorean/pkg/domain"
)

const (
	btrfsType      = "btrfs"
	findmntOptions = "SOURCE,TARGET,LABEL,UUID,FSROOT"
)

type findMntOutput struct {
	Filesystems []findMntVolume `json:"filesystems"`
}

type findMntVolume struct {
	Source string `json:"source"`
	Target string `json:"target"`
	Label  string `json:"label"`
	UUID   string `json:"uuid"`
	FsRoot string `json:"fsroot"`
}

// GetVolumes returns volumes using findmnt.
// It skips mounted by delorean volumes.
func GetVolumes() ([]domain.Volume, error) {
	// -t btrfs: return only btrfs type
	// -J: return in json
	// -l: return as a flat list
	// -o: options {findmntOptions}
	cmd := exec.Command("findmnt", "-t", "btrfs", "-J", "-l", "-o", findmntOptions)
	var cmdErr bytes.Buffer
	cmd.Stderr = &cmdErr
	output, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("can't execute %s: %s", cmd.String(), cmdErr.String())
	}

	fmo := findMntOutput{}

	err = json.Unmarshal([]byte(output), &fmo)
	if err != nil {
		return nil, fmt.Errorf("can't unmarshal data: %v", err)
	}

	vv := []domain.Volume{}
	for i := range fmo.Filesystems {
		if strings.HasPrefix(fmo.Filesystems[i].Target, domain.DeloreanMountPoint) {
			continue
		}

		v, err := buildVolume(fmo.Filesystems[i])
		if err != nil {
			return nil, fmt.Errorf("can't build volume: %v", err)
		}
		vv = append(vv, v)
	}

	return vv, nil
}

func buildVolume(fmv findMntVolume) (domain.Volume, error) {
	v := domain.Volume{
		Label:      fmv.Label,
		Device:     strings.Split(fmv.Source, "[")[0], // safe even if string without `[` character
		MountPoint: fmv.Target,
		UUID:       fmv.UUID,
		Mounted:    true,
		Subvol:     getSubvol(fmv.FsRoot),
	}

	if v.Label == "" {
		switch v.MountPoint {
		case "/":
			v.Label = "Root"
		case "/home":
			v.Label = "Home"
		default:
			v.Label = v.Device
		}
	}

	return v, nil
}

func getSubvol(fsRoot string) string {
	if fsRoot == "/" {
		return domain.Subvol5
	}

	if strings.HasPrefix(fsRoot, "/") { // remove slash from fsRoot like `/@`
		ss := strings.Split(fsRoot, "/")
		return ss[1]
	}

	return fsRoot
}
