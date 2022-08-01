package config

import (
	"encoding/json"
	"os"
	"path"
	"testing"

	"github.com/EmmettCorp/delorean/pkg/domain"
	"github.com/stretchr/testify/require"
)

func TestSave(t *testing.T) {
	t.Parallel()
	rq := require.New(t)

	t.Run("ok", func(t *testing.T) {
		t.Parallel()
		dir := "test_save"
		err := domain.CheckDir(dir, 0o777)
		rq.NoError(err)

		ph := path.Join(dir, "config.json")

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

		_, err = os.Stat(path.Join(ph, domain.Manual.String()))
		rq.NoError(err)

		_, err = os.Stat(path.Join(ph, domain.Monthly.String()))
		rq.NoError(err)

		_, err = os.Stat(path.Join(ph, domain.Weekly.String()))
		rq.NoError(err)

		_, err = os.Stat(path.Join(ph, domain.Daily.String()))
		rq.NoError(err)

		_, err = os.Stat(path.Join(ph, domain.Hourly.String()))
		rq.NoError(err)

		_, err = os.Stat(path.Join(ph, domain.Boot.String()))
		rq.NoError(err)

		_, err = os.Stat(path.Join(ph, domain.Restore.String()))
		rq.NoError(err)

		err = os.RemoveAll(ph)
		rq.NoError(err)
	})
}
