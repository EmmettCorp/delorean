package ui

import (
	"fmt"

	"github.com/EmmettCorp/delorean/pkg/ui/components/tabs"
)

var tabsTitles = []string{
	"Snapshots",
	"Schedule",
	"Settings",
}

func Draw() {
	tabModel, err := tabs.NewModel(tabsTitles)
	if err != nil {
		return
	}

	fmt.Println(tabModel.View())
}
