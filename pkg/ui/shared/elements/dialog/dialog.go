package dialog

import (
	"github.com/EmmettCorp/delorean/pkg/ui/shared"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const (
	dHeight = 5
	dWidth  = 50
)

// Model is a dialog model.
type Model struct {
	Title        string
	OkButton     *Button
	CancelButton *Button
	w            int
	h            int
	keys         keyMap
}

// New creates and returns a new dialog model.
func New(title, okText, cancelText string, w, h int, okFunc, cancelFunc func()) *Model {
	m := Model{
		Title: title,
		OkButton: &Button{
			Text:     okText,
			Callback: okFunc,
			active:   true,
		},
		CancelButton: &Button{
			Text:     cancelText,
			Callback: cancelFunc,
			active:   false,
		},
		w:    w,
		h:    h + dHeight,
		keys: getKeyMaps(),
	}
	m.setButtonsCoords()

	return &m
}

func (m *Model) Init() tea.Cmd {
	return nil
}

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.h = msg.Height + dHeight
		m.w = msg.Width
		m.setButtonsCoords()
	case tea.KeyMsg:
		if key.Matches(msg, m.keys.Left) {
			m.confirm(true)
		} else if key.Matches(msg, m.keys.Right) {
			m.confirm(false)
		} else if key.Matches(msg, m.keys.Enter) {
			m.getActiveButton().Callback()
		}
	}

	var cmd tea.Cmd

	return m, cmd
}

func (m *Model) View() string {
	question := lipgloss.NewStyle().Width(dWidth).Align(lipgloss.Center).Render(m.Title)
	buttons := lipgloss.JoinHorizontal(lipgloss.Top, m.OkButton.Render(), m.CancelButton.Render())
	ui := lipgloss.JoinVertical(lipgloss.Center, question, buttons)

	return lipgloss.Place(m.w, m.h,
		lipgloss.Center, lipgloss.Center,
		dialogBoxStyle.Render(ui),
		lipgloss.WithWhitespaceChars(" "),
		lipgloss.WithWhitespaceForeground(subtle),
	)
}

func (m *Model) confirm(ok bool) {
	m.OkButton.active = ok
	m.CancelButton.active = !ok
}

func (m *Model) getActiveButton() *Button {
	if m.OkButton.active {
		return m.OkButton
	}

	return m.CancelButton
}

func (m *Model) setButtonsCoords() {
	center := (m.w / 2) - 1
	y1 := m.h/2 + dHeight - 1

	m.OkButton.SetCoords(shared.Coords{
		X1: center - 10,
		Y1: y1,
		X2: center - 3,
		Y2: y1 + 1,
	})
	m.CancelButton.SetCoords(shared.Coords{
		X1: center,
		Y1: y1,
		X2: center + 11,
		Y2: y1 + 1,
	})
}
