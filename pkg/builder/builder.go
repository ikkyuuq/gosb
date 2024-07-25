package builder

import (
	"go-structure-builder/pkg/ui"

	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	Form    *ui.FormModel
	Spinner *ui.SpinnerModel
	State   int
}

const (
	StateForm = iota
	StateSpinner
	StateComplete
)

func NewBuilder() *Model {
	return initialBuilder()
}

func initialBuilder() *Model {
	form := ui.NewForm()
	return &Model{
		Form: form,
	}
}

func (m Model) Init() tea.Cmd {
	return m.Form.Init()
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}

func (m Model) View() string {
	switch m.State {
	case StateComplete:
		return "Complete"
	case StateSpinner:
		return m.Spinner.View()
	default:
		return m.Form.View()
	}
}
