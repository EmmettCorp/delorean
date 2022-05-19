package snapshots

import (
	"github.com/EmmettCorp/delorean/pkg/logger"
	"github.com/EmmettCorp/delorean/pkg/ui/shared"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const (
	clickableGap                   = 1
	clickableExtendedInfoRowHeight = 1
)

type (
	clickableItem struct {
		index  int
		coords shared.Coords
		model  *Model
	}

	deleteItemButton struct {
		coords shared.Coords
		item   *clickableItem
	}
)

func (ci *clickableItem) OnClick(event tea.MouseMsg) error {
	ci.model.list.Select(ci.index)

	return nil
}

func (ci *clickableItem) GetCoords() shared.Coords {
	return ci.coords
}

func (ci *clickableItem) SetCoords(coords shared.Coords) {
	ci.coords = coords
}

func (db *deleteItemButton) OnClick(event tea.MouseMsg) error {
	return db.item.model.deleteByIndex(db.item.index)
}

func (db *deleteItemButton) GetCoords() shared.Coords {
	return db.coords
}

func (db *deleteItemButton) SetCoords(coords shared.Coords) {
	db.coords = coords
}

func updateClickable(m *Model) {
	h := m.state.Areas.TabBar.Height + CreateButtonHeight + tabLineDividerHeight
	m.state.CleanClickable(shared.SnapshotsList)
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
				Y2: itemY + clickableExtendedInfoRowHeight,
			},
			model: m,
		}
		itemY = itemY + clickableExtendedInfoRowHeight + clickableGap + 1

		err := m.state.AppendClickable(shared.SnapshotsList, &sn)
		if err != nil {
			m.err = err
		}
		logger.Client.InfoLog.Println("clickable", shared.Max(m.state.ScreenWidth-(infoColumnWidth+idColumnWidth+tabLineDividerHeight+minColumnGapLen),
			m.state.ScreenWidth-clickableGap-lipgloss.Width(deleteIcon)))
		db := deleteItemButton{
			coords: shared.Coords{
				X1: shared.Max(m.state.ScreenWidth-(infoColumnWidth+idColumnWidth+tabLineDividerHeight+minColumnGapLen),
					m.state.ScreenWidth-clickableGap-lipgloss.Width(deleteIcon)),
				Y1: sn.coords.Y1,
				X2: shared.Max(m.state.ScreenWidth-(infoColumnWidth+idColumnWidth+tabLineDividerHeight+minColumnGapLen),
					m.state.ScreenWidth-clickableGap),
				Y2: sn.coords.Y2,
			},
			item: &sn,
		}
		err = m.state.AppendClickable(shared.SnapshotsList, &db)
		if err != nil {
			m.err = err
		}
	}
}
