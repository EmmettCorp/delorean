package schedule

const scriptsDirectoryName = "scripts"
const systemdDirectoryName = "systemd"

func CreateUnits() error {
	err := createAfterRebootUnits()
	if err != nil {
		return err
	}

	return nil
}
