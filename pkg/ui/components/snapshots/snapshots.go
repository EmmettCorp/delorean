/*
Package snapshots keeps all the logic for snapshots component.
*/
package snapshots

import (
	"fmt"
	"os"
	"strings"

	"github.com/EmmettCorp/delorean/pkg/commands/btrfs"
	"github.com/EmmettCorp/delorean/pkg/ui/components/button"
	"github.com/EmmettCorp/delorean/pkg/ui/components/divider"
	"github.com/EmmettCorp/delorean/pkg/ui/components/tabs"
	"github.com/EmmettCorp/delorean/pkg/ui/shared"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"golang.org/x/term"
)

type snapshot struct {
	Label       string
	VolumeLabel string
	Type        string
}

func (s snapshot) Title() string { return s.Label }
func (s snapshot) Description() string {
	return fmt.Sprintf("type: %s | volume: %s ", s.Type, s.VolumeLabel)
}
func (s snapshot) FilterValue() string { return s.Label }

type Model struct {
	state     *shared.State
	createBtn button.Model
	list      list.Model
	err       error
}

func NewModel(st *shared.State) (*Model, error) {
	itemsModel := list.New([]list.Item{}, list.NewDefaultDelegate(), 0, 0)
	itemsModel.SetFilteringEnabled(false)
	itemsModel.SetShowFilter(false)
	itemsModel.SetShowTitle(false)
	itemsModel.SetShowStatusBar(false)
	itemsModel.SetShowHelp(false)

	createBtn := newCreateButton("Create", shared.Coords{
		X1: 0,
		Y1: 5,
		X2: 3,
		Y2: 7,
	})

	return &Model{
		list:      itemsModel,
		state:     st,
		createBtn: createBtn,
	}, nil
}

func (m *Model) Init() tea.Cmd {
	return nil
}

func (m *Model) View() string {
	w, h, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		return err.Error()
	}

	s := strings.Builder{}
	s.WriteString(button.DrawButton(m.createBtn.GetTitle()))
	s.WriteString("\n")
	s.WriteString(lipgloss.NewStyle().SetString("  Info\t\t\t\t\tID\t\tKernel").Foreground(inactive).String())
	s.WriteString("\n")
	s.WriteString(divider.Horizontal(w, subtle))
	s.WriteString("\n")
	m.list.SetSize(w, h-((tabs.TabsHeigh+1)+2+2+2)) // nolint:gomnd // (TabsHeigh + bottom line) + Header + Divider
	s.WriteString(docStyle.Render(m.list.View()))

	return s.String()
}

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if msg, ok := msg.(tea.WindowSizeMsg); ok {
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
		})
	}

	m.list.SetItems(items)
}
