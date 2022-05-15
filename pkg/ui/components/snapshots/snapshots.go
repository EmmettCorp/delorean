/*
Package snapshots keeps all the logic for snapshots component.
*/
package snapshots

import (
	"strings"

	"github.com/EmmettCorp/delorean/pkg/commands/btrfs"
	"github.com/EmmettCorp/delorean/pkg/ui/shared"
	"github.com/EmmettCorp/delorean/pkg/ui/shared/elements/button"
	"github.com/EmmettCorp/delorean/pkg/ui/shared/elements/divider"
	"github.com/EmmettCorp/delorean/pkg/ui/shared/styles"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const (
	infoTitle            = "Info"
	idTitle              = "ID"
	typeTitle            = "Type"
	infoColumnWidth      = 30
	idColumnWidth        = 6
	tabLineDeviderHeight = 4

	minColumnGap    = "  "
	minColumnGapLen = len(minColumnGap)
)

type buttonModel interface {
	shared.Clickable
	SetTitle(title string)
	GetTitle() string
}

type snapshot struct {
	Label       string
	VolumeLabel string
	Type        string
	VolumeID    string
	Kernel      string
}

func (s *snapshot) FilterValue() string { return s.Label }

type Model struct {
	state       *shared.State
	createBtn   buttonModel
	list        list.Model
	height      int
	currentPage int
	itemsCount  int
	err         error
}

func NewModel(st *shared.State) (*Model, error) {
	m := Model{
		state:       st,
		currentPage: -1,
		itemsCount:  -1,
	}

	itemsModel := list.New([]list.Item{}, itemDelegate{
		state:  st,
		styles: list.NewDefaultItemStyles(),
	}, 0, 0)
	itemsModel.SetFilteringEnabled(false)
	itemsModel.SetShowFilter(false)
	itemsModel.SetShowTitle(false)
	itemsModel.SetShowStatusBar(false)
	itemsModel.SetShowHelp(false)
	m.list = itemsModel
	m.UpdateList()

	btnTitle := "Create"
	createButtongY1 := st.Areas.TabBar.Height + 1
	createBtn := newCreateButton(st, btnTitle, shared.Coords{
		Y1: createButtongY1,
		X2: len(btnTitle) + 3, // nolint:gomnd // left and right borders + 1
		Y2: createButtongY1 + CreateButtonHeight,
	}, m.UpdateList)
	m.createBtn = createBtn

	err := st.AppendClickable(shared.SnapshotsButtonsBar, createBtn)
	if err != nil {
		return nil, err
	}

	return &m, nil
}

func (m *Model) Init() tea.Cmd {
	return nil
}

func (m *Model) View() string {
	var s strings.Builder
	s.WriteString(button.New(m.createBtn.GetTitle()))
	s.WriteString("\n")
	s.WriteString(lipgloss.NewStyle().SetString(getSnapshotsHeader()).
		Foreground(styles.DefaultTheme.InactiveText).String())
	s.WriteString("\n")
	s.WriteString(divider.HorizontalLine(m.state.ScreenWidth, styles.DefaultTheme.InactiveText))
	s.WriteString("\n")
	m.list.SetSize(m.state.ScreenWidth, m.height)
	s.WriteString(docStyle.Render(m.list.View()))

	return s.String()
}

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var needUpdateClickable bool

	if _, ok := msg.(tea.WindowSizeMsg); ok {
		m.height = m.getHeight()
		m.list.SetSize(m.state.ScreenWidth, m.height)
		needUpdateClickable = true
	}

	var cmd tea.Cmd
	// do not call btrfs commands for just ui update
	if len(m.list.Items()) != m.itemsCount {
		m.UpdateList()
		m.itemsCount = len(m.list.Items())
	}
	m.list, cmd = m.list.Update(msg)

	if m.currentPage != m.list.Paginator.Page {
		m.currentPage = m.list.Paginator.Page
		needUpdateClickable = true
	}

	if needUpdateClickable {
		updateClickable(m)
	}

	return m, cmd
}

func (m *Model) UpdateList() {
	snaps, err := btrfs.SnapshotsList(m.state.Config.Volumes)
	if err != nil {
		m.err = err

		return
	}

	items := make([]list.Item, len(snaps))
	for i := range snaps {
		sn := snapshot{
			Label:       snaps[i].Label,
			VolumeLabel: snaps[i].VolumeLabel,
			Type:        snaps[i].Type,
			VolumeID:    snaps[i].VolumeID,
		}
		items[i] = &sn
	}

	m.list.SetItems(items)
}

func (m *Model) getHeight() int {
	return m.state.Areas.MainContent.Height - (CreateButtonHeight + tabLineDeviderHeight)
}

func getSnapshotsHeader() string {
	var header strings.Builder
	header.WriteString(minColumnGap)
	header.WriteString(infoTitle)
	header.WriteString(strings.Repeat(" ", infoColumnWidth-len(infoTitle)-minColumnGapLen))
	header.WriteString(idTitle)
	header.WriteString(strings.Repeat(" ", idColumnWidth-len(idTitle)))
	header.WriteString(typeTitle)

	return header.String()
}
