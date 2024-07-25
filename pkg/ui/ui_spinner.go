package ui

import (
	"fmt"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	textStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("252")).Render
	spinnerStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("69"))
)

type SpinnerModel struct {
	Text    string
	Spinner spinner.Model
	Done    bool
}

func NewSpinner(s string) SpinnerModel {
	return InitialSpinner(s)
}

func InitialSpinner(s string) SpinnerModel {
	sp := spinner.New()
	sp.Spinner = spinner.Dot
	return SpinnerModel{
		Text:    textStyle(s),
		Spinner: sp,
		Done:    false,
	}
}

func (m SpinnerModel) Init() tea.Cmd {
	return m.Spinner.Tick
}

func (m SpinnerModel) Update(msg tea.Msg) (SpinnerModel, tea.Cmd) {
	return m, nil
}

func (m SpinnerModel) View() (s string) {
	s += fmt.Sprintf("\n %s %s\n\n", m.Spinner.View(), m.Text)
	return
}
