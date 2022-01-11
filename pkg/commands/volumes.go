package commands

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"os/exec"
	"strings"

	"github.com/EmmettCorp/delorean/pkg/domain"
)

const (
	btrfsType      = "btrfs"
	findmntOptions = "SOURCE,TARGET,FS-OPTIONS,LABEL,UUID"
)

type findMntOutput struct {
	Filesystems []findMntVolume `json:"filesystems"`
}

type findMntVolume struct {
	Source    string `json:"source"`
	Target    string `json:"target"`
	FsOptions string `json:"fs-options"`
	Label     string `json:"label"`
	UUID      string `json:"uuid"`
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

// return value of string with equal sign
// like `UUID=some-uuid-here`
// will return `some-uuid-here`.
func getValueFromAssignmet(s string) (string, error) {
	ss := strings.Split(s, "=")
	if len(ss) != 2 {
		return "", errors.New("there is no equal sign in the string")
	}

	v := strings.Trim(ss[1], `"`)

	return v, nil
}

func buildVolume(fmv findMntVolume) (domain.Volume, error) {
	v := domain.Volume{
		Label:      fmv.Label,
		Device:     strings.Split(fmv.Source, "[")[0],
		MountPoint: fmv.Target,
		UUID:       fmv.UUID,
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

	var err error
	v.Subvol, err = getSubvolFromOptions(fmv.FsOptions)
	if err != nil {
		return domain.Volume{}, err
	}

	return v, nil
}

func getSubvolFromOptions(options string) (string, error) {
	opts := strings.Split(options, ",")
	if len(opts) == 0 {
		return "", errors.New("empty options")
	}

	for i := range opts {
		if strings.HasPrefix(opts[i], "subvol=") {
			sbvl, err := getValueFromAssignmet(opts[i])
			if err != nil {
				return "", err
			}

			if sbvl == "" {
				return "", errors.New("subvol is empty")
			}

			if sbvl == "/" {
				return "root", nil
			}

			if strings.HasPrefix(sbvl, "/") {
				ss := strings.Split(sbvl, "/")
				if len(ss) != 2 {
					return "", fmt.Errorf("subvol is invalid: %s", sbvl)
				}
				sbvl = ss[1]
			}

			return sbvl, nil
		}
	}

	return "", errors.New("there is no subvol in options")
}
