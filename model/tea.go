package model

import tea "github.com/charmbracelet/bubbletea"

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:

		switch msg.String() {

		// Quitting
		case "ctrl+c":
			return m, tea.Quit

		// Movement by cell
		case "up", "w":
			m = m.CursorCellUp()

		case "down", "s":
			m = m.CursorCellDown()

		case "right", "d":
			m = m.CursorCellRight()

		case "left", "a":
			m = m.CursorCellLeft()

		// Movement by section
		case "ctrl+up", "ctrl+w", "shift+up", "shift+w":
			m = m.CursorSectionUp()

		case "ctrl+down", "ctrl+s", "shift+down", "shift+s":
			m = m.CursorSectionDown()

		case "ctrl+right", "ctrl+d", "shift+right", "shift+d":
			m = m.CursorSectionRight()

		case "ctrl+left", "ctrl+a", "shift+left", "shift+a":
			m = m.CursorSectionRight()

		// Marking/flagging
		case " ", "enter":
			m = m.Mark()

		case "x":
			m = m.Flag()

		case "delete", "backspace":
			m = m.ClearGuess()
		}
	}

	return m, nil
}

func (m Model) View() string {
	return ""
}
