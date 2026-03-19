package main

import (
	"fmt"
	"os"

	"git.makyo.dev/makyo/gogogogogram/ui"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	p := tea.NewProgram(ui.NewModel(4, 4))
	if _, err := p.Run(); err != nil {
		fmt.Printf("Ah drat — %v\n", err)
		os.Exit(1)
	}
}
