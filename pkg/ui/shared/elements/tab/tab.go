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
	shared.ClickableItem
	id    shared.TabItem
	state *shared.State
	title string
}

func New(state *shared.State, id shared.TabItem) (*Tab, error) {
	t := Tab{
		title: id.String(),
		state: state,
		id:    id,
	}
	err := t.state.AppendClickable(shared.TabHeader, &t)
	if err != nil {
		return nil, err
	}

	return &t, nil
}

func (t *Tab) GetID() shared.TabItem {
	return t.id
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
