package ui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
)

type FormModel struct {
	form *huh.Form
}

func NewForm() FormModel {
	return initialForm()
}

func initialForm() FormModel {
	m := FormModel{}
	m.form = huh.NewForm()
	return m
}

func (m FormModel) Init() tea.Cmd {
	return m.form.Init()
}

func (m FormModel) Update(msg tea.Msg) (FormModel, tea.Cmd) {
	return m, nil
}

func (m FormModel) View() string {
	return m.form.View()
}
