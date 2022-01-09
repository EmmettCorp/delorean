package commands

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/EmmettCorp/delorean/pkg/domain"
)

const (
	deviceIDx = iota
	pathIDx
	fsTypeIDx
	optionsIDx
)

const btrfsType = "btrfs"

// GetVolumes returns volumes data from /etc/fstab.
func GetVolumes() ([]domain.Volume, error) {
	fp, err := os.Open("/etc/fstab")
	if err != nil {
		return nil, err
	}
	defer fp.Close()

	vv := []domain.Volume{}

	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		if len(fields) < optionsIDx {
			continue
		}
		if fields[fsTypeIDx] != btrfsType {
			continue
		}

		device := fields[deviceIDx]
		if strings.HasPrefix(device, "UUID") {
			d, err := getValueFromAssignmet(device)
			if err != nil {
				return nil, fmt.Errorf("can't get uuid %s: %v", device, err)
			}

			device = d
		}

		subID, err := getSubvolumeID(fields[optionsIDx])
		if err != nil {
			return nil, err
		}

		v := domain.Volume{
			ID:     subID,
			Point:  fields[pathIDx],
			Device: device,
		}

		switch v.Point {
		case "/":
			v.Label = "Root"
		case "/home":
			v.Label = "Home"
		default:
			v.Label = v.Point
		}

		vv = append(vv, v)

	}

	if scanner.Err() != nil {
		return nil, scanner.Err()
	}

	return vv, nil
}

// return value of string with equal sign
// like `UUID=some-uuid-here`
// will return `some-uuid-here`
func getValueFromAssignmet(s string) (string, error) {
	ss := strings.Split(s, "=")
	if len(ss) != 2 {
		return "", errors.New("there is no equal sign in the string")
	}

	return ss[1], nil
}

func getSubvolumeID(options string) (string, error) {
	opts := strings.Split(options, ",")
	if len(opts) == 0 {
		return "", errors.New("empty options")
	}

	for i := range opts {
		if strings.HasPrefix(opts[i], "subvol") {
			id, err := getValueFromAssignmet(opts[i])
			if err != nil {
				return "", err
			}

			if id == "" {
				return "", errors.New("subvolume id is empty")
			}

			if id == "/" {
				return "root", nil
			}

			if strings.HasPrefix(id, "/") {
				ss := strings.Split(id, "/")
				if len(ss) != 2 {
					return "", fmt.Errorf("id is invalid: %s", id)
				}
				id = ss[1]
			}

			return id, nil
		}
	}

	return "", errors.New("there is no subvolume id in options")
}
