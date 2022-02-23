package colors

import (
	"testing"

	"github.com/jroimartin/gocui"
	"github.com/stretchr/testify/require"
)

func Test_GetColorByName(t *testing.T) {
	t.Parallel()
	rq := require.New(t)

	t.Run("all colors", func(t *testing.T) {
		t.Parallel()

		rq.Equal(gocui.ColorBlack, GetColorByName("black"))
		rq.Equal(gocui.ColorWhite, GetColorByName("white"))
		rq.Equal(gocui.ColorBlue, GetColorByName("blue"))
		rq.Equal(gocui.ColorGreen, GetColorByName("green"))
		rq.Equal(gocui.ColorMagenta, GetColorByName("magenta"))
		rq.Equal(gocui.ColorRed, GetColorByName("red"))
		rq.Equal(gocui.ColorYellow, GetColorByName("yellow"))
		rq.Equal(gocui.ColorCyan, GetColorByName("cyan"))
		rq.Equal(gocui.ColorDefault, GetColorByName(""))
		rq.Equal(gocui.ColorDefault, GetColorByName("some weird value"))
	})
}

func Test_GetHighlightBG(t *testing.T) {
	t.Parallel()
	rq := require.New(t)

	t.Run("all colors", func(t *testing.T) {
		t.Parallel()

		rq.Equal(gocui.ColorBlack, GetHighlightBG("black"))
		rq.Equal(gocui.ColorWhite, GetHighlightBG("white"))
		rq.Equal(gocui.ColorBlue, GetHighlightBG("blue"))
		rq.Equal(gocui.ColorGreen, GetHighlightBG("green"))
		rq.Equal(gocui.ColorMagenta, GetHighlightBG("magenta"))
		rq.Equal(gocui.ColorRed, GetHighlightBG("red"))
		rq.Equal(gocui.ColorYellow, GetHighlightBG("yellow"))
		rq.Equal(gocui.ColorCyan, GetHighlightBG("cyan"))
		rq.Equal(defaultHighlightBG, GetHighlightBG(""))
		rq.Equal(defaultHighlightBG, GetHighlightBG("some weird value"))
	})
}
