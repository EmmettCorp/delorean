package gui

import (
	"fmt"

	"github.com/EmmettCorp/delorean/pkg/version"
)

type state struct {
	status string
}

func initState() *state {
	return &state{
		status: fmt.Sprintf(" application is running. version %s ", version.Number),
	}
}
