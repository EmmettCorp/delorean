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
			Width:  -1,
		},
		MainContent: uiArea{
			Width: -1,
		},
		HelpBar: uiArea{
			Height: HelpBarHeight,
			Width:  -1,
		},
	}

	return &ua
}
