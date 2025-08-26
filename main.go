package main

import (
	"fmt"
	"os"

	"git.makyo.dev/makyo/gogogogogram/model"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	m, err := model.New(4, 4)
	if err != nil {
		fmt.Printf("Ah, drat — %v", err)
		os.Exit(1)
	}
	p := tea.NewProgram(m)
	if _, err := p.Run(); err != nil {
		fmt.Printf("Ah drat. %v", err)
		os.Exit(1)
	}
}
