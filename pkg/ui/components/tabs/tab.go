package tabs

import (
	"github.com/EmmettCorp/delorean/pkg/ui/shared"
	tea "github.com/charmbracelet/bubbletea"
)

type Tab struct {
	shared.Clickable
	title string
	state *shared.State
	id    shared.TabItem
}

func NewTab(state *shared.State, id shared.TabItem, title string, X1, Y1, X2, Y2 int) *Tab {
	t := Tab{
		title: title,
		state: state,
		id:    id,
	}
	t.X1, t.Y1, t.X2, t.Y2 = X1, Y1, X2, Y2

	return &t
}

func (t *Tab) OnClick(event tea.MouseMsg) {
	t.state.CurrentTab = t.id
}

func (t *Tab) GetCoords() (int, int, int, int) {
	return t.X1, t.Y1, t.X2, t.Y2
}
