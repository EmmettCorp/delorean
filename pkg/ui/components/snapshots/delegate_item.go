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
	restoreIcon        = "↻"
	deleteIcon         = "✖"
	iconsGap           = 4
	itemDelegateHeight = 2
	spacing            = 1
)

type rowButton struct {
	action rowAction
	row    *itemDelegate
	coords shared.Coords
}

// itemDelegate is responsible for item rendering.
type itemDelegate struct {
	index  int
	model  *Model
	coords shared.Coords
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
		err := d.setRowClickable(index, m.Paginator.PerPage)
		if err != nil {
			logger.Client.ErrLog.Printf("can't prepare visible row with index `%d`: %v", index, err)
		}
	}

	var rowBuilder strings.Builder
	rowBuilder.WriteString(s.Label)
	rowBuilder.WriteString(strings.Repeat(" ", infoColumnWidth-lipgloss.Width(s.Label)-minColumnGapLen))
	rowBuilder.WriteString(s.VolumeID)
	rowBuilder.WriteString(strings.Repeat(" ", idColumnWidth-lipgloss.Width(s.VolumeID)))
	rowBuilder.WriteString(s.Type)
	row := rowBuilder.String()
	// minus 1 here because getRowButtonX1 for restore button returns it's x1 value, we need get -1 of it
	gap := strings.Repeat(" ", shared.Max(minColumnGapLen,
		d.getRowButtonX1(restoreItem)-1-lipgloss.Width(row)-spacing))
	itemRow := lipgloss.JoinHorizontal(lipgloss.Left, row, gap, restoreIcon, strings.Repeat(" ", iconsGap), deleteIcon)

	var description string
	if s.VolumeLabel == "Root" {
		description = fmt.Sprintf("volume: %s | kernel: %s ", s.VolumeLabel, s.Kernel)
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
	return d.model.state.Areas.TabBar.Height + createButtonHeight + tabLineDividerHeight
}

func (d itemDelegate) setRowClickable(index, perPage int) error {
	d.index = index
	itemY := d.getFirstItemY() + (spacing+itemDelegateHeight)*(index%perPage)
	d.coords = shared.Coords{
		X1: 1,
		Y1: itemY,
		X2: d.model.state.ScreenWidth,
		Y2: itemY + spacing,
	}
	err := d.model.state.AppendClickable(shared.SnapshotsList, &d)
	if err != nil {
		return fmt.Errorf("can't set row as clickable: %v", err)
	}

	deleteItem := rowButton{
		coords: shared.Coords{
			X1: d.getRowButtonX1(deleteItem),
			Y1: itemY,
			X2: d.getRowButtonX1(deleteItem) + lipgloss.Width(deleteIcon),
			Y2: itemY + spacing,
		},
		row: &d,
	}

	return d.model.state.AppendClickable(shared.SnapshotsList, &deleteItem)
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

func (d *itemDelegate) OnClick(event tea.MouseMsg) error {
	d.model.list.Select(d.index)

	return nil
}

func (d *itemDelegate) GetCoords() shared.Coords {
	return d.coords
}

func (d *itemDelegate) SetCoords(coords shared.Coords) {
	d.coords = coords
}

func (re *rowButton) OnClick(event tea.MouseMsg) error {
	switch re.action {
	case deleteItem:
		return re.deleteItem()
	case restoreItem:
		return nil // not implemented yet
	}

	return nil
}

func (re *rowButton) GetCoords() shared.Coords {
	return re.coords
}

func (re *rowButton) SetCoords(coords shared.Coords) {
	re.coords = coords
}

func (re *rowButton) deleteItem() error {
	return re.row.model.deleteByIndex(re.row.index)
}
