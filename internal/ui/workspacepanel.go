// Package ui provides Terminal UI components for T.U.I (To Unify Imagination).
// workspacepanel.go — Workspace switcher panel (left rail).
// Displays active workspaces, agent count, and status indicators.
package ui

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/Phoenixai36/To-You-I/internal/model"
)

// WorkspacePanelModel is the Bubble Tea model for the workspace rail.
type WorkspacePanelModel struct {
	workspaces []model.Workspace
	active     int
	width      int
	styles     workspacePanelStyles
}

type workspacePanelStyles struct {
	Panel       lipgloss.Style
	Header      lipgloss.Style
	Item        lipgloss.Style
	ActiveItem  lipgloss.Style
	Badge       lipgloss.Style
	StatusDot   lipgloss.Style
	Footer      lipgloss.Style
}

func defaultWorkspacePanelStyles(width int) workspacePanelStyles {
	return workspacePanelStyles{
		Panel: lipgloss.NewStyle().
			Width(width).
			BorderRight(true).
			BorderStyle(lipgloss.ThickBorder()).
			BorderForeground(lipgloss.Color("#252A34")).
			Background(lipgloss.Color("#0D0D0D")).
			PaddingTop(1),
		Header: lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FF2E63")).
			Bold(true).
			PaddingLeft(2).
			MarginBottom(1),
		Item: lipgloss.NewStyle().
			Foreground(lipgloss.Color("#EAEAEA")).
			PaddingLeft(2).
			Width(width - 2),
		ActiveItem: lipgloss.NewStyle().
			Foreground(lipgloss.Color("#0D0D0D")).
			Background(lipgloss.Color("#08D9D6")).
			Bold(true).
			PaddingLeft(2).
			Width(width - 2),
		Badge: lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FF2E63")).
			Bold(true),
		StatusDot: lipgloss.NewStyle().
			Foreground(lipgloss.Color("#08D9D6")),
		Footer: lipgloss.NewStyle().
			Foreground(lipgloss.Color("#444466")).
			PaddingLeft(2).
			MarginTop(1),
	}
}

// NewWorkspacePanelModel creates a new workspace panel.
func NewWorkspacePanelModel(workspaces []model.Workspace, panelWidth int) WorkspacePanelModel {
	return WorkspacePanelModel{
		workspaces: workspaces,
		width:      panelWidth,
		styles:     defaultWorkspacePanelStyles(panelWidth),
	}
}

// Init implements tea.Model.
func (w WorkspacePanelModel) Init() tea.Cmd { return nil }

// Update implements tea.Model.
func (w WorkspacePanelModel) Update(msg tea.Msg) (WorkspacePanelModel, tea.Cmd) {
	switch m := msg.(type) {
	case tea.KeyMsg:
		switch m.String() {
		case "j", "down":
			if w.active < len(w.workspaces)-1 {
				w.active++
			}
		case "k", "up":
			if w.active > 0 {
				w.active--
			}
		}
	}
	return w, nil
}

// ActiveWorkspace returns the currently focused workspace.
func (w WorkspacePanelModel) ActiveWorkspace() *model.Workspace {
	if len(w.workspaces) == 0 {
		return nil
	}
	return &w.workspaces[w.active]
}

// View implements tea.Model.
func (w WorkspacePanelModel) View() string {
	var b strings.Builder

	// Header
	b.WriteString(w.styles.Header.Render("🔥 T.U.I"))
	b.WriteString("\n")

	// Workspace list
	for i, ws := range w.workspaces {
		dot := w.styles.StatusDot.Render("•")
		if ws.Status == model.WorkspaceStatusError {
			dot = lipgloss.NewStyle().Foreground(lipgloss.Color("#FF2E63")).Render("•")
		} else if ws.Status == model.WorkspaceStatusIdle {
			dot = lipgloss.NewStyle().Foreground(lipgloss.Color("#444466")).Render("•")
		}

		agentCount := ""
		if len(ws.Agents) > 0 {
			agentCount = w.styles.Badge.Render(fmt.Sprintf(" [%d]", len(ws.Agents)))
		}

		line := dot + " " + ws.Name + agentCount

		if i == w.active {
			b.WriteString(w.styles.ActiveItem.Render(line))
		} else {
			b.WriteString(w.styles.Item.Render(line))
		}
		b.WriteString("\n")
	}

	// Footer hint
	b.WriteString(w.styles.Footer.Render("j/k navigate"))

	return w.styles.Panel.Render(b.String())
}
