package main

import (
	"fmt"
	"log"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type model struct {
	questions []Question
	index     int
	width     int
	height    int
	styles    *Styles
	done      bool
	viewport  viewport.Model
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	current := &m.questions[m.index]
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "ctrl+n":
			current.answer = current.input.Value()
			if m.index == len(m.questions)-1 {
				m.GlamourRender(m.BuildResponseMarkdown())
				m.done = true
			}
			log.Printf("question: %s, answer: %s", current.question, current.answer)
			m.Next()
			return m, current.input.Blur
		case "ctrl+p":
			if m.index == 0 {
				return m, nil
			}
			current.answer = current.input.Value()
			log.Printf("question: %s, answer: %s", current.question, current.answer)
			m.Prev()
			return m, current.input.Blur
		}

		if m.done {
			m.viewport, cmd = m.viewport.Update(msg)
			return m, cmd
		}
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
	}

	current.input, cmd = current.input.Update(msg)
	return m, cmd
}

func (m model) View() string {
	current := m.questions[m.index]
	if m.done {
		var output string
		for _, q := range m.questions {
			output += fmt.Sprintf("%s: %s\n", q.question, q.answer)
		}
		return lipgloss.Place(
			m.width,
			m.height,
			lipgloss.Center,
			lipgloss.Center,
			lipgloss.JoinVertical(
				lipgloss.Left,
				m.viewport.View(),
				m.HelpView("↑/↓: Navigate • q: Quit"),
			),
		)
	}
	if m.width == 0 {
		return "loading..."
	}
	return lipgloss.Place(
		m.width,
		m.height,
		lipgloss.Center,
		lipgloss.Center,
		lipgloss.JoinVertical(
			lipgloss.Center,
			m.questions[m.index].question,
			m.styles.InputField.Render(current.input.View()),
			lipgloss.Place(
				120,
				1,
				lipgloss.Left,
				lipgloss.Center,
				m.HelpView("Next [ctrl+n]\tPrevious [ctrl+p]\tQuit [q]"),
			),
		),
	)
}

func main() {
	questions := []Question{
		NewShortQuestion("question1?"),
		NewShortQuestion("Question2?"),
		NewLongQuestion("Question3?"),
	}
	m := NewModel(questions)
	f, err := tea.LogToFile("debug.log", "debug")
	if err != nil {
		log.Fatalf("err: %v", err)
	}
	defer f.Close()
	p := tea.NewProgram(m, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
