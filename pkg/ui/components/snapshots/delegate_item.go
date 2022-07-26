package snapshots

import (
	"fmt"
	"io"
	"strings"

	"github.com/EmmettCorp/delorean/pkg/logger"
	"github.com/EmmettCorp/delorean/pkg/ui/shared"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type rowAction int

const (
	deleteItem rowAction = iota
	restoreItem
)

const (
	restoreIcon        = "⟲"
	deleteIcon         = "✖"
	iconsGap           = 4
	itemDelegateHeight = 2
	spacing            = 1
)

type rowButton struct {
	shared.ClickableItem
	row *itemDelegate
}

// itemDelegate is responsible for item rendering.
type itemDelegate struct {
	shared.ClickableItem
	index int
	model *Model
}

func (d *itemDelegate) Height() int  { return itemDelegateHeight }
func (d *itemDelegate) Spacing() int { return spacing }
func (d *itemDelegate) Update(msg tea.Msg, m *list.Model) tea.Cmd {
	return nil
}

func (d *itemDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	s, ok := listItem.(*snapshot)
	if !ok {
		return
	}

	if d.model.updateClickable {
		d.setRowClickable(index, m.Paginator.PerPage)
	}

	var rowBuilder strings.Builder
	rowBuilder.WriteString(s.Label)
	rowBuilder.WriteString(strings.Repeat(" ", infoColumnWidth-lipgloss.Width(s.Label)-minColumnGapLen))
	rowBuilder.WriteString(s.VolumeID)
	rowBuilder.WriteString(strings.Repeat(" ", idColumnWidth-lipgloss.Width(s.VolumeID)))
	rowBuilder.WriteString(s.Type)
	row := rowBuilder.String()
	// restoreItem is left most button in row
	// minus 1 here because getRowButtonX1 for restore button returns it's x1 value, we need get -1 of it
	gap := strings.Repeat(" ", shared.Max(minColumnGapLen,
		d.getRowButtonX1(restoreItem)-1-lipgloss.Width(row)-spacing))
	itemRow := lipgloss.JoinHorizontal(lipgloss.Left, row, gap, restoreIcon, strings.Repeat(" ", iconsGap), deleteIcon)

	var description string
	if s.VolumeLabel == "Root" {
		krn := s.Kernel
		if krn == d.model.state.Config.KernelVersion {
			krn = "current"
		}
		description = fmt.Sprintf("volume: %s | kernel: %s ", s.VolumeLabel, krn)
	} else {
		description = fmt.Sprintf("volume: %s", s.VolumeLabel)
	}

	if index == m.Index() {
		itemRow = d.model.styles.SelectedTitle.Render(itemRow)
		description = d.model.styles.SelectedDesc.Render(description)
	} else {
		itemRow = d.model.styles.NormalTitle.Render(itemRow)
		description = d.model.styles.NormalDesc.Render(description)
	}

	fmt.Fprintf(w, "%s\n%s", itemRow, description)
}

func (d *itemDelegate) getFirstItemY() int {
	return d.model.state.Areas.TabBar.Height + createButtonHeight + listHeaderHeight
}

func (d itemDelegate) setRowClickable(index, perPage int) {
	d.index = index
	itemY := d.getFirstItemY() + (spacing+itemDelegateHeight)*(index%perPage)

	d.SetCoords(shared.Coords{
		X1: spacing,
		Y1: itemY,
		X2: d.model.state.ScreenWidth,
		Y2: itemY + spacing,
	})
	d.SetCallback(func() error {
		return d.model.selectByIndex(d.getIndex())
	})

	err := d.model.state.AppendClickable(shared.SnapshotsList, &d)
	if err != nil {
		logger.Client.ErrLog.Printf("append clickable row `%d`: %v", index, err)
	}

	deleteX1 := d.getRowButtonX1(deleteItem)
	deleteItem := rowButton{
		row: &d,
	}
	deleteItem.SetCoords(shared.Coords{
		X1: deleteX1,
		Y1: itemY,
		X2: deleteX1 + lipgloss.Width(deleteIcon),
		Y2: itemY + itemDelegateHeight,
	})

	deleteItem.SetCallback(func() error {
		return d.model.deleteWithDialog(d.index)
	})
	err = d.model.state.AppendClickable(shared.SnapshotsList, &deleteItem)
	if err != nil {
		logger.Client.ErrLog.Printf("append clickable delete button `%d`: %v", index, err)
	}
}

func (d *itemDelegate) getRowButtonX1(action rowAction) int {
	switch action {
	case deleteItem:
		return shared.Max(infoColumnsWidth+lipgloss.Width(restoreIcon)+iconsGap,
			d.model.state.ScreenWidth-spacing-lipgloss.Width(deleteIcon))
	case restoreItem:
		return shared.Max(infoColumnsWidth,
			d.model.state.ScreenWidth-spacing-lipgloss.Width(deleteIcon)-iconsGap-lipgloss.Width(restoreIcon))
	default:
		return 1
	}
}

func (d *itemDelegate) getIndex() int {
	return d.index
}
