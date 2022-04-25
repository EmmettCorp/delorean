/*
Package snapshots keeps all the logic for snapshots component.
*/
package snapshots

import (
	"strings"

	"github.com/EmmettCorp/delorean/pkg/commands/btrfs"
	"github.com/EmmettCorp/delorean/pkg/ui/elements/button"
	"github.com/EmmettCorp/delorean/pkg/ui/elements/divider"
	"github.com/EmmettCorp/delorean/pkg/ui/shared"
	"github.com/EmmettCorp/delorean/pkg/ui/shared/styles"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var screenWidth = 0

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

func (s snapshot) FilterValue() string { return s.Label }

type Model struct {
	state     *shared.State
	createBtn buttonModel
	list      list.Model
	height    int
	err       error
}

func NewModel(st *shared.State) (*Model, error) {
	m := Model{
		state: st,
	}

	screenWidth = m.state.ScreenWidth

	itemsModel := list.New([]list.Item{},
		itemDelegate{
			state:  st,
			styles: list.NewDefaultItemStyles(),
		},
		0, 0)
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

	err := st.AppendClickable(shared.SnapshotsTab, createBtn)
	if err != nil {
		return nil, err
	}

	return &m, nil
}

func (m *Model) Init() tea.Cmd {
	return nil
}

func (m *Model) View() string {
	s := strings.Builder{}
	s.WriteString(button.New(m.createBtn.GetTitle()))
	s.WriteString("\n")

	s.WriteString(lipgloss.NewStyle().SetString("  Info\t\t\t\tID\tType").
		Foreground(styles.DefaultTheme.InactiveText).String())
	s.WriteString("\n")
	s.WriteString(divider.HorizontalLine(m.state.ScreenWidth, styles.DefaultTheme.InactiveText))
	s.WriteString("\n")
	m.list.SetSize(m.state.ScreenWidth, m.height)
	s.WriteString(docStyle.Render(m.list.View()))

	return s.String()
}

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if msg, ok := msg.(tea.WindowSizeMsg); ok {
		screenWidth = m.state.ScreenWidth
		m.height = m.getHeight()
		h, v := docStyle.GetFrameSize()
		m.list.SetSize(msg.Width-h, msg.Height-v)
	}

	if len(m.list.Items()) == 0 {
		m.UpdateList()
	}
	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)

	return m, cmd
}

func (m *Model) UpdateList() {
	snaps, err := btrfs.SnapshotsList(m.state.Config.Volumes)
	if err != nil {
		m.err = err

		return
	}

	items := []list.Item{}
	for i := range snaps {
		items = append(items, snapshot{
			Label:       snaps[i].Label,
			VolumeLabel: snaps[i].VolumeLabel,
			Type:        snaps[i].Type,
			VolumeID:    snaps[i].VolumeID,
		})
	}

	m.list.SetItems(items)
}

func (m *Model) getHeight() int {
	return m.state.Areas.MainContent.Height - (CreateButtonHeight +
		2 + // divider height with padding
		2) // nolint:gomnd // list header
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
