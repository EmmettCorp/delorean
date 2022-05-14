package snapshots

import (
	"github.com/EmmettCorp/delorean/pkg/ui/shared"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type clickableItem struct {
	index  int
	coords shared.Coords
	list   *list.Model
}

func (ci *clickableItem) OnClick(event tea.MouseMsg) error {
	ci.list.Select(ci.index)

	return nil
}

func (ci *clickableItem) GetCoords() shared.Coords {
	return ci.coords
}

func (ci *clickableItem) SetCoords(coords shared.Coords) {
	ci.coords = coords
}

func updateClickable(m *Model) error {
	h := m.state.Areas.TabBar.Height + CreateButtonHeight + tabLineDeviderHeight
	m.state.CleanClickable(shared.SnapshotsList)
	rowHeight := 1
	gap := 2
	itemY := h

	items := m.list.Items()
	first := m.list.Paginator.PerPage * m.list.Paginator.Page

	for i := first; i < first+m.list.Paginator.PerPage && i < len(items); i++ {
		sn := clickableItem{
			index: i,
			coords: shared.Coords{
				X1: 1,
				Y1: itemY,
				X2: m.state.ScreenWidth,
				Y2: itemY + rowHeight,
			},
			list: &m.list,
		}
		itemY = itemY + rowHeight + gap

		err := m.state.AppendClickable(shared.SnapshotsList, &sn)
		if err != nil {
			return err
		}
	}

	return nil
}
