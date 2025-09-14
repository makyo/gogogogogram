package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	p := tea.NewProgram(newModel(4, 4))
	if _, err := p.Run(); err != nil {
		fmt.Printf("Ah drat — %v\n", err)
		os.Exit(1)
	}
}
