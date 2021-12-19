package config

import (
	"encoding/json"
	"fmt"
	"io"
	"io/fs"
	"io/ioutil"
	"os"
	"time"

	"github.com/EmmettCorp/delorean/pkg/closer"
)

const (
	configPath    = ".config/delorean"
	appDir        = ".delorean"
	logNameFormat = "2006-01-02_15-04-05"
	// The path where all snapshots will be stored.
	// TODO: put in in `run` directory on release
	storePath = "$HOME/.delorean"
	fileMode  = 0666
)

type (
	Config struct {
		Path      string   `json:"path"` // needs to save config file from app.
		Mouse     bool     `json:"mouse"`
		LogPath   string   `json:"log_path"`
		StorePath string   `json:"store_path"`
		Schedule  Schedule `json:"schedule"`
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
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, fmt.Errorf("can't get user home dir: %v", err)
	}
	configDir := fmt.Sprintf("%s/%s", homeDir, configPath)
	err = checkDir(configDir, os.ModePerm)
	if err != nil {
		return nil, fmt.Errorf("can't create config directory: %v", err)
	}
	configPath := fmt.Sprintf("%s/config.json", configDir)

	f, err := os.OpenFile(configPath, os.O_CREATE, fileMode)
	if err != nil {
		return nil, fmt.Errorf("can't open file: %v", err)
	}
	defer closer.CloseOrLog(f)

	appDir := fmt.Sprintf("%s/%s", homeDir, appDir)
	err = checkDir(appDir, os.ModePerm)
	if err != nil {
		return nil, fmt.Errorf("can't create log directory: %v", err)
	}
	err = checkDir(fmt.Sprintf("%s/logs", appDir), os.ModePerm)
	if err != nil {
		return nil, fmt.Errorf("can't create log directory: %v", err)
	}

	cfg := Config{
		Path: configPath,
	}

	err = json.NewDecoder(f).Decode(&cfg)
	if err != nil && err != io.EOF {
		return nil, fmt.Errorf("can't decode config: %v", err)
	}

	// update it on each init
	cfg.LogPath = fmt.Sprintf("%s/logs/%s.log", appDir, time.Now().Format(logNameFormat))
	cfg.Path = configPath

	if cfg.StorePath == "" {
		cfg.StorePath = storePath
	}

	err = cfg.Save()
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}

func checkDir(path string, mode fs.FileMode) error {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		err = os.Mkdir(path, mode)
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
