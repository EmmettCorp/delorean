/*
Package colors keeps helpers to colorize text output.
*/
package colors

import "fmt"

// Color represents the terminal color in int.
type Color int

// Available colors.
const (
	Red    Color = 31
	Green  Color = 32
	Yellow Color = 33
)

// Paint colorizes string `s` to color `c`.
func Paint(s string, c Color) string {
	return fmt.Sprintf("\033[1;%dm%s\033[0m", c, s)
}

// Bold makes string `s` bold.
func Bold(s string) string {
	return fmt.Sprintf("%s%s\033[0m", "\033[1m", s)
}
