package rate

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestSetTimer(t *testing.T) {
	t.Parallel()

	t.Run("ok", func(t *testing.T) {
		t.Parallel()

		lim := NewLimiter(time.Second)
		lim.expired = false
		lim.setNewTimer()

		rq := require.New(t)
		time.Sleep(time.Second * 2)
		rq.True(lim.expired)
	})
}

func TestAllow(t *testing.T) {
	t.Parallel()

	t.Run("ok", func(t *testing.T) {
		t.Parallel()

		lim := NewLimiter(time.Second)

		rq := require.New(t)
		rq.True(lim.Allow())
		rq.False(lim.Allow())
		time.Sleep(time.Second * 2)
		rq.True(lim.Allow())
	})
}
