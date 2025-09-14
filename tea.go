package main

import tea "github.com/charmbracelet/bubbletea"

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:

		switch msg.String() {

		// Quitting
		case "ctrl+c":
			return m, tea.Quit

		// Movement by cell
		case "up", "w":
			m.state.CursorCellUp()

		case "down", "s":
			m.state.CursorCellDown()

		case "right", "d":
			m.state.CursorCellRight()

		case "left", "a":
			m.state.CursorCellLeft()

		// Movement by section
		case "ctrl+up", "ctrl+w", "shift+up", "shift+w":
			m.state.CursorSectionUp()

		case "ctrl+down", "ctrl+s", "shift+down", "shift+s":
			m.state.CursorSectionDown()

		case "ctrl+right", "ctrl+d", "shift+right", "shift+d":
			m.state.CursorSectionRight()

		case "ctrl+left", "ctrl+a", "shift+left", "shift+a":
			m.state.CursorSectionRight()

		// Marking/flagging
		case " ", "enter":
			m.state.Mark()

		case "x":
			m.state.Flag()

		case "delete", "backspace":
			m.state.Clear()
		}
	}

	return m, nil
}

func (m model) View() string {
	return ""
}
