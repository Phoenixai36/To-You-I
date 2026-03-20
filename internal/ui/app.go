// T.U.I — To Unify Imagination
// internal/ui/app.go — Root Bubble Tea model
package ui

import (
	"github.com/Phoenixai36/To-You-I/internal/model"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// Panel represents which panel has focus.
type Panel int

const (
	PanelRail Panel = iota
	PanelSidebar
	PanelFeed
	PanelDetail
	PanelPalette
)

// App is the root Bubble Tea model for T.U.I.
type App struct {
	// Layout
	width  int
	height int
	focus  Panel

	// Data
	workspaces     []*model.Workspace
	activeWS       int
	selectedAgent  int
	paletteOpen    bool
	paletteQuery   string
	paletteEntries []model.PaletteEntry

	// Sub-models (future: each panel is its own Bubble Tea model)
	// rail    Rail
	// sidebar Sidebar
	// feed    Feed
}

// NewApp creates the root app with demo workspaces.
func NewApp() *App {
	// Seed demo data until adapters are wired in
	ws := []*model.Workspace{
		{
			ID:    "backend-api",
			Name:  "Backend API",
			Shell: "pwsh",
			Status: model.WSBusy,
			Agents: []*model.Agent{
				{ID: "claude-1", Name: "Claude", Kind: "claude", Status: model.StatusWaitingInput},
				{ID: "lint-1",   Name: "Linter", Kind: "custom", Status: model.StatusRunning},
				{ID: "test-1",   Name: "Tests",  Kind: "custom", Status: model.StatusIdle},
			},
		},
		{
			ID:    "frontend-web",
			Name:  "Frontend Web",
			Shell: "zsh",
			Status: model.WSIdle,
			Agents: []*model.Agent{
				{ID: "goose-1", Name: "Goose", Kind: "goose", Status: model.StatusIdle},
			},
		},
		{
			ID:    "infra",
			Name:  "Infra",
			Shell: "bash",
			Status: model.WSError,
			Agents: []*model.Agent{
				{ID: "deploy-1", Name: "Deploy Bot", Kind: "custom", Status: model.StatusError},
			},
		},
	}

	return &App{
		workspaces: ws,
		focus:      PanelSidebar,
	}
}

// Init satisfies tea.Model.
func (a *App) Init() tea.Cmd {
	return nil
}

// Update handles all keyboard events and messages.
func (a *App) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		a.width = msg.Width
		a.height = msg.Height

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			if !a.paletteOpen {
				return a, tea.Quit
			}
		case "ctrl+p":
			a.paletteOpen = !a.paletteOpen
			if a.paletteOpen {
				a.buildPalette()
			}
		case "esc":
			if a.paletteOpen {
				a.paletteOpen = false
			}
		case "left":
			if !a.paletteOpen && a.activeWS > 0 {
				a.activeWS--
				a.selectedAgent = 0
			}
		case "right":
			if !a.paletteOpen && a.activeWS < len(a.workspaces)-1 {
				a.activeWS++
				a.selectedAgent = 0
			}
		case "up":
			if !a.paletteOpen {
				ws := a.workspaces[a.activeWS]
				if a.selectedAgent > 0 {
					a.selectedAgent--
				} else {
					a.selectedAgent = len(ws.Agents) - 1
				}
			}
		case "down":
			if !a.paletteOpen {
				ws := a.workspaces[a.activeWS]
				if a.selectedAgent < len(ws.Agents)-1 {
					a.selectedAgent++
				} else {
					a.selectedAgent = 0
				}
			}
		}
	}
	return a, nil
}

// View renders the full HUD.
func (a *App) View() string {
	if a.width == 0 {
		return "Loading T.U.I..."
	}

	rail    := a.renderRail()
	sidebar := a.renderSidebar()
	feed    := a.renderFeed()
	statBar := a.renderStatusBar()

	sidebarW := 22
	feedW    := a.width - sidebarW - 1
	feedH    := a.height - 3 // rail + statusbar

	sidebarView := StyleSidebar.Width(sidebarW).Height(feedH).Render(sidebar)
	feedView    := StyleFeed.Width(feedW).Height(feedH).Render(feed)

	main := lipgloss.JoinHorizontal(lipgloss.Top, sidebarView, feedView)
	view := lipgloss.JoinVertical(lipgloss.Left, rail, main, statBar)

	if a.paletteOpen {
		view = a.renderPaletteOverlay(view)
	}

	return view
}

func (a *App) buildPalette() {
	ws := a.workspaces[a.activeWS]
	a.paletteEntries = model.DefaultPaletteEntries(ws)
}

func (a *App) activeWorkspace() *model.Workspace {
	return a.workspaces[a.activeWS]
}
