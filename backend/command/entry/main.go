package main

import (
	"backend/common/agentic"
	"backend/common/config"
	"backend/common/docker"
	"backend/tea"
	"os"

	"github.com/bsthun/gut"
	bubbletea "github.com/charmbracelet/bubbletea"

	"go.uber.org/fx"
)

func main() {
	fx.New(
		fx.Provide(
			config.Init,
			agentic.Init,
			docker.Init,
			tea.Init,
		),
		fx.Invoke(
			invoke,
		),
	).Run()
}

func invoke(t *tea.Model) {
	p := bubbletea.NewProgram(t)
	if _, err := p.Run(); err != nil {
		gut.Fatal("failed to run tea program", err)
	}
	os.Exit(0)
}
