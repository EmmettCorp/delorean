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
		s.SetLabel()

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
