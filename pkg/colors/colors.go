package colors

import "fmt"

const (
	start    = "\033[1;"
	end      = "\033[0m"
	fg_red   = 31
	fg_green = 32
)

func FgRed(s string) string {
	return fmt.Sprintf("%s%dm%s%s", start, fg_red, s, end)
}

func FgGreen(s string) string {
	return fmt.Sprintf("%s%dm%s%s", start, fg_green, s, end)
}

func Bold(s string) string {
	return fmt.Sprintf("%s%s%s", "\033[1m", s, end)
}
