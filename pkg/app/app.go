package app

type (
	App struct{}
)

func New() (*App, error) {
	return &App{}, nil
}

func (a *App) Run() {

}
