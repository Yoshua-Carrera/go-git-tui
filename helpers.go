package main

import (
	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/viewport"
	"github.com/charmbracelet/glamour"
	"github.com/charmbracelet/lipgloss"
)

type Styles struct {
	BorderColor lipgloss.Color
	InputField  lipgloss.Style
}

type ShortAnswerField struct {
	textinput textinput.Model
}

type LongAnswerField struct {
	textarea textarea.Model
}

func (m *model) Next() {
	if m.index < len(m.questions)-1 {
		m.index++
	} else {
		m.index = 0
	}
}

func (m *model) Prev() {
	if m.index == 0 {
		m.index--
	} else {
		m.index = 0
	}
}

func (m model) HelpView(msg string) string {
	return lipgloss.NewStyle().Foreground(lipgloss.Color("241")).Render(msg)
}

func DefaultStyles() *Styles {
	s := new(Styles)
	s.BorderColor = lipgloss.Color("150")
	s.InputField = lipgloss.NewStyle().
		BorderForeground(s.BorderColor).
		BorderStyle(lipgloss.NormalBorder()).
		Padding(1).
		Width(120)
	return s
}

func DefaultViewportStyles() lipgloss.Style {
	s := lipgloss.NewStyle().
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("62")).
		PaddingRight(2)
	return s
}

func NewModel(questions []Question) *model {
	styles := DefaultStyles()
	vp := viewport.New(120, 40)
	vp.Style = DefaultViewportStyles()
	return &model{
		questions: questions,
		styles:    styles,
		viewport:  vp,
	}
}

func (m *model) GlamourRender(content string) error {
	gutter := 2
	glamourRenderWidth := m.width - m.viewport.Style.GetHorizontalFrameSize() - gutter
	glamourRenderer, err := glamour.NewTermRenderer(
		glamour.WithAutoStyle(),
		glamour.WithWordWrap(glamourRenderWidth),
	)
	if err != nil {
		return err
	}

	str, err := glamourRenderer.Render(content)
	if err != nil {
		return err
	}

	m.viewport.SetContent(str)
	return nil
}
