package main

import tea "github.com/charmbracelet/bubbletea"

type Input interface {
	Value() string
	Blur() tea.Msg
	Update(tea.Msg) (Input, tea.Cmd)
	View() string
}
