package shared

const (
	TabsBarHeight = 2
	HelpBarHeight = 3
	FullScreen    = -1
)

type uiArea struct {
	Height int
	Width  int
	Coords Coords
}

type uiAreas struct {
	TabBar      uiArea
	MainContent uiArea
	HelpBar     uiArea
}

func initAreas() *uiAreas {
	ua := uiAreas{
		TabBar: uiArea{
			Height: TabsBarHeight,
			Width:  FullScreen,
		},
		MainContent: uiArea{
			Width: FullScreen,
		},
		HelpBar: uiArea{
			Height: HelpBarHeight,
			Width:  FullScreen,
		},
	}

	return &ua
}
