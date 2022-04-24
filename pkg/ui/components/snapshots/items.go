package snapshots

import (
	"fmt"
	"io"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type itemDelegate struct {
}

func (d itemDelegate) Height() int                               { return 1 }
func (d itemDelegate) Spacing() int                              { return 0 }
func (d itemDelegate) Update(msg tea.Msg, m *list.Model) tea.Cmd { return nil }
func (d itemDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	i, ok := listItem.(snapshot)
	if !ok {
		return
	}

	str := fmt.Sprintf("%s\t\t\t%d\t\t%s\n%s | %s", i.Label, index+1, i.Type, i.VolumeLabel, "5.16.20-200.fc35.x86_64")

	fn := itemStyle.Render
	if index == m.Index() {
		fn = func(s string) string {
			return selectedItemStyle.Render(s)
		}
	}

	fmt.Fprintf(w, fn(str))
}
