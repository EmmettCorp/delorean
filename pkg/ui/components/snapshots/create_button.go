package snapshots

import (
	"time"

	"github.com/EmmettCorp/delorean/pkg/rate"
	"github.com/EmmettCorp/delorean/pkg/ui/shared"
	"github.com/charmbracelet/lipgloss"
)

const (
	createButtonHeight = 3 // top border + title + bottom border
	createLimitInSec   = 2
	buttonBorderColor  = "#AD58B4"
)

type createButton struct {
	shared.ClickableItem
	title string
	state *shared.State
	// Limiter for create button is needed to allow to finish create snapshot operation.
	// There is no real point in real life doing snapshots every second.
	// If allow user to call btrfs.CreateSnapshot several times a second it could cause a exec.Command call error.
	limiter  *rate.Limiter
	callback func() error
}

func newCreateButton(st *shared.State, title string, coords shared.Coords) *createButton {
	cb := createButton{
		state:   st,
		title:   title,
		limiter: rate.NewLimiter(time.Second * createLimitInSec),
	}

	cb.SetCoords(coords)

	return &cb
}

func (cb *createButton) SetTitle(title string) {
	cb.title = title
}

func (cb *createButton) OnClick() error {
	if !cb.limiter.Allow() {
		// TODO: consider to return errors.New("too many create calls per second")
		// and write it to to status bar

		return nil
	}

	return cb.callback()
}

func (cb *createButton) Render() string {
	return lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder(), true).
		Padding(0, 1).
		BorderForeground(lipgloss.AdaptiveColor{Dark: buttonBorderColor}).
		Render(cb.title)
}
