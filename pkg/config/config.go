package config

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"time"

	"github.com/EmmettCorp/delorean/pkg/closer"
	"github.com/EmmettCorp/delorean/pkg/commands"
	"github.com/EmmettCorp/delorean/pkg/domain"
)

const (
	deloreanPath        = "/opt/delorean"
	defaultLogDir       = "/var/log/delorean"
	logNameFormat       = "2006-01-02_15-04-05"
	defaultSnapshotsDir = ".snapshots"
	fileMode            = 0600
)

type (
	Config struct {
		Path       string          `json:"path"` // needs to save config file from app.
		LogPath    string          `json:"log_path"`
		RootDevice string          `json:"root_device"`
		Schedule   Schedule        `json:"schedule"`
		Volumes    []domain.Volume `json:"volumes"`
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
func New() (*Config, error) {
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
	defer closer.CloseOrLog(f)
	var cfg Config
	err = json.NewDecoder(f).Decode(&cfg)
	if err != nil && err != io.EOF {
		return nil, fmt.Errorf("can't decode config: %v", err)
	}
	cfg.Path = configPath

	// create a new log file
	err = checkDir(defaultLogDir)
	if err != nil {
		return nil, fmt.Errorf("can't create log directory: %v", err)
	}
	cfg.LogPath = fmt.Sprintf("%s/%s.log", defaultLogDir, time.Now().Format(logNameFormat))

	// get root device
	cfg.RootDevice, err = commands.GetRootDevice()
	if err != nil {
		return nil, fmt.Errorf("can't get root device: %v", err)
	}

	rootSnapshotsPath := fmt.Sprintf("/%s", defaultSnapshotsDir)
	err = createSnapshotsPaths(rootSnapshotsPath)
	if err != nil {
		return nil, fmt.Errorf("can't create snapshots paths: %v", err)
	}

	// volumes
	vv, err := commands.GetVolumes()
	if err != nil {
		return nil, fmt.Errorf("can't get volumes: %v", err)
	}

OUT:
	for i := range vv {
		for j := range cfg.Volumes {
			if vv[i].Point == cfg.Volumes[j].Point { // check if this path has been already added
				continue OUT
			}
		}

		if vv[i].Device != cfg.RootDevice {
			snapshotsPath := fmt.Sprintf("%s/%s", vv[i].Point, defaultSnapshotsDir)
			err = createSnapshotsPaths(fmt.Sprintf("%s/%s", vv[i].Point, defaultSnapshotsDir))
			if err != nil {
				return nil, fmt.Errorf("can't create snapshots paths: %v", err)
			}
			vv[i].SnapshotsPath = snapshotsPath
		} else {
			vv[i].SnapshotsPath = rootSnapshotsPath
		}

		cfg.Volumes = append(cfg.Volumes, vv[i])
	}

	err = cfg.Save()
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}

func checkDir(path string) error {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		err = os.Mkdir(path, fileMode)
		if err != nil {
			return err
		}
	}

	return nil
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
	err := checkDir(p)
	if err != nil {
		return fmt.Errorf("can't create snapshot directory: %v", err)
	}

	for _, v := range []string{domain.Manual, domain.Monthly, domain.Weekly, domain.Daily, domain.Hourly, domain.Boot} {
		err := checkDir(fmt.Sprintf("%s/%s", p, v))
		if err != nil {
			return fmt.Errorf("can't create snapshot directory for %s: %v", v, err)
		}
	}

	return nil
}
