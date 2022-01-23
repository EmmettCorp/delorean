package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"

	"github.com/EmmettCorp/delorean/pkg/commands"
	"github.com/EmmettCorp/delorean/pkg/commands/btrfs"
	"github.com/EmmettCorp/delorean/pkg/commands/findmnt"
	"github.com/EmmettCorp/delorean/pkg/domain"
	"github.com/EmmettCorp/delorean/pkg/logger"
)

const (
	deloreanPath = "/usr/local/delorean"
	fileMode     = 0600
)

type (
	Config struct {
		BtrfsSupported bool            `json:"btrfs_supported"`
		Path           string          `json:"path"` // needs to save config file from app.
		Schedule       Schedule        `json:"schedule"`
		Volumes        []domain.Volume `json:"volumes"`
		RootDevice     string          `json:"root_device"`
	}

	Schedule struct {
		Monthly int `json:"monthly"`
		Weekly  int `json:"weekly"`
		Daily   int `json:"daily"`
		Hourly  int `json:"hourly"`
		Boot    int `json:"boot"`
	}
)

// New returns config that is stored in default config path.
func New(log *logger.Client) (*Config, error) {
	// delorean path
	err := checkDir(deloreanPath)
	if err != nil {
		return nil, fmt.Errorf("can't get delorean dir: %v", err)
	}

	// get config
	configDir := fmt.Sprintf("%s/config", deloreanPath)
	err = checkDir(configDir)
	if err != nil {
		return nil, fmt.Errorf("can't create config directory: %v", err)
	}
	configPath := fmt.Sprintf("%s/config.json", configDir)
	f, err := os.OpenFile(configPath, os.O_CREATE, fileMode)
	if err != nil {
		return nil, fmt.Errorf("can't open file: %v", err)
	}
	defer log.CloseOrLog(f)

	var cfg Config
	err = json.NewDecoder(f).Decode(&cfg)
	if err != nil && err != io.EOF {
		return nil, fmt.Errorf("can't decode config: %v", err)
	}
	cfg.Path = configPath
	for i := range cfg.Volumes {
		cfg.Volumes[i].Device.Mounted = false
	}

	if !cfg.BtrfsSupported { // check on first run only
		cfg.BtrfsSupported, err = btrfs.SupportedByKernel()
		if err != nil {
			log.Errorf("can't check if btrfs is supported by kernel: %v", err)
			return nil, fmt.Errorf("can't check if btrfs is supported by kernel: %v", err)
		}

		if !cfg.BtrfsSupported {
			log.Error("btrfs is not supported by the kernel")
			return nil, errors.New("btrfs is not supported by the kernel")
		}
	}

	// volumes
	vv, err := findmnt.GetVolumes()
	if err != nil {
		return nil, fmt.Errorf("can't get volumes: %v", err)
	}

OUT:
	for i := range vv {
		for j := range cfg.Volumes {
			if vv[i].Device.MountPoint == cfg.Volumes[j].Device.MountPoint { // check if this path has been already added
				cfg.Volumes[j].Device.Mounted = true
				continue OUT
			}
		}

		if vv[i].Device.MountPoint == "/" {
			cfg.RootDevice = vv[i].Device.Path
		}

		vv[i].ID, err = btrfs.GetVolumeID(vv[i].Device.MountPoint)
		if err != nil {
			return nil, err
		}

		cfg.Volumes = append(cfg.Volumes, vv[i])
	}

	err = mountTopLevelSubvolume(cfg.RootDevice)
	if err != nil {
		return nil, fmt.Errorf("can't mount top level subvolume: %v", err)
	}

	err = checkDir(path.Join(domain.DeloreanMountPoint, domain.SnapshotsDirName))
	if err != nil {
		return nil, fmt.Errorf("can't check snapshots directory: %v", err)
	}

	err = cfg.createSnapshotsPaths()
	if err != nil {
		return nil, err
	}

	err = cfg.Save()
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}

func (cfg *Config) createSnapshotsPaths() error {
	for i := range cfg.Volumes {
		p := path.Join(domain.DeloreanMountPoint, domain.SnapshotsDirName, cfg.Volumes[i].Subvol)
		if cfg.Volumes[i].Device.Path != cfg.RootDevice {
			p = path.Join(cfg.Volumes[i].Device.MountPoint, domain.SnapshotsDirName, cfg.Volumes[i].Subvol)
		}

		err := createSnapshotsPaths(p)
		if err != nil {
			return fmt.Errorf("can't create snapshots paths: %v", err)
		}

		cfg.Volumes[i].SnapshotsPath = p
	}

	return nil
}

func checkDir(path string) error {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return os.MkdirAll(path, fileMode)
	}

	return err
}

// Save flushes current config to file.
func (cfg *Config) Save() error {
	data, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return fmt.Errorf("can't marshal data: %v", err)
	}

	return ioutil.WriteFile(cfg.Path, data, fileMode)
}

func createSnapshotsPaths(p string) error {
	for _, v := range []string{
		domain.Manual,
		domain.Monthly,
		domain.Weekly,
		domain.Daily,
		domain.Hourly,
		domain.Boot,
		domain.Revert,
	} {
		err := checkDir(path.Join(p, v))
		if err != nil {
			return fmt.Errorf("can't create snapshot directory for %s: %v", v, err)
		}
	}

	return nil
}

func mountTopLevelSubvolume(device string) error {
	// snapshots path
	err := checkDir(domain.DeloreanMountPoint)
	if err != nil {
		return fmt.Errorf("can't create default snapshots dir: %v", err)
	}

	if findmnt.IsDeviceMount(device, domain.DeloreanMountPoint) {
		return nil
	}

	return commands.Mount(device, domain.DeloreanMountPoint)
}
