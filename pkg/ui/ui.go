/*
Package ui implements the UI for the delorean application.
*/
package ui

import (
	"strings"

	"github.com/EmmettCorp/delorean/pkg/config"
	"github.com/EmmettCorp/delorean/pkg/ui/components/tabs"
	"github.com/EmmettCorp/delorean/pkg/ui/shared"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	tabs   tabs.Model
	keys   KeyMap
	state  *shared.State
	config *config.Config
}

func NewModel(cfg *config.Config) (*Model, error) {
	st := shared.State{}
	tabModel, err := tabs.NewModel(&st, getCurrentTabTitles())
	if err != nil {
		return &Model{}, err
	}

	// default current section id is 0
	return &Model{
		tabs:   tabModel,
		keys:   getKeyMaps(),
		state:  &st,
		config: cfg,
	}, nil
}

func (m *Model) Init() tea.Cmd {
	return tea.Batch(tea.EnterAltScreen)
}

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmd  tea.Cmd
		cmds []tea.Cmd
	)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		if key.Matches(msg, m.keys.Quit) {
			cmd = tea.Quit
		}
	case tea.MouseMsg:
		if msg.Type == tea.MouseLeft {
			m.onClick(msg)
		}
	}

	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

func (m *Model) View() string {
	s := strings.Builder{}
	s.WriteString(m.tabs.View())
	s.WriteString("\n")

	return s.String()
}

func (m *Model) onClick(event tea.MouseMsg) {
	if event.Y <= tabs.TabsHeigh {
		m.tabs.OnClick(event)
	}
}

func getCurrentTabTitles() []string {
	return []string{
		"Snapshots",
		// "Schedule",
		"Settings",
	}
}
