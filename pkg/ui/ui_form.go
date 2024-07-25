package ui

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
)

type FormModel struct {
	form *huh.Form
}

type FormData struct {
	ProjectName string
	Database    string
	Framework   string
	WorkDir     string
}

func NewForm() FormModel {
	return initialForm()
}

func initialForm() FormModel {
	m := FormModel{}
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("cannot find your current directory")
		os.Exit(1)
	}
	fd := FormData{}
	fd.WorkDir = dir
	m.form = huh.NewForm(
		huh.NewGroup(
			huh.NewNote().Title("GOSB").
				Description("A common question developers new to Go have is “How do I organize my Go project?”\nAnd, Yes! that why the goal of GOSB is to build your project structure easier,\nyou might learn by yourself from go.dev\n\nHuge credit: Melkey\n\n").Next(true).NextLabel("Let's Build!"),
		),
		huh.NewGroup(
			huh.NewInput().Title("Project Name").
				Value(&fd.ProjectName).Key("projectName").
				Description("Project directory is on: ").
				DescriptionFunc(
					func() string {
						s := fmt.Sprintf("Project directory is on: %v/%v", fd.WorkDir, fd.ProjectName)
						return s
					}, &fd.ProjectName),
		),
		huh.NewGroup(
			huh.NewSelect[string]().Title("Database").Description("Choose a database").
				Options(huh.NewOptions("None", "MySQL", "MongoDB", "ProtgreSQL")...).
				Value(&fd.Database).Key("database"),
		),
		huh.NewGroup(
			huh.NewSelect[string]().Title("Framework").Description("Choose a Go web framework").
				Options(huh.NewOptions("Standard Library", "Chi", "Echo", "Fiber", "Gin", "Gorilla", "HttpRouter")...).
				Value(&fd.Framework).Key("framework"),
			huh.NewConfirm().Title("Are you sure?").Affirmative("Build!").Negative("Wait!"),
		),
	)
	return m
}

func (m FormModel) State() huh.FormState {
	return m.form.State
}

func (m FormModel) Init() tea.Cmd {
	return m.form.Init()
}

func (m FormModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd
	form, cmd := m.form.Update(msg)
	if f, ok := form.(*huh.Form); ok {
		m.form = f
		cmds = append(cmds, cmd)
	}
	return m, tea.Batch(cmds...)
}

func (m FormModel) View() string {
	return m.form.View()
}
