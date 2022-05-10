package shared

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_TabItem_String(t *testing.T) {
	t.Parallel()

	t.Run("ok", func(t *testing.T) {
		t.Parallel()

		rq := require.New(t)

		rq.Equal(snapshotsTabTitle, SnapshotsTab.String())
		rq.Equal(settingsTabTitle, SettingsTab.String())
		rq.Equal("", AnyTab.String())
	})
}
