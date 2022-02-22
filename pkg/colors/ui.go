package colors

import (
	"github.com/jroimartin/gocui"
)

// Absolutely disputable constant.
// By default highlight background gocui is green that seems a bit more obvious and universal.
// Green one looks ok on the terminal theme with black background
// and background that ubuntu has out of the box (burgundy I guess) or white.
// The reason why delorean has black color for highlight is only because author uses black background and in that case
// this value looks way better. Author naively believes that all around uses black backround. ))
// But fortunately this option is changable and can be set in config. Look README.
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

// GetHighlightBG returns color by name if set and valid otherwise defaultHighlightBG.
func GetHighlightBG(color string) gocui.Attribute {
	c := GetColorByName(color)
	if c != gocui.ColorDefault {
		return c
	}

	return defaultHighlightBG
}
