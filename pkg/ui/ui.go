package ui

import (
	"fmt"
	"strings"

	"github.com/EmmettCorp/delorean/pkg/config"
	"github.com/EmmettCorp/delorean/pkg/ui/components/tabs"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

var tabsTitles = []string{
	"Snapshots",
	// "Schedule",
	"Settings",
}

type Model struct {
	currentTab int
	tabs       tabs.Model
	keys       KeyMap
	config     *config.Config
	mouseEvent tea.MouseEvent
	err        error
}

func NewModel(cfg *config.Config) (*Model, error) {
	tabModel, err := tabs.NewModel(tabsTitles)
	if err != nil {
		return &Model{}, err
	}

	// default current section id is 0
	return &Model{
		tabs:   tabModel,
		keys:   Keys,
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
		switch {
		case key.Matches(msg, m.keys.Quit):
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

func Draw() {
	tabModel, err := tabs.NewModel(tabsTitles)
	if err != nil {
		return
	}

	fmt.Println(tabModel.View())
}

func (m *Model) onClick(event tea.MouseMsg) {
	if event.Y < 3 {
		m.currentTab = m.tabs.OnClick(event)
		m.tabs.SetcurrentTabID(m.currentTab)
	}
}
