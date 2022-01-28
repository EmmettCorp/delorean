package config

import (
	"encoding/json"
	"os"
	"path"
	"testing"

	"github.com/EmmettCorp/delorean/pkg/domain"
	"github.com/stretchr/testify/require"
)

func TestCheckDir(t *testing.T) {
	t.Parallel()
	rq := require.New(t)

	t.Run("ok", func(t *testing.T) {
		t.Parallel()

		ph := "./test_check_dir"
		_, err := os.Stat(ph)
		rq.Error(err)

		err = checkDir(ph, 0o006)
		rq.NoError(err)

		_, err = os.Stat(ph)
		rq.NoError(err)

		err = os.Remove(ph)
		rq.NoError(err)
	})
}

func TestSave(t *testing.T) {
	t.Parallel()
	rq := require.New(t)

	t.Run("ok", func(t *testing.T) {
		t.Parallel()
		dir := "test_save"
		ph, err := getConfigPath(dir, 0o777)
		rq.NoError(err)

		cfg := Config{
			Path:     ph,
			FileMode: 0o777,
		}

		err = cfg.Save()
		rq.NoError(err)

		f, err := os.OpenFile(ph, os.O_CREATE, domain.RWFileMode) // nolint:gosec // only from test
		rq.NoError(err)

		var fileCfg Config
		err = json.NewDecoder(f).Decode(&fileCfg)
		rq.NoError(err)
		err = f.Close()
		rq.NoError(err)
		rq.Equal(cfg.Path, fileCfg.Path)
		rq.Equal(cfg.FileMode, fileCfg.FileMode)

		err = os.RemoveAll(dir)
		rq.NoError(err)
	})
}

func TestCreateSnapshotsPaths(t *testing.T) {
	t.Parallel()
	rq := require.New(t)

	t.Run("ok", func(t *testing.T) {
		t.Parallel()

		ph := "test_create_snapshots_paths"

		err := createSnapshotsPaths(ph, 0o777)
		rq.NoError(err)

		_, err = os.Stat(path.Join(ph, domain.Manual))
		rq.NoError(err)

		_, err = os.Stat(path.Join(ph, domain.Monthly))
		rq.NoError(err)

		_, err = os.Stat(path.Join(ph, domain.Weekly))
		rq.NoError(err)

		_, err = os.Stat(path.Join(ph, domain.Daily))
		rq.NoError(err)

		_, err = os.Stat(path.Join(ph, domain.Hourly))
		rq.NoError(err)

		_, err = os.Stat(path.Join(ph, domain.Boot))
		rq.NoError(err)

		_, err = os.Stat(path.Join(ph, domain.Restore))
		rq.NoError(err)

		err = os.RemoveAll(ph)
		rq.NoError(err)
	})
}

func TestGetConfigPath(t *testing.T) {
	t.Parallel()
	rq := require.New(t)

	t.Run("ok", func(t *testing.T) {
		t.Parallel()
		dir := "test_get_config_path"
		ph, err := getConfigPath(dir, 0o777)
		rq.NoError(err)
		rq.Equal(path.Join(dir, "config", "config.json"), ph)

		err = os.RemoveAll(dir)
		rq.NoError(err)
	})
}
