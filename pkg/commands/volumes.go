package commands

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"os/exec"
	"strings"

	"github.com/EmmettCorp/delorean/pkg/domain"
)

const (
	btrfsType    = "btrfs"
	lsblkOptions = "PATH,UUID,LABEL,MOUNTPOINT,FSTYPE,HOTPLUG"
)

// order depends on lsblkOptions
const (
	deviceIDx = iota
	uuidIDx
	labelIDx
	mountPointIDx
	fsTypeIDx
	pluggableIDx
)

// GetVolumes returns volumes data from /etc/fstab.
func GetVolumes() ([]domain.Volume, error) {
	// Using `-P`` here guaranties 6 fields.
	// If don't use it and as example is not set 5 fields will be returned.
	cmd := exec.Command("lsblk", "-P", "-o", lsblkOptions)
	var cmdErr bytes.Buffer
	cmd.Stderr = &cmdErr
	output, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("can't execute %s: %s", cmd.String(), cmdErr.String())
	}

	vv := []domain.Volume{}

	scanner := bufio.NewScanner(bytes.NewReader(output))
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		if len(fields) < pluggableIDx-1 {
			continue
		}
		tp := getValueFromAssignmetOrEmpty(fields[fsTypeIDx])
		if tp != btrfsType {
			continue
		}

		v, err := buildVolume(fields)
		if err != nil {
			return nil, fmt.Errorf("can't build volume: %v", err)
		}

		v.Subvol, err = getSubvolume(v.MountPoint)
		if err != nil {
			return nil, fmt.Errorf("can't get subvolume id: %v", err)
		}

		vv = append(vv, v)
	}

	if scanner.Err() != nil {
		return nil, scanner.Err()
	}

	return vv, nil
}

func buildVolume(fields []string) (domain.Volume, error) {
	dev, err := getValueFromAssignmet(fields[deviceIDx])
	if err != nil {
		return domain.Volume{}, fmt.Errorf("can't get device path `%s`: %v", fields[deviceIDx], err)
	}

	uuid, err := getValueFromAssignmet(fields[uuidIDx])
	if err != nil {
		return domain.Volume{}, fmt.Errorf("can't get uuid `%s`: %v", fields[uuidIDx], err)
	}

	mountPoint, err := getValueFromAssignmet(fields[mountPointIDx])
	if err != nil {
		return domain.Volume{}, fmt.Errorf("can't get mount point `%s`: %v", fields[mountPointIDx], err)
	}

	v := domain.Volume{
		Device:     dev,
		UUID:       uuid,
		Label:      getValueFromAssignmetOrEmpty(fields[labelIDx]),
		MountPoint: mountPoint,
		Pluggable:  getValueFromAssignmetOrEmpty(fields[pluggableIDx]) == "1",
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

func getValueFromAssignmetOrEmpty(s string) string {
	v, err := getValueFromAssignmet(s)
	if err != nil {
		return ""
	}

	return v
}
