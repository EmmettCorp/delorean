package tabs

import (
	"github.com/EmmettCorp/delorean/pkg/ui/shared"
	tea "github.com/charmbracelet/bubbletea"
)

type tab struct {
	id     shared.TabItem
	state  *shared.State
	coords shared.Coords
	title  string
}

func NewTab(state *shared.State, id shared.TabItem, coords shared.Coords) (*tab, error) {
	t := tab{
		title: id.String(),
		state: state,
		id:    id,
	}
	t.SetCoords(coords)
	err := t.state.AppendClickable(shared.AnyTab, &t)
	if err != nil {
		return nil, err
	}

	return &t, nil
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
	t.state.Update(t.id)
}
