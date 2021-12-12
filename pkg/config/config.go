package config

import (
	"encoding/json"
	"fmt"
	"io"
	"io/fs"
	"os"
	"time"
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

type Config struct {
	LogPath   string
	StorePath string
}

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

	f, err := os.OpenFile(fmt.Sprintf("%s/config.json", configDir), os.O_CREATE, 0666)
	if err != nil {
		return nil, fmt.Errorf("can't open file: %v", err)
	}

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
		LogPath: fmt.Sprintf("%s/logs/%s.log", appDir, time.Now().Format(logNameFormat)),
	}

	err = json.NewDecoder(f).Decode(&cfg)
	if err != nil && err != io.EOF {
		return nil, fmt.Errorf("can't decode config: %v", err)
	}

	if cfg.StorePath == "" {
		cfg.StorePath = storePath
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
