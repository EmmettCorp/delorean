/*
Package colors keeps helpers to colorize text output.
*/
package colors

import "fmt"

const (
	start   = "\033[1;"
	end     = "\033[0m"
	fgRed   = 31
	fgGreen = 32
)

func FgRed(s string) string {
	return fmt.Sprintf("%s%dm%s%s", start, fgRed, s, end)
}

func FgGreen(s string) string {
	return fmt.Sprintf("%s%dm%s%s", start, fgGreen, s, end)
}

func Bold(s string) string {
	return fmt.Sprintf("%s%s%s", "\033[1m", s, end)
}
