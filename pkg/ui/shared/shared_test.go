package shared

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Max(t *testing.T) {
	t.Parallel()

	t.Run("ok", func(t *testing.T) {
		t.Parallel()
		rq := require.New(t)

		rq.Equal(5, Max(3, 5))
		rq.Equal(21, Max(21, 5))
	})
}
