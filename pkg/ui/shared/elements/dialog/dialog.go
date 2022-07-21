/*
Package dialog keeps helpers to create a standard dialog window.
*/
package dialog

import (
	"github.com/EmmettCorp/delorean/pkg/ui/shared"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const (
	dHeight = 6
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
func New(title, okText, cancelText string, w, h int, okFunc, cancelFunc func() error) *Model {
	m := Model{
		Title: title,
		OkButton: &Button{
			Text:   okText,
			active: true,
		},
		CancelButton: &Button{
			Text:   cancelText,
			active: false,
		},
		w:    w,
		h:    h - dHeight,
		keys: getKeyMaps(),
	}
	m.OkButton.SetCallback(okFunc)
	m.CancelButton.SetCallback(cancelFunc)
	m.setButtonsCoords()

	return &m
}

func (m *Model) Init() tea.Cmd {
	return nil
}

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.h = msg.Height - dHeight
		m.w = msg.Width
		m.setButtonsCoords()
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keys.Left):
			m.confirm(true)
		case key.Matches(msg, m.keys.Right):
			m.confirm(false)
		case key.Matches(msg, m.keys.Enter):
			m.getActiveButton().OnClick()
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

// nolint:gomnd // we divide all `w` and `h` to 2 in order to get center.
// minus - 1 every time to include first row position of element
// plus  + 1 every time to set button height.
func (m *Model) setButtonsCoords() {
	center := (m.w / 2) - 1
	y1 := (m.h + dHeight) / 2

	okLen := lipgloss.Width(m.OkButton.Text) + buttonPadding*2
	cancelLen := lipgloss.Width(m.CancelButton.Text) + buttonPadding*2

	okCoords := shared.Coords{
		X1: center - okLen - buttonMargin*2,
		Y1: y1,
		Y2: y1 + lipgloss.Height(m.OkButton.Text),
	}
	okCoords.X2 = okCoords.X1 + okLen - 1
	m.OkButton.SetCoords(okCoords)

	cancelCoords := shared.Coords{
		X1: center,
		Y1: y1,
		Y2: y1 + lipgloss.Height(m.OkButton.Text),
	}
	cancelCoords.X2 = cancelCoords.X1 + cancelLen - 1
	m.CancelButton.SetCoords(cancelCoords)
}
