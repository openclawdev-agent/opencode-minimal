package tui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/opencode-ai/opencode/internal/app"
)

type stubModel struct{}

func (m stubModel) Init() tea.Cmd            { return nil }
func (m stubModel) Update(tea.Msg) (tea.Model, tea.Cmd) { return m, nil }
func (m stubModel) View() string             { return "" }

func New(_ *app.App) tea.Model { return stubModel{} }
