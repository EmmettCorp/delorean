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
		status: fmt.Sprintf(" delorean version %s | type ctrl+h to call help", version.Number),
	}
}
