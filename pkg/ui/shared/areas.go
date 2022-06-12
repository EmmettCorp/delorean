package shared

const (
	TabsBarHeight = 3
	HelpBarHeight = 2
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
