/*
Package config is responsible for application configuration init.
*/
package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/fs"
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
)

type (
	Config struct {
		BtrfsSupported bool            `json:"btrfs_supported"`
		Path           string          `json:"path"` // needs to save config file from app.
		Schedule       domain.Schedule `json:"schedule"`
		Volumes        []domain.Volume `json:"volumes"`
		RootDevice     string          `json:"root_device"`
		FileMode       os.FileMode     `json:"file_mode"`
	}
)

// New returns config that is stored in default config path.
func New(log *logger.Client) (*Config, error) {
	// delorean path
	configPath, err := getConfigPath(deloreanPath, domain.RWFileMode)
	if err != nil {
		return nil, err
	}

	f, err := os.OpenFile(path.Clean(configPath), os.O_CREATE, domain.RWFileMode)
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
	cfg.FileMode = domain.RWFileMode
	for i := range cfg.Volumes {
		cfg.Volumes[i].Device.Mounted = false
	}

	err = cfg.checkIfKernelSupportsBtrfs()
	if err != nil {
		log.ErrLog.Printf("can't check if btrfs is supported by kernel: %v", err)

		return nil, fmt.Errorf("can't check if btrfs is supported by kernel: %v", err)
	}

	// volumes
	vv, err := findmnt.GetVolumes()
	if err != nil {
		return nil, fmt.Errorf("can't get volumes: %v", err)
	}

	err = cfg.setupVolumes(vv)
	if err != nil {
		return nil, fmt.Errorf("can't setup volumes: %v", err)
	}

	err = mountTopLevelSubvolume(cfg.RootDevice, domain.RWFileMode)
	if err != nil {
		return nil, fmt.Errorf("can't mount top level subvolume: %v", err)
	}

	err = checkDir(path.Join(domain.DeloreanMountPoint, domain.SnapshotsDirName), domain.RWFileMode)
	if err != nil {
		return nil, fmt.Errorf("can't check snapshots directory: %v", err)
	}

	err = cfg.createSnapshotsPaths(domain.RWFileMode)
	if err != nil {
		return nil, err
	}

	err = cfg.Save()
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}

func (cfg *Config) createSnapshotsPaths(fm fs.FileMode) error {
	for i := range cfg.Volumes {
		p := path.Join(domain.DeloreanMountPoint, domain.SnapshotsDirName, cfg.Volumes[i].Subvol)
		if cfg.Volumes[i].Device.Path != cfg.RootDevice {
			p = path.Join(cfg.Volumes[i].Device.MountPoint, domain.SnapshotsDirName, cfg.Volumes[i].Subvol)
		}

		err := createSnapshotsPaths(p, fm)
		if err != nil {
			return fmt.Errorf("can't create snapshots paths: %v", err)
		}

		cfg.Volumes[i].SnapshotsPath = p
	}

	return nil
}

func checkDir(ph string, fm fs.FileMode) error {
	_, err := os.Stat(ph)
	if os.IsNotExist(err) {
		return os.MkdirAll(ph, fm)
	}

	return err
}

// Save flushes current config to file.
func (cfg *Config) Save() error {
	data, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return fmt.Errorf("can't marshal data: %v", err)
	}

	return os.WriteFile(cfg.Path, data, cfg.FileMode)
}

// must be run after getting config from file.
func (cfg *Config) checkIfKernelSupportsBtrfs() error {
	if cfg.BtrfsSupported { // has been checked before and saved in config
		return nil
	}

	var err error
	cfg.BtrfsSupported, err = btrfs.SupportedByKernel()
	if err != nil {
		return fmt.Errorf("can't check if btrfs is supported by kernel: %v", err)
	}

	if !cfg.BtrfsSupported {
		return errors.New("btrfs is not supported by the kernel")
	}

	return nil
}

func (cfg *Config) setupVolumes(vv []domain.Volume) error {
	var err error

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
			return err
		}

		cfg.Volumes = append(cfg.Volumes, vv[i])
	}

	return nil
}

func createSnapshotsPaths(p string, fm fs.FileMode) error {
	for _, v := range []string{
		domain.Manual,
		domain.Monthly,
		domain.Weekly,
		domain.Daily,
		domain.Hourly,
		domain.Boot,
		domain.Revert,
	} {
		err := checkDir(path.Join(p, v), fm)
		if err != nil {
			return fmt.Errorf("can't create snapshot directory for %s: %v", v, err)
		}
	}

	return nil
}

func mountTopLevelSubvolume(device string, fm fs.FileMode) error {
	// snapshots path
	err := checkDir(domain.DeloreanMountPoint, fm)
	if err != nil {
		return fmt.Errorf("can't create default snapshots dir: %v", err)
	}

	if findmnt.IsDeviceMount(device, domain.DeloreanMountPoint) {
		return nil
	}

	return commands.Mount(device, domain.DeloreanMountPoint)
}

func getConfigPath(dir string, fm fs.FileMode) (string, error) {
	configDir := path.Join(dir, "config")
	err := checkDir(configDir, fm)
	if err != nil {
		return "", fmt.Errorf("can't create config directory: %v", err)
	}

	return path.Join(configDir, "config.json"), nil
}
