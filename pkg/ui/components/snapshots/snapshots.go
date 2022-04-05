package snapshots

import (
	"fmt"
	"os"

	"github.com/EmmettCorp/delorean/pkg/domain"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
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
	list list.Model
}

func NewModel() (Model, error) {
	itemsModel := list.New([]list.Item{}, list.NewDefaultDelegate(), 0, 0)
	itemsModel.SetFilteringEnabled(false)
	itemsModel.SetShowFilter(false)
	itemsModel.SetShowTitle(false)

	return Model{
		list: itemsModel,
	}, nil
}

func (m Model) View() string {
	w, h, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		return err.Error()
	}
	m.list.SetSize(w, h-6)
	return docStyle.Render(m.list.View())
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		h, v := docStyle.GetFrameSize()
		m.list.SetSize(msg.Width-h, msg.Height-v)
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m *Model) UpdateList(snps []domain.Snapshot) {
	items := []list.Item{}
	for i := range snps {
		items = append(items, snapshot{
			Label:       snps[i].Label,
			VolumeLabel: snps[i].VolumeLabel,
			Type:        snps[i].Type,
		})
	}

	m.list.SetItems(items)
}
