package ui

import (
	"fmt"

	"github.com/EmmettCorp/delorean/pkg/ui/components/tabs"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

var tabsTitles = []string{
	"Snapshots",
	"Schedule",
	"Settings",
}

type Model struct {
	currSectionId int
	tabs          tabs.Model
	keys          KeyMap
	err           error
}

func NewModel() (Model, error) {
	tabModel, err := tabs.NewModel(tabsTitles)
	if err != nil {
		return Model{}, err
	}

	// default current section id is 0
	return Model{
		tabs: tabModel,
	}, nil
}

func (m Model) Init() tea.Cmd {
	return tea.Batch(tea.EnterAltScreen)
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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
	}

	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
	// return m, tea.Batch(tea.ClearScreen, tea.Draw)
}

func (m Model) View() string {
	return fmt.Sprintf("%v", m.err)
}

func Draw() {
	tabModel, err := tabs.NewModel(tabsTitles)
	if err != nil {
		return
	}

	fmt.Println(tabModel.View())
}
