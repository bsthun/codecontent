package tea

import (
	"os"
	"time"

	"github.com/bsthun/gut"
	"github.com/charmbracelet/bubbletea"
)

type Model struct {
	CurrentScreen string
}

func (r *Model) Init() tea.Cmd {
	return tea.Tick(3*time.Second, func(t time.Time) tea.Msg {
		r.CurrentScreen = "hello"
		return r.CurrentScreen
	})
}

func (r *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case string:
		r.CurrentScreen = msg
		return r, nil

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return r, tea.Quit
		}
	}

	return r, nil
}

func (r *Model) View() string {
	switch r.CurrentScreen {
	case "banner":
		bytes, err := os.ReadFile("./resource/banner.txt")
		if err != nil {
			gut.Fatal("failed to read banner file", err)
		}
		return string(bytes)
	case "hello":
		return "Hello from tea"
	default:
		return "Unknown"
	}
}
