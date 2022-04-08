package tabs

import (
	"github.com/EmmettCorp/delorean/pkg/ui/shared"
	tea "github.com/charmbracelet/bubbletea"
)

type tab struct {
	title  string
	state  *shared.State
	id     shared.TabItem
	coords shared.Coords
}

func NewTab(state *shared.State, id shared.TabItem, coords shared.Coords) *tab {
	t := tab{
		title: id.String(),
		state: state,
		id:    id,
	}
	t.SetCoords(coords)
	t.state.AppendClickable(&t)
	return &t
}

func (t *tab) getTitle() string {
	return t.title
}

func (t *tab) getID() shared.TabItem {
	return t.id
}

func (t *tab) GetCoords() shared.Coords {
	return t.coords
}

func (t *tab) SetCoords(c shared.Coords) {
	t.coords = c
}

func (t *tab) OnClick(event tea.MouseMsg) {
	t.state.CurrentTab = t.id
}
