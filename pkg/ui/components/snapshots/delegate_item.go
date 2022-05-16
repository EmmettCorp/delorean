package snapshots

import (
	"fmt"
	"io"
	"strings"

	"github.com/EmmettCorp/delorean/pkg/ui/shared"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const (
	restoreIcon        = "↻"
	deleteIcon         = "✖"
	iconsGap           = 4
	itemDelegateHeight = 2
)

// itemDelegate is responsible for item rendering
type itemDelegate struct {
	state  *shared.State // state here is needed to track screenWidth
	styles list.DefaultItemStyles
}

func (d itemDelegate) Height() int                               { return itemDelegateHeight }
func (d itemDelegate) Spacing() int                              { return 1 }
func (d itemDelegate) Update(msg tea.Msg, m *list.Model) tea.Cmd { return nil }
func (d itemDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	s, ok := listItem.(*snapshot)
	if !ok {
		return
	}

	var rowBuilder strings.Builder
	rowBuilder.WriteString(s.Label)
	rowBuilder.WriteString(strings.Repeat(" ", infoColumnWidth-lipgloss.Width(s.Label)-minColumnGapLen))
	rowBuilder.WriteString(s.VolumeID)
	rowBuilder.WriteString(strings.Repeat(" ", idColumnWidth-lipgloss.Width(s.VolumeID)))
	rowBuilder.WriteString(s.Type)
	row := rowBuilder.String()

	rowIcons := fmt.Sprintf("%s%s%s", restoreIcon, strings.Repeat(" ", iconsGap), deleteIcon)

	gap := strings.Repeat(" ", shared.Max(minColumnGapLen, d.state.ScreenWidth-lipgloss.Width(row)-lipgloss.Width(rowIcons)-iconsGap))

	title := lipgloss.JoinHorizontal(lipgloss.Left, row, gap, rowIcons)

	var description string
	if s.VolumeLabel == "Root" {
		description = fmt.Sprintf("volume: %s | kernel: %s ", s.VolumeLabel, s.Kernel)
	} else {
		description = fmt.Sprintf("volume: %s", s.VolumeLabel)
	}

	if index == m.Index() {
		title = d.styles.SelectedTitle.Render(title)
		description = d.styles.SelectedDesc.Render(description)
	} else {
		title = d.styles.NormalTitle.Render(title)
		description = d.styles.NormalDesc.Render(description)
	}

	fmt.Fprintf(w, "%s\n%s", title, description)
}
