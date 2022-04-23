/*
Package tab keeps helpers to create tabs.
*/
package tab

import (
	"github.com/EmmettCorp/delorean/pkg/ui/shared"
	tea "github.com/charmbracelet/bubbletea"
)

type Tab struct {
	id     shared.TabItem
	state  *shared.State
	coords shared.Coords
	title  string
}

func New(state *shared.State, id shared.TabItem, coords shared.Coords) (*Tab, error) {
	t := Tab{
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

func (t *Tab) GetTitle() string {
	return t.title
}

func (t *Tab) GetID() shared.TabItem {
	return t.id
}

func (t *Tab) GetCoords() shared.Coords {
	return t.coords
}

func (t *Tab) SetCoords(c shared.Coords) {
	t.coords = c
}

func (t *Tab) OnClick(event tea.MouseMsg) error {
	t.state.Update(t.id)

	return nil
}
