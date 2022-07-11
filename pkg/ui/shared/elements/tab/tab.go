/*
Package tab keeps helpers to create tabs.
*/
package tab

import (
	"strings"

	"github.com/EmmettCorp/delorean/pkg/ui/shared"
	"github.com/charmbracelet/lipgloss"
)

const (
	tabsLeftRightIndents = 2
)

type Tab struct {
	id       shared.TabItem
	state    *shared.State
	coords   shared.Coords
	title    string
	callback func() error
}

func New(state *shared.State, id shared.TabItem, coords shared.Coords) (*Tab, error) {
	t := Tab{
		title: id.String(),
		state: state,
		id:    id,
	}
	t.SetCoords(coords)
	err := t.state.AppendClickable(shared.TabHeader, &t)
	if err != nil {
		return nil, err
	}

	return &t, nil
}

func (t *Tab) GetID() shared.TabItem {
	return t.id
}

func (t *Tab) GetCoords() shared.Coords {
	return t.coords
}

func (t *Tab) SetCoords(c shared.Coords) {
	t.coords = c
}

func (t *Tab) SetCallback(callback func() error) {
	t.callback = callback
}

func (t *Tab) OnClick() error {
	t.state.Update(t.id)

	return nil
}

func (t *Tab) Render() string {
	if t.id == t.state.CurrentTab {
		return activeTab.Render(t.title)
	}

	return inactiveTab.Render(t.title)
}

func RenderTabBar(screenWidth int, tabs []string) string {
	row := lipgloss.JoinHorizontal(
		lipgloss.Top,
		tabs...,
	)

	gap := tabGap.Render(strings.Repeat(" ", shared.Max(0, screenWidth-lipgloss.Width(row)-tabsLeftRightIndents)))

	return docStyle.Render(lipgloss.JoinHorizontal(lipgloss.Bottom, row, gap))
}
