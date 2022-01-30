package domain

import (
	"os"
	"testing"

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

		err = CheckDir(ph, 0o006)
		rq.NoError(err)

		_, err = os.Stat(ph)
		rq.NoError(err)

		err = os.Remove(ph)
		rq.NoError(err)
	})
}
