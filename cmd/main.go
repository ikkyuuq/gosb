package main

import (
	"fmt"
	"go-structure-builder/pkg/builder"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	program := tea.NewProgram(builder.NewBuilder())
	if _, err := program.Run(); err != nil {
		fmt.Printf("Oh no %v", err)
		os.Exit(1)
	}
}
