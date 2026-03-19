package ui

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

		// Saving
		case "ctrl+s":
			return m, nil

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
		case "shift+up", "W":
			m.state.CursorSectionUp()

		case "shift+down", "S":
			m.state.CursorSectionDown()

		case "shift+right", "D":
			m.state.CursorSectionRight()

		case "shift+left", "A":
			m.state.CursorSectionLeft()

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
	return m.state.View()
}
