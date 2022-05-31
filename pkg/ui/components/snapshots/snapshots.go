/*
Package snapshots keeps all the logic for snapshots component.
*/
package snapshots

import (
	"errors"
	"fmt"
	"strings"

	"github.com/EmmettCorp/delorean/pkg/commands/btrfs"
	"github.com/EmmettCorp/delorean/pkg/ui/shared"
	"github.com/EmmettCorp/delorean/pkg/ui/shared/elements/button"
	"github.com/EmmettCorp/delorean/pkg/ui/shared/elements/dialog"
	"github.com/EmmettCorp/delorean/pkg/ui/shared/elements/divider"
	"github.com/EmmettCorp/delorean/pkg/ui/shared/styles"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const (
	infoTitle       = "Info"
	idTitle         = "ID"
	typeTitle       = "Type"
	infoColumnWidth = 30
	idColumnWidth   = 6
	typeColumnWidth = 10

	listHeaderHeight    = 2
	pageIndicatorHeight = 1

	minColumnGap     = "  "
	minColumnGapLen  = len(minColumnGap)
	infoColumnsWidth = infoColumnWidth + idColumnWidth + typeColumnWidth + minColumnGapLen
)

type buttonModel interface {
	shared.Clickable
	SetTitle(title string)
	GetTitle() string
}

type snapshot struct {
	Label       string
	VolumeLabel string
	Type        string
	VolumeID    string
	Kernel      string
	Path        string
}

func (s *snapshot) FilterValue() string { return s.Label }
func (s *snapshot) GetPath() string     { return s.Path }

type Model struct {
	state           *shared.State
	createBtn       buttonModel
	list            list.Model
	keys            keyMap
	styles          list.DefaultItemStyles
	height          int
	currentPage     int
	itemsCount      int
	updateClickable bool
	err             error
	dialog          tea.Model
}

func NewModel(st *shared.State) (*Model, error) {
	m := Model{
		state:       st,
		currentPage: -1,
		itemsCount:  -1,
		styles:      list.NewDefaultItemStyles(),
		keys:        getKeyMaps(),
	}

	itemsModel := list.New([]list.Item{}, &itemDelegate{
		model: &m,
	}, 0, 0)
	itemsModel.SetFilteringEnabled(false)
	itemsModel.SetShowFilter(false)
	itemsModel.SetShowTitle(false)
	itemsModel.SetShowStatusBar(false)
	itemsModel.SetShowHelp(false)
	m.list = itemsModel
	m.UpdateList()

	btnTitle := "Create"
	createButtonY1 := st.Areas.TabBar.Height + 1
	createBtn := newCreateButton(st, btnTitle, shared.Coords{
		Y1: createButtonY1,
		X2: lipgloss.Width(btnTitle) + 3,            // nolint:gomnd // left and right borders + 1
		Y2: createButtonY1 + createButtonHeight - 1, // we don't need make bottom border line clickable
	}, m.UpdateList)
	m.createBtn = createBtn
	err := st.AppendClickable(shared.SnapshotsButtonsBar, createBtn)
	if err != nil {
		return nil, err
	}

	return &m, nil
}

func (m *Model) Init() tea.Cmd {
	return nil
}

func (m *Model) View() string {
	if m.dialog != nil {
		return m.dialog.View()
	}

	var s strings.Builder
	s.WriteString(button.New(m.createBtn.GetTitle()))
	s.WriteString("\n")
	s.WriteString(lipgloss.NewStyle().SetString(getSnapshotsHeader()).
		Foreground(styles.DefaultTheme.InactiveText).String())
	s.WriteString("\n")
	s.WriteString(divider.HorizontalLine(m.state.ScreenWidth, styles.DefaultTheme.InactiveText))
	s.WriteString("\n")
	m.list.SetSize(m.state.ScreenWidth, m.height)
	s.WriteString(styles.MainDocStyle.Render(m.list.View()))
	// set updateClickable = false after list page rendering only
	// otherwise there can be not set clickable elements
	m.updateClickable = false

	return s.String()
}

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if m.dialog != nil {
		return m.updateDialog(msg)
	}

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.height = m.getHeight()
		m.list.SetSize(m.state.ScreenWidth, m.height)
		m.updateClickable = true
	case tea.MouseMsg:
		if msg.Type == tea.MouseWheelDown {
			m.list.Paginator.NextPage()
		} else if msg.Type == tea.MouseWheelUp {
			m.list.Paginator.PrevPage()
		}
	case tea.KeyMsg:
		if key.Matches(msg, m.keys.Delete) {
			m.err = m.deleteSelectedKey()
		}
	}

	var cmd tea.Cmd
	// do not call btrfs commands for just ui update
	if len(m.list.Items()) != m.itemsCount {
		m.UpdateList()
		m.itemsCount = len(m.list.Items())
		m.updateClickable = true
	}
	if m.currentPage != m.list.Paginator.Page {
		m.currentPage = m.list.Paginator.Page
		m.updateClickable = true
	}

	if m.updateClickable {
		m.state.CleanClickable(shared.SnapshotsList)
	}
	m.list, cmd = m.list.Update(msg)

	return m, cmd
}

func (m *Model) updateDialog(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg.(type) {
	case tea.WindowSizeMsg:
		m.height = m.getHeight()
		m.updateClickable = true
		m, cmd := m.dialog.Update(tea.WindowSizeMsg{
			Width:  m.state.ScreenWidth,
			Height: m.height,
		})

		return m, cmd
	}

	mod, cmd := m.dialog.Update(msg)

	return mod, cmd
}

func (m *Model) UpdateList() {
	snaps, err := btrfs.SnapshotsList(m.state.Config.Volumes)
	if err != nil {
		m.err = err

		return
	}

	items := make([]list.Item, len(snaps))
	for i := range snaps {
		items[i] = &snapshot{
			Label:       snaps[i].Label,
			VolumeLabel: snaps[i].VolumeLabel,
			Type:        snaps[i].Type,
			VolumeID:    snaps[i].VolumeID,
			Path:        snaps[i].Path,
		}
	}

	m.list.SetItems(items)
}

func (m *Model) getHeight() int {
	return m.state.Areas.MainContent.Height - createButtonHeight - listHeaderHeight - pageIndicatorHeight
}

func (m *Model) deleteSelectedKey() error {
	return m.deleteWithDialog(m.list.Index())
}

func (m *Model) deleteWithDialog(idx int) error {
	sn, err := m.getSnapshotByIndex(idx)
	if err != nil {
		return fmt.Errorf("can't get snapshot by index `%d`: %v", idx, err)
	}

	m.dialog = dialog.New(fmt.Sprintf("Remove snapshot %s?", sn.Label), "Ok", "Cancel", m.state.ScreenWidth, m.height,
		func() {
			m.deleteByIndex(idx)
			m.dialog = nil
		}, func() {
			m.list.Select(idx)
			m.dialog = nil
		})

	return nil
}

func (m *Model) deleteByIndex(idx int) error {
	sn, err := m.getSnapshotByIndex(idx)
	if err != nil {
		return fmt.Errorf("can't get snapshot by index `%d`: %v", idx, err)
	}

	err = btrfs.DeleteSnapshot(sn.GetPath())
	if err != nil {
		return err
	}

	m.list.RemoveItem(idx)

	return nil
}

func (m *Model) getSnapshotByIndex(idx int) (*snapshot, error) {
	items := m.list.Items()
	if idx >= len(items) {
		return nil, fmt.Errorf("index `%d` is out of range", idx)
	}

	sn, ok := items[idx].(*snapshot)
	if !ok {
		return nil, errors.New("can't assert item to snapshot type")
	}

	return sn, nil
}

func getSnapshotsHeader() string {
	var header strings.Builder
	header.WriteString(minColumnGap)
	header.WriteString(infoTitle)
	header.WriteString(strings.Repeat(" ", infoColumnWidth-lipgloss.Width(infoTitle)-minColumnGapLen))
	header.WriteString(idTitle)
	header.WriteString(strings.Repeat(" ", idColumnWidth-lipgloss.Width(idTitle)))
	header.WriteString(typeTitle)

	return header.String()
}
