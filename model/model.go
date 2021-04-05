package model

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/glamour"
	"github.com/charmbracelet/lipgloss"
	"github.com/maaslalani/slides/styles"
)

type Model struct {
	Slides   []string
	Page     int
	Author   string
	Date     string
	viewport viewport.Model
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.viewport.Width = msg.Width
		m.viewport.Height = msg.Height
		return m, nil

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up", "k", "right", "l", "enter", "n":
			if m.Page < len(m.Slides)-1 {
				m.Page++
			}
		case "down", "j", "left", "h", "p":
			if m.Page > 0 {
				m.Page--
			}
		}
	}

	return m, nil
}

func (m Model) View() string {
	slide, _ := glamour.Render(m.Slides[m.Page], "dark")
	slide = styles.Slide.Render(slide)

	left := styles.Author.Render(m.Author) + styles.Date.Render(m.Date)
	right := styles.Page.Render(fmt.Sprintf("Slide %d / %d", m.Page, len(m.Slides)-1))
	status := styles.Status.Render(styles.SpreadHorizontal(left, right, m.viewport.Width))

	padding := strings.Repeat("\n", max(m.viewport.Height-lipgloss.Height(slide)-lipgloss.Height(status), 0))

	return slide + padding + status
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
