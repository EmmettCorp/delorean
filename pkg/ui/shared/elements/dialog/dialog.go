package dialog

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const (
	dHeight = 5
	dWidth  = 50
)

type Model struct {
	Ok         bool
	Text       string
	OkText     string
	CancelText string
	OkFunc     func()
	CancelFunc func()
	w          int
	h          int
	keys       keyMap
}

func New(text, okText, cancelText string, w, h int, okFunc, cancelFunc func()) *Model {
	return &Model{
		Ok:         true,
		Text:       text,
		OkText:     okText,
		CancelText: cancelText,
		w:          w,
		h:          h,
		OkFunc:     okFunc,
		CancelFunc: cancelFunc,
		keys:       getKeyMaps(),
	}
}

func (m *Model) Init() tea.Cmd {
	return nil
}

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		// m.height = m.getHeight()
		// m.list.SetSize(m.state.ScreenWidth, m.height)
		// m.updateClickable = true
	// case tea.MouseMsg:
	// 	if msg.Type == tea.MouseWheelDown {
	// 		m.list.Paginator.NextPage()
	// 	} else if msg.Type == tea.MouseWheelUp {
	// 		m.list.Paginator.PrevPage()
	// 	}
	case tea.KeyMsg:
		if key.Matches(msg, m.keys.Left) {
			m.Ok = true
		} else if key.Matches(msg, m.keys.Right) {
			m.Ok = false
		} else if key.Matches(msg, m.keys.Enter) {
			m.callback()
		}
	}

	var cmd tea.Cmd

	return m, cmd
}

func (m *Model) View() string {
	okButton := activeButtonStyle.Render(m.OkText)
	cancelButton := buttonStyle.Render(m.CancelText)
	if !m.Ok {
		okButton = buttonStyle.Render(m.OkText)
		cancelButton = activeButtonStyle.Render(m.CancelText)
	}

	question := lipgloss.NewStyle().Width(dWidth).Align(lipgloss.Center).Render(m.Text)
	buttons := lipgloss.JoinHorizontal(lipgloss.Top, okButton, cancelButton)
	ui := lipgloss.JoinVertical(lipgloss.Center, question, buttons)

	return lipgloss.Place(m.w, m.h+dHeight,
		lipgloss.Center, lipgloss.Center,
		dialogBoxStyle.Render(ui),
		lipgloss.WithWhitespaceChars(" "),
		lipgloss.WithWhitespaceForeground(subtle),
	)
}

func (m *Model) callback() {
	if m.Ok {
		m.OkFunc()
		return
	}

	m.CancelFunc()
}
