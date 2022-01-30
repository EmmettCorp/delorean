package config

import (
	"fmt"
	"path"

	"github.com/EmmettCorp/delorean/pkg/domain"
)

// initPaths checks if all required directories are exist and create them if not.
func initPaths() error {
	configDir := path.Join(domain.DeloreanPath, "config")
	err := domain.CheckDir(configDir, domain.RWFileMode)
	if err != nil {
		return fmt.Errorf("can't create config directory: %v", err)
	}

	err = domain.CheckDir(path.Join(domain.DeloreanPath, "scripts"), domain.RWFileMode)
	if err != nil {
		return fmt.Errorf("can't create scripts directory: %v", err)
	}

	err = domain.CheckDir(path.Join(domain.DeloreanPath, "systemd"), domain.RWFileMode)
	if err != nil {
		return fmt.Errorf("can't create systemd directory: %v", err)
	}

	return nil
}
