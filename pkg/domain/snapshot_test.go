package domain

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSetLabel(t *testing.T) {
	t.Parallel()

	t.Run("empty", func(t *testing.T) {
		t.Parallel()

		s := Snapshot{}
		s.SetLabel()

		rq := require.New(t)

		rq.Empty(s.Label)
	})

	t.Run("set", func(t *testing.T) {
		t.Parallel()

		s := Snapshot{
			Path: "/home",
		}
		s.SetLabel()

		rq := require.New(t)

		rq.Equal("home", s.Label)
	})
}

func TestSetType(t *testing.T) {
	t.Parallel()

	t.Run("empty", func(t *testing.T) {
		t.Parallel()

		s := Snapshot{}
		s.SetType()

		rq := require.New(t)

		rq.Empty(s.Label)
	})

	t.Run("set", func(t *testing.T) {
		t.Parallel()

		s := Snapshot{
			Path: "/home/manual/time",
		}
		s.SetType()

		rq := require.New(t)

		rq.Equal("manual", s.Type)
	})
}

func TestSetTimestamp(t *testing.T) {
	t.Parallel()

	t.Run("empty", func(t *testing.T) {
		t.Parallel()

		s := Snapshot{}
		s.SetTimestamp()

		rq := require.New(t)

		rq.Empty(s.Timestamp)
	})

	t.Run("set", func(t *testing.T) {
		t.Parallel()

		s := Snapshot{
			Label: "2022-01-25_16:56:37",
		}
		s.SetTimestamp()

		rq := require.New(t)

		rq.Equal(int64(1643129797), s.Timestamp)
	})
}
