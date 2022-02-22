package colors

import (
	"github.com/jroimartin/gocui"
)

// Absolutely disputable constant.
// By default highlight background gocui is gocui.ColorGreen that looks a bit more obvious and universal.
// Green one looks ok on black terminal theme and theme like ubuntu has out of the box.
// The reason why delorean has black color for highlight is only because author uses black background and in that case
// this value looks way better. ))
// But fortunately this options is changable in config. Look README.
const defaultHighlightBG = gocui.ColorBlack

// GetColorByName returns the gocui.Attribute color if name is appropriate.
// If invalid color name it returns ColorDefault.
func GetColorByName(color string) gocui.Attribute {
	switch color {
	case "black":
		return gocui.ColorBlack
	case "white":
		return gocui.ColorWhite
	case "blue":
		return gocui.ColorBlue
	case "green":
		return gocui.ColorGreen
	case "magenta":
		return gocui.ColorMagenta
	case "red":
		return gocui.ColorRed
	case "yellow":
		return gocui.ColorYellow
	case "cyan":
		return gocui.ColorCyan
	default:
		return gocui.ColorDefault
	}
}

// GetHighlightBG returns
func GetHighlightBG(color string) gocui.Attribute {
	c := GetColorByName(color)
	if c != gocui.ColorDefault {
		return c
	}

	return defaultHighlightBG
}
