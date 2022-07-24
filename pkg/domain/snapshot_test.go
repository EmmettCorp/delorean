package domain

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewSnapshot(t *testing.T) {
	t.Parallel()
	rq := require.New(t)

	t.Run("err: empty path", func(t *testing.T) {
		t.Parallel()

		_, err := SnapshotByPath("", "some_label", "some_id")

		rq.Error(err)
	})

	t.Run("err: invalid snapshot id", func(t *testing.T) {
		t.Parallel()

		_, err := SnapshotByPath("manual/invalid snapshot id", "some_label", "some_id")

		rq.Error(err)
	})

	t.Run("err: invalid snapshot id", func(t *testing.T) {
		t.Parallel()

		sn, err := SnapshotByPath("/run/delorean/.snapshots/@/manual/2022-01-25_16:56:37", "some_label", "some_id")

		rq.NoError(err)
		rq.Equal("/run/delorean/.snapshots/@/manual/2022-01-25_16:56:37", sn.Path)
		rq.Equal("manual", sn.Type)
		rq.Equal("2022-01-25_16:56:37", sn.Label)
		rq.Equal("some_label", sn.VolumeLabel)
		rq.Equal("some_id", sn.VolumeID)
		rq.Equal(int64(1643129797), sn.Timestamp)
	})
}
