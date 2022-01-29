package schedule

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path"

	"github.com/EmmettCorp/delorean/pkg/domain"
)

const (
	afterRestoreShFile      = "after-restore.sh"
	afterRestoreServiceFile = "after-restore.service"
	afterRestoreTimerFile   = "after-restore.timer"
)

const afterRestoreTemplate = `#!/bin/bash

if [ ! -d %s ]; then mkdir %s && mount %s %s; fi
rm -rf %s
`

// RemoveOldRootAfterReboot deletes old root directory after the next reboot.
// Call it on restore.
func RemoveOldRootAfterReboot(rootDevice, ph string) error {
	unitFilePath := path.Join(domain.DeloreanPath, scriptsDirectoryName, afterRestoreShFile)
	err := os.WriteFile(unitFilePath,
		[]byte(fmt.Sprintf(afterRestoreTemplate,
			domain.DeloreanMountPoint, domain.DeloreanMountPoint, rootDevice, domain.DeloreanMountPoint, ph),
		),
		0o100,
	)
	if err != nil {
		return err
	}

	cmd := exec.Command("systemctl", "enable", path.Join(
		domain.DeloreanPath, systemdDirectoryName, afterRestoreTimerFile))
	var cmdErr bytes.Buffer
	cmd.Stderr = &cmdErr
	err = cmd.Run()
	if err != nil {
		return fmt.Errorf("can't execute %s: %s", cmd.String(), cmdErr.String())
	}

	return nil
}

func createAfterRebootUnits() error {
	pathToService := path.Join(domain.DeloreanPath, systemdDirectoryName, afterRestoreServiceFile)
	pathToTimer := path.Join(domain.DeloreanPath, systemdDirectoryName, afterRestoreTimerFile)
	_, errSrv := os.Stat(pathToService)
	_, errTmr := os.Stat(pathToTimer)
	if errSrv == nil && errTmr == nil {
		return nil
	}

	pathToSh := path.Join(domain.DeloreanPath, scriptsDirectoryName, afterRestoreShFile)

	err := os.WriteFile(pathToService, []byte(fmt.Sprintf(`[Unit]
Description=Delorean scripts running on boot after restore.

[Service]
ExecStart=/bin/bash %s
ExecStartPost=/bin/echo > %s
Restart=no
Type=oneshot

[Install]
WantedBy=multi-user.target
`, pathToSh, pathToSh)), domain.RWFileMode)
	if err != nil {
		return err
	}

	err = os.WriteFile(path.Join(domain.DeloreanPath, systemdDirectoryName, afterRestoreTimerFile),
		[]byte(`[Unit]
Description=Delorean scripts running on boot after restore.

[Timer]
OnUnitActiveSec=10s
OnBootSec=10s

[Install]
WantedBy=timers.target
`), domain.RWFileMode)
	if err != nil {
		return err
	}

	return nil
}
