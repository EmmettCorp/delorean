/*
Package snapshots keeps all the logic for snapshots component.
*/
package snapshots

import (
	"errors"
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/EmmettCorp/delorean/pkg/commands/btrfs"
	"github.com/EmmettCorp/delorean/pkg/domain"
	"github.com/EmmettCorp/delorean/pkg/ui/shared"
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

type (
	snapshot struct {
		Label       string
		VolumeLabel string
		Type        string
		VolumeID    string
		Kernel      string
		Path        string
	}

	snapshotRepo interface {
		Put(sn domain.Snapshot) error
		List(vIDs []string) ([]domain.Snapshot, error)
		Delete(path string) error
	}

	garbageRepo interface {
		Put(ph string) error
	}
)

func (s *snapshot) FilterValue() string { return s.Label }
func (s *snapshot) GetPath() string     { return s.Path }

type Model struct {
	state           *shared.State
	createBtn       *createButton
	list            list.Model
	keys            keyMap
	styles          list.DefaultItemStyles
	height          int
	currentPage     int
	itemsCount      int
	updateClickable bool
	err             error
	dialog          *dialog.Model
	snapshotRepo    snapshotRepo
	garbageRepo     garbageRepo
}

func New(st *shared.State, sr snapshotRepo, gr garbageRepo) (*Model, error) {
	m := Model{
		state:        st,
		currentPage:  -1,
		itemsCount:   -1,
		styles:       list.NewDefaultItemStyles(),
		keys:         getKeyMaps(),
		snapshotRepo: sr,
		garbageRepo:  gr,
	}

	itemD := itemDelegate{
		model: &m,
	}

	itemsModel := list.New([]list.Item{}, &itemD, 0, 0)
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
	})
	m.createBtn = createBtn
	m.createBtn.SetCallback(m.createSnapshots)
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
	s.WriteString(m.createBtn.Render())
	s.WriteString("\n")
	s.WriteString(lipgloss.NewStyle().Foreground(styles.DefaultTheme.InactiveText).Render(getSnapshotsHeader()))
	s.WriteString("\n")
	s.WriteString(divider.HorizontalLine(m.state.ScreenWidth, styles.DefaultTheme.InactiveText))
	s.WriteString("\n")
	m.list.SetSize(m.state.ScreenWidth, m.getHeight())
	s.WriteString(styles.MainDocStyle.Render(m.list.View()))
	s.WriteString("\n")
	s.WriteString(divider.HorizontalLine(m.state.ScreenWidth, styles.DefaultTheme.InactiveText))
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

	if len(m.list.Items()) != m.itemsCount {
		m.state.UpdateSnapshots = true
	}

	var cmd tea.Cmd
	// do not call btrfs commands for just ui update
	if m.state.UpdateSnapshots {
		m.UpdateList()
		m.itemsCount = len(m.list.Items())
		m.updateClickable = true
		m.state.UpdateSnapshots = false
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
	if _, ok := msg.(tea.WindowSizeMsg); ok {
		m.height = m.getHeight()
		m.updateClickable = true
		m, cmd := m.dialog.Update(tea.WindowSizeMsg{
			Width:  m.state.ScreenWidth,
			Height: m.state.ScreenHeight,
		})

		return m, cmd
	}

	mod, cmd := m.dialog.Update(msg)

	return mod, cmd
}

func (m *Model) UpdateList() {
	snaps, err := m.snapshotRepo.List(m.state.GetActiveVolumesIDs())
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
			Kernel:      snaps[i].Kernel,
		}
	}

	m.list.SetItems(items)
}

func (m *Model) selectByIndex(idx int) error {
	m.list.Select(idx)

	return nil
}

func (m *Model) getHeight() int {
	return m.state.Areas.MainContent.Height - listHeaderHeight - createButtonHeight - pageIndicatorHeight
}

func (m *Model) deleteSelectedKey() error {
	return m.deleteWithDialog(m.list.Index())
}

func (m *Model) deleteWithDialog(idx int) error {
	sn, err := m.getSnapshotByIndex(idx)
	if err != nil {
		return fmt.Errorf("can't get snapshot by index `%d`: %v", idx, err)
	}

	m.dialog = dialog.New(fmt.Sprintf("Remove snapshot %s?", sn.Label), "Ok", "Cancel", m.state.ScreenWidth,
		m.state.ScreenHeight,
		func() error {
			err := m.deleteByIndex(idx)
			if err != nil {
				return fmt.Errorf("can't delete by index `%d`: %v", idx, err)
			}
			m.dialog = nil
			m.updateClickable = true

			return nil
		}, func() error {
			m.list.Select(idx)
			m.dialog = nil
			m.updateClickable = true

			return nil
		})

	m.state.CleanClickable(shared.SnapshotsList)

	err = m.state.AppendClickable(shared.SnapshotsList, m.dialog.OkButton, m.dialog.CancelButton)
	if err != nil {
		return fmt.Errorf("can't append delete with dialog: %v", err)
	}

	return nil
}

func (m *Model) restoreWithDialog(idx int) error {
	sn, err := m.getSnapshotByIndex(idx)
	if err != nil {
		return fmt.Errorf("can't get snapshot by index `%d`: %v", idx, err)
	}

	m.dialog = dialog.New(fmt.Sprintf("Restore from snapshot %s?", sn.Label), "Ok", "Cancel", m.state.ScreenWidth,
		m.state.ScreenHeight,
		func() error {
			err := m.restoreByIndex(idx)
			if err != nil {
				return fmt.Errorf("can't restore by index `%d`: %v", idx, err)
			}
			m.dialog = nil
			m.updateClickable = true

			return nil
		}, func() error {
			m.list.Select(idx)
			m.dialog = nil
			m.updateClickable = true

			return nil
		})

	m.state.CleanClickable(shared.SnapshotsList)

	err = m.state.AppendClickable(shared.SnapshotsList, m.dialog.OkButton, m.dialog.CancelButton)
	if err != nil {
		return fmt.Errorf("can't append restore with dialog: %v", err)
	}

	return nil
}

func (m *Model) deleteByIndex(idx int) error {
	sn, err := m.getSnapshotByIndex(idx)
	if err != nil {
		return fmt.Errorf("can't get snapshot by index `%d`: %v", idx, err)
	}

	ph := sn.GetPath()
	err = btrfs.DeleteSnapshot(ph)
	if err != nil {
		return fmt.Errorf("can't delete from btrfs: %v", err)
	}

	err = m.snapshotRepo.Delete(ph)
	if err != nil {
		return fmt.Errorf("can't delete info from storage: %v", err)
	}

	m.list.RemoveItem(idx)

	return nil
}

func (m *Model) restoreByIndex(idx int) error {
	sn, err := m.getSnapshotByIndex(idx)
	if err != nil {
		return fmt.Errorf("can't get snapshot by index `%d`: %v", idx, err)
	}

	vol, err := m.getVolumeByID(sn.VolumeID)
	if err != nil {
		return err
	}

	if !m.state.Config.VolumeInRootFs(vol) {
		return fmt.Errorf("volume %s is not a child subvolume top level subvolume", vol.Label)
	}

	err = m.createSnapshotForVolume(vol, domain.Restore)
	if err != nil {
		return err
	}

	subvolumeDelorianMountPoint := path.Join(domain.DeloreanMountPoint, vol.Subvol)
	oldFsDelorianMountPoint := path.Join(domain.DeloreanMountPoint, fmt.Sprintf("%s.old", vol.Subvol))

	err = os.Rename(subvolumeDelorianMountPoint, oldFsDelorianMountPoint)
	if err != nil {
		return fmt.Errorf("can't rename directory %s: %v", oldFsDelorianMountPoint, err)
	}

	err = btrfs.Restore(sn.GetPath(), subvolumeDelorianMountPoint)
	if err != nil {
		return fmt.Errorf("can't create snapshot for %s: %v", vol.Device.MountPoint, err)
	}

	err = m.garbageRepo.Put(oldFsDelorianMountPoint)
	if err != nil {
		return fmt.Errorf("can't put old filesystem path to garbage storage %s: %v", oldFsDelorianMountPoint, err)
	}

	// force reboot logic

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

func (m *Model) createSnapshots() error {
	if !m.createBtn.Available() {
		// TODO: consider to return errors.New("too many create calls per second")
		// and write it to to status bar
		return nil
	}

	var activeVolumeFound bool

	for _, vol := range m.state.Config.Volumes {
		if !vol.Active {
			continue
		}

		activeVolumeFound = true

		err := m.createSnapshotForVolume(vol, domain.Manual)
		if err != nil {
			return err
		}
	}

	if !activeVolumeFound {
		// TODO: after creation write message to the status bar
		// put the message to a status bar errors.New("there are no active volumes")

		return nil
	}

	m.UpdateList()

	return nil
}

func (m *Model) createSnapshotForVolume(vol domain.Volume, snapType string) error {
	snap := domain.NewSnapshot(vol.SnapshotsPath, snapType, vol.Label, vol.ID, m.state.Config.KernelVersion)

	err := btrfs.CreateSnapshot(vol.Device.MountPoint, snap)
	if err != nil {
		return fmt.Errorf("can't create snapshot for %s: %v", snap.Path, err)
	}

	err = m.snapshotRepo.Put(snap)
	if err != nil {
		return fmt.Errorf("can't put snapshot to storage with path %s: %v", snap.Path, err)
	}

	return nil
}

func (m *Model) getVolumeByID(id string) (domain.Volume, error) {
	for i := range m.state.Config.Volumes {
		if m.state.Config.Volumes[i].ID == id {
			return m.state.Config.Volumes[i], nil
		}
	}

	return domain.Volume{}, fmt.Errorf("can't find volume by id `%s`", id)
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
