package builder

import (
	"go-structure-builder/pkg/ui"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
)

type Model struct {
	Form    ui.FormModel
	Spinner ui.SpinnerModel
	State   int
}

const (
	StateForm = iota
	StateSpinner
	StateComplete
)

func NewBuilder() Model {
	return initialBuilder()
}

func initialBuilder() Model {
	form := ui.NewForm()
	return Model{
		Form: form,
	}
}

func (m Model) Init() tea.Cmd {
	return m.Form.Init()
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc", "q", "ctrl+c":
			return m, tea.Quit
		}
	}
	switch m.State {
	case StateForm:
		_, cmd := m.Form.Update(msg)
		if m.Form.State() == huh.StateCompleted {
			m.State = StateSpinner
			m.Spinner = ui.NewSpinner("Preparing...")
			return m, m.Spinner.Init()
		}
		return m, cmd
	case StateSpinner:
		spinnerModel, cmd := m.Spinner.Update(msg)
		m.Spinner = spinnerModel.(ui.SpinnerModel)
		if m.Spinner.State == ui.StateDone {
			m.State = StateComplete
		}
		return m, cmd
	default:
		return m, nil
	}
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
