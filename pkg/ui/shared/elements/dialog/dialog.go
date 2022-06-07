package dialog

import (
	"github.com/EmmettCorp/delorean/pkg/logger"
	"github.com/EmmettCorp/delorean/pkg/ui/shared"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const (
	dHeight = 5
	dWidth  = 50
)

type Button struct {
	Text     string
	Callback func()
	active   bool
	coords   shared.Coords
}

type Model struct {
	Title        string
	OkButton     Button
	CancelButton Button
	w            int
	h            int
	keys         keyMap
}

func New(title, okText, cancelText string, w, h int, okFunc, cancelFunc func()) *Model {
	return &Model{
		Title: title,
		OkButton: Button{
			Text:     okText,
			Callback: okFunc,
			active:   true,
		},
		CancelButton: Button{
			Text:     cancelText,
			Callback: cancelFunc,
			active:   false,
		},
		w:    w,
		h:    h + dHeight,
		keys: getKeyMaps(),
	}
}

func (m *Model) Init() tea.Cmd {
	return nil
}

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.h = msg.Height + dHeight
		m.w = msg.Width
		logger.Client.InfoLog.Printf("resize height: %d", msg.Height)
	// case tea.MouseMsg:
	// 	if msg.Type == tea.MouseWheelDown {
	// 		m.list.Paginator.NextPage()
	// 	} else if msg.Type == tea.MouseWheelUp {
	// 		m.list.Paginator.PrevPage()
	// 	}
	case tea.KeyMsg:
		if key.Matches(msg, m.keys.Left) {
			m.Ok(true)
		} else if key.Matches(msg, m.keys.Right) {
			m.Ok(false)
		} else if key.Matches(msg, m.keys.Enter) {
			m.getActiveButton().Callback()
		}
	}

	var cmd tea.Cmd

	return m, cmd
}

func (m *Model) View() string {
	okButton := m.OkButton.Render()
	cancelButton := m.CancelButton.Render()

	question := lipgloss.NewStyle().Width(dWidth).Align(lipgloss.Center).Render(m.Title)
	buttons := lipgloss.JoinHorizontal(lipgloss.Top, okButton, cancelButton)
	ui := lipgloss.JoinVertical(lipgloss.Center, question, buttons)

	return lipgloss.Place(m.w, m.h,
		lipgloss.Center, lipgloss.Center,
		dialogBoxStyle.Render(ui),
		lipgloss.WithWhitespaceChars(" "),
		lipgloss.WithWhitespaceForeground(subtle),
	)
}

func (m *Model) Ok(ok bool) {
	m.OkButton.active = ok
	m.CancelButton.active = !ok
}

func (m *Model) getActiveButton() Button {
	if m.OkButton.active {
		return m.OkButton
	}

	return m.CancelButton
}

func (b *Button) Render() string {
	var style lipgloss.Style

	if b.active {
		style = activeButtonStyle
	} else {
		style = buttonStyle
	}

	return style.Render(b.Text)
}
