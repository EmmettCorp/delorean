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

type (
	// Config represents configuration of application.
	// It keeps all needed settings. Config is saved on a disk.
	Config struct {
		RootDevice     string          `json:"root_device"`
		Path           string          `json:"path"` // needs to save config file from app.
		ToRemove       []string        `json:"to_remove"`
		BtrfsSupported bool            `json:"btrfs_supported"`
		Schedule       domain.Schedule `json:"schedule"`
		Volumes        []domain.Volume `json:"volumes"`
		FileMode       os.FileMode     `json:"file_mode"`
	}
)

// New returns config that is stored in default config path.
func New() (*Config, error) {
	err := initPaths()
	if err != nil {
		return nil, fmt.Errorf("can't init paths: %v", err)
	}

	configPath := path.Clean(path.Join(domain.DeloreanPath, "config", "config.json"))
	f, err := os.OpenFile(configPath, os.O_CREATE, domain.RWFileMode) // nolint:gosec // path constructed from consts
	if err != nil {
		return nil, fmt.Errorf("can't open file: %v", err)
	}
	defer logger.Client.CloseOrLog(f)

	var cfg Config
	err = json.NewDecoder(f).Decode(&cfg)
	if err != nil && err != io.EOF {
		return nil, fmt.Errorf("can't decode config: %v", err)
	}
	cfg.Path = configPath
	cfg.FileMode = domain.RWFileMode

	err = cfg.checkIfKernelSupportsBtrfs()
	if err != nil {
		logger.Client.ErrLog.Printf("can't check if btrfs is supported by kernel: %v", err)

		return nil, fmt.Errorf("can't check if btrfs is supported by kernel: %v", err)
	}

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

	err = cfg.createSnapshotsPaths(domain.RWFileMode)
	if err != nil {
		return nil, err
	}

	cfg.removeOldSubvolumes()

	err = cfg.Save()
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}

func (cfg *Config) createSnapshotsPaths(fm fs.FileMode) error {
	err := domain.CheckDir(path.Join(domain.DeloreanMountPoint, domain.SnapshotsDirName), fm)
	if err != nil {
		return fmt.Errorf("can't check snapshots directory: %v", err)
	}

	for i := range cfg.Volumes {
		p := path.Join(domain.DeloreanMountPoint, domain.SnapshotsDirName, cfg.Volumes[i].Subvol)
		if cfg.Volumes[i].Device.Path != cfg.RootDevice {
			p = path.Join(cfg.Volumes[i].Device.MountPoint, domain.SnapshotsDirName, cfg.Volumes[i].Subvol)
		}

		err = createSnapshotsPaths(p, fm)
		if err != nil {
			return fmt.Errorf("can't create snapshots paths: %v", err)
		}

		cfg.Volumes[i].SnapshotsPath = p
	}

	return nil
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

	for i := range vv {
		vv[i].ID, err = btrfs.GetVolumeID(vv[i].Device.MountPoint)
		if err != nil {
			return err
		}

		for j := range cfg.Volumes {
			if vv[i].ID == cfg.Volumes[j].ID {
				vv[i].Active = cfg.Volumes[j].Active
			}
		}

		if vv[i].Device.MountPoint == "/" {
			cfg.RootDevice = vv[i].Device.Path
		}
	}

	cfg.Volumes = vv

	return nil
}

func (cfg *Config) removeOldSubvolumes() {
	for _, v := range cfg.ToRemove {
		go removeOld(v)
	}

	cfg.ToRemove = []string{}
}

func removeOld(dir string) {
	err := os.RemoveAll(dir)
	if err != nil {
		logger.Client.ErrLog.Printf("can't remove old subvolume `%s`: %v", dir, err)
	}
}

func createSnapshotsPaths(p string, fm fs.FileMode) error {
	for _, v := range []string{
		domain.Manual,
		domain.Monthly,
		domain.Weekly,
		domain.Daily,
		domain.Hourly,
		domain.Boot,
		domain.Restore,
	} {
		err := domain.CheckDir(path.Join(p, v), fm)
		if err != nil {
			return fmt.Errorf("can't create snapshot directory for %s: %v", v, err)
		}
	}

	return nil
}

func mountTopLevelSubvolume(device string, fm fs.FileMode) error {
	// snapshots path
	err := domain.CheckDir(domain.DeloreanMountPoint, fm)
	if err != nil {
		return fmt.Errorf("can't create default snapshots dir: %v", err)
	}

	if findmnt.IsDeviceMount(device, domain.DeloreanMountPoint) {
		return nil
	}

	return commands.Mount(device, domain.DeloreanMountPoint)
}

// VolumeInRootFs checks if the volume is in root.
func (cfg *Config) VolumeInRootFs(vol domain.Volume) bool {
	return vol.Device.Path == cfg.RootDevice
}
