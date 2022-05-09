package shared

import (
	"testing"

	"github.com/EmmettCorp/delorean/pkg/config"
	"github.com/stretchr/testify/require"
)

const (
	testRoot         = "root"
	testPathToRemove = "/some/path/to/remove"
)

func Test_NewState(t *testing.T) {
	t.Parallel()

	t.Run("ok", func(t *testing.T) {
		t.Parallel()

		rq := require.New(t)

		st := NewState(&config.Config{
			BtrfsSupported: true,
			RootDevice:     testRoot,
			ToRemove:       []string{testPathToRemove},
		})

		rq.True(st.Config.BtrfsSupported)
		rq.Equal(testRoot, st.Config.RootDevice)
		rq.Len(st.Config.ToRemove, 1)
		rq.Equal(testPathToRemove, st.Config.ToRemove[0])

		rq.NotNil(st.ClickableElements)

		for _, cm := range getAllClickableComponents() {
			_, ok := st.ClickableElements[cm]
			rq.False(ok)
		}

		rq.Equal(TabsBarHeight, st.Areas.TabBar.Height)
		rq.Equal(FullScreen, st.Areas.TabBar.Width)
	})
}

func Test_State_Update(t *testing.T) {
	t.Parallel()

	t.Run("ok", func(t *testing.T) {
		t.Parallel()

		rq := require.New(t)

		root := "root"
		somePathToRemove := "/some/path/to/remove"

		st := NewState(&config.Config{
			BtrfsSupported: true,
			RootDevice:     root,
			ToRemove:       []string{somePathToRemove},
		})

		rq.Equal(SnapshotsTab, st.CurrentTab)

		st.Update(SettingsTab)

		rq.Equal(SettingsTab, st.CurrentTab)
	})
}

func Test_State_AppendClickable(t *testing.T) {
	t.Parallel()

	t.Run("ok", func(t *testing.T) {
		t.Parallel()

		rq := require.New(t)

		st := NewState(&config.Config{
			BtrfsSupported: true,
			RootDevice:     testRoot,
			ToRemove:       []string{testPathToRemove},
		})

		rq.Empty(st.ClickableElements[SnapshotsButtonsBar])

		cl := clickableMock{}
		cl.SetCoords(Coords{
			X1: 1,
			Y1: 1,
			X2: 2,
			Y2: 2,
		})

		err := st.AppendClickable(SnapshotsButtonsBar, &cl)
		rq.NoError(err)
	})

	t.Run("err: validate", func(t *testing.T) {
		t.Parallel()

		rq := require.New(t)

		st := NewState(&config.Config{
			BtrfsSupported: true,
			RootDevice:     testRoot,
			ToRemove:       []string{testPathToRemove},
		})

		rq.Empty(st.ClickableElements[SnapshotsButtonsBar])

		cl := clickableMock{}
		cl.SetCoords(Coords{
			X1: 3,
			Y1: 3,
			X2: 2,
			Y2: 2,
		})

		err := st.AppendClickable(SnapshotsButtonsBar, &cl)
		rq.Error(err)
	})
}

func Test_State_CleanClickable(t *testing.T) {
	t.Parallel()

	t.Run("ok", func(t *testing.T) {
		t.Parallel()

		rq := require.New(t)

		st := NewState(&config.Config{
			BtrfsSupported: true,
			RootDevice:     testRoot,
			ToRemove:       []string{testPathToRemove},
		})

		rq.Empty(st.ClickableElements[SnapshotsButtonsBar])

		cl := clickableMock{}
		cl.SetCoords(Coords{
			X1: 1,
			Y1: 1,
			X2: 2,
			Y2: 2,
		})

		err := st.AppendClickable(SnapshotsButtonsBar, &cl)
		rq.NoError(err)

		tags, ok := st.ClickableElements[SnapshotsButtonsBar]
		rq.True(ok)
		rq.Len(tags, 1)

		st.CleanClickable(SnapshotsButtonsBar)

		tags, ok = st.ClickableElements[SnapshotsButtonsBar]
		rq.True(ok)
		rq.Len(tags, 0)
	})
}

func Test_State_FindClickable(t *testing.T) {
	t.Parallel()

	t.Run("ok", func(t *testing.T) {
		t.Parallel()

		rq := require.New(t)

		st := NewState(&config.Config{
			BtrfsSupported: true,
			RootDevice:     testRoot,
			ToRemove:       []string{testPathToRemove},
		})

		rq.Empty(st.ClickableElements[SnapshotsList])

		line1 := clickableMock{}
		line1.SetCoords(Coords{
			X1: 0,
			Y1: 0,
			X2: 20,
			Y2: 4,
		})
		button1 := clickableMock{}
		button1.SetCoords(Coords{
			X1: 1,
			Y1: 1,
			X2: 3,
			Y2: 3,
		})

		line2 := clickableMock{}
		line2.SetCoords(Coords{
			X1: 5,
			Y1: 5,
			X2: 20,
			Y2: 9,
		})
		button2 := clickableMock{}
		button2.SetCoords(Coords{
			X1: 1,
			Y1: 6,
			X2: 3,
			Y2: 8,
		})

		err := st.AppendClickable(SnapshotsList, &line1)
		rq.NoError(err)
		err = st.AppendClickable(SnapshotsList, &button1)
		rq.NoError(err)
		err = st.AppendClickable(SnapshotsList, &line2)
		rq.NoError(err)
		err = st.AppendClickable(SnapshotsList, &button2)
		rq.NoError(err)

		tags, ok := st.ClickableElements[SnapshotsList]
		rq.True(ok)
		rq.Len(tags, 4)

		l1 := st.FindClickable(5, 1)
		rq.Equal(line1.GetCoords(), l1.GetCoords())
		b1 := st.FindClickable(2, 2)
		rq.Equal(button1.GetCoords(), b1.GetCoords())

		l2 := st.FindClickable(5, 6)
		rq.Equal(line2.GetCoords(), l2.GetCoords())
		b2 := st.FindClickable(3, 6)
		rq.Equal(button2.GetCoords(), b2.GetCoords())
	})
}

func Test_ResizeAreas(t *testing.T) {
	t.Parallel()
	t.Run("ok", func(t *testing.T) {
		t.Parallel()

		rq := require.New(t)

		st := NewState(&config.Config{
			BtrfsSupported: true,
			RootDevice:     testRoot,
			ToRemove:       []string{testPathToRemove},
		})

		rq.Equal(0, st.Areas.MainContent.Height)

		st.ScreenHeight = 10
		st.ResizeAreas()

		rq.Equal(st.ScreenHeight-(TabsBarHeight+HelpBarHeight), st.Areas.MainContent.Height)

		st.ScreenHeight = 40
		st.ResizeAreas()

		rq.Equal(st.ScreenHeight-(TabsBarHeight+HelpBarHeight), st.Areas.MainContent.Height)
	})
}
