package btrfs

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestOsReadDir(t *testing.T) {
	t.Parallel()
	rq := require.New(t)

	t.Run("ok", func(t *testing.T) {
		t.Parallel()

		ff, err := osReadDir(".")
		rq.NoError(err)
		rq.Len(ff, 0)

		ff, err = osReadDir("./..")
		rq.NoError(err)
		rq.Contains(ff, "btrfs")
	})
}

func TestGetSnapshots(t *testing.T) {
	t.Parallel()
	rq := require.New(t)

	t.Run("err", func(t *testing.T) {
		t.Parallel()

		ff, err := getSnapshots("./..")
		rq.NoError(err)
		rq.Len(ff, 0)
	})

	t.Run("ok", func(t *testing.T) {
		t.Parallel()

		ff, err := getSnapshots("./../..")
		rq.NoError(err)
		rq.Contains(ff, "../../commands/btrfs")
		rq.Contains(ff, "../../commands/findmnt")
	})
}
