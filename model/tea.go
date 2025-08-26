package model

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
			m.cursorCellUp()

		case "down", "s":
			m.cursorCellDown()

		case "right", "d":
			m.cursorCellRight()

		case "left", "a":
			m.cursorCellLeft()

		// Movement by section
		case "ctrl+up", "ctrl+w", "shift+up", "shift+w":
			m.cursorSectionUp()

		case "ctrl+down", "ctrl+s", "shift+down", "shift+s":
			m.cursorSectionDown()

		case "ctrl+right", "ctrl+d", "shift+right", "shift+d":
			m.cursorSectionRight()

		case "ctrl+left", "ctrl+a", "shift+left", "shift+a":
			m.cursorSectionRight()

		// Marking/flagging
		case " ", "enter":
			m.state.mark(*m.cursor)

		case "x":
			m.state.flag(*m.cursor)

		case "delete", "backspace":
			m.state.clear(*m.cursor)
		}
	}

	return m, nil
}

func (m model) View() string {
	return ""
}
