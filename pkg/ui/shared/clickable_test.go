package shared

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_validateClickable(t *testing.T) {
	t.Parallel()

	t.Run("err: nil", func(t *testing.T) {
		t.Parallel()

		rq := require.New(t)

		rq.Error(validateClickable(nil))
	})

	t.Run("ok", func(t *testing.T) {
		t.Parallel()

		rq := require.New(t)

		cm := &clickableMock{}
		cm.SetCoords(Coords{
			X2: 1,
			Y2: 1,
		})

		rq.NoError(validateClickable(cm))
	})

	t.Run("err: coords are not set", func(t *testing.T) {
		t.Parallel()

		rq := require.New(t)

		cm := &clickableMock{}

		rq.Error(validateClickable(cm))
	})

	t.Run("err: coords are not correct", func(t *testing.T) {
		t.Parallel()

		rq := require.New(t)

		cm := &clickableMock{}
		cm.SetCoords(Coords{
			X1: 22,
			Y1: 22,
			X2: 12,
			Y2: 12,
		})

		rq.Error(validateClickable(cm))
	})
}
