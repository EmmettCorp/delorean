package colors

import (
	"errors"

	"github.com/jroimartin/gocui"
)

// GetColorByName returns the gocui.Attribute color if name is appropriate.
func GetColorByName(color string) (gocui.Attribute, error) {
	switch color {
	case "black":
		return gocui.ColorBlack, nil
	case "white":
		return gocui.ColorWhite, nil
	case "blue":
		return gocui.ColorBlue, nil
	case "green":
		return gocui.ColorGreen, nil
	case "magenta":
		return gocui.ColorMagenta, nil
	case "red":
		return gocui.ColorRed, nil
	case "yellow":
		return gocui.ColorYellow, nil
	case "cyan":
		return gocui.ColorCyan, nil
	default:
		return 0, errors.New("unknown color")
	}
}
