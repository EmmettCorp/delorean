package ui

import (
	"strings"

	"github.com/EmmettCorp/delorean/pkg/config"
	"github.com/EmmettCorp/delorean/pkg/ui/components/help"
	"github.com/EmmettCorp/delorean/pkg/ui/components/tabs"
	"github.com/EmmettCorp/delorean/pkg/ui/context"
	"github.com/EmmettCorp/delorean/pkg/utils"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
	keys          utils.KeyMap
	err           error
	currSectionId int
	help          help.Model
	ready         bool
	isSidebarOpen bool
	tabs          tabs.Model
	ctx           context.ProgramContext
}

func NewModel() Model {
	tabsModel := tabs.NewModel()
	return Model{
		keys:          utils.Keys,
		help:          help.NewModel(),
		currSectionId: 0,
		tabs:          tabsModel,
	}
}

func initScreen() tea.Msg {
	config, err := config.New()
	if err != nil {
		return errMsg{err}
	}

	return initMsg{Config: *config}
}

func (m Model) Init() tea.Cmd {
	return tea.Batch(initScreen, tea.EnterAltScreen)
}

func (m Model) View() string {
	if m.err != nil {
		return m.err.Error()
	}

	// if m.ctx.Config == nil {
	// 	return "Reading config...\n"
	// }

	s := strings.Builder{}
	s.WriteString(m.tabs.View(m.ctx))
	s.WriteString("\n")
	mainContent := lipgloss.JoinHorizontal(
		lipgloss.Top,
		m.getCurrSection().View(),
		// m.sidebar.View(),
	)
	s.WriteString(mainContent)
	s.WriteString("\n")
	s.WriteString(m.help.View(m.ctx))
	return s.String()
}

type initMsg struct {
	Config config.Config
}

type errMsg struct {
	error
}

func (e errMsg) Error() string { return e.error.Error() }
