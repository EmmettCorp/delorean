/*
Package button keeps helpers to draw buttons.
*/
package button

func New(title string) string {
	return buttonWithBorder.Render(title)
}
