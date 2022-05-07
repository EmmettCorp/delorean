package shared

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_initAreas(t *testing.T) {
	t.Parallel()
	t.Run("ok", func(t *testing.T) {
		t.Parallel()

		rq := require.New(t)

		ua := initAreas()

		rq.Equal(TabsBarHeight, ua.TabBar.Height)
		rq.Equal(FullScreen, ua.TabBar.Width)

		rq.Equal(HelpBarHeight, ua.HelpBar.Height)
		rq.Equal(FullScreen, ua.HelpBar.Width)

		rq.Equal(0, ua.MainContent.Height)
		rq.Equal(FullScreen, ua.MainContent.Width)
	})
}
