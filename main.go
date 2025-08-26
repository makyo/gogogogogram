package main

import (
	"fmt"
	"os"

	"git.makyo.dev/makyo/gogogogogram/model"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	p := tea.NewProgram(model.New(4, 4))
	if _, err := p.Run(); err != nil {
		fmt.Printf("Ah drat — %v\n", err)
		os.Exit(1)
	}
}
