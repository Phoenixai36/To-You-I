// T.U.I — To Unify Imagination
// internal/ui/styles.go — Dark Phoenix Lip Gloss theme
// Palette: deep black + burnt orange + fire red + metallic gold + glitch cyan
package ui

import "github.com/charmbracelet/lipgloss"

// Dark Phoenix color palette
const (
	// Background
	ColorBg      = lipgloss.Color("#0A0A0A")
	ColorBgPanel = lipgloss.Color("#111111")
	ColorBgHover = lipgloss.Color("#1A1A1A")

	// Primary — Burnt orange / fire
	ColorOrange = lipgloss.Color("#E8890C")
	ColorRed    = lipgloss.Color("#CC2200")
	ColorFire   = lipgloss.Color("#FF4500")

	// Accent — Gold metallic
	ColorGold      = lipgloss.Color("#D4A017")
	ColorGoldLight = lipgloss.Color("#F5C842")

	// Glitch — Cyan electric
	ColorCyan      = lipgloss.Color("#00FFFF")
	ColorCyanDim   = lipgloss.Color("#007777")

	// Text
	ColorText    = lipgloss.Color("#E8E8E8")
	ColorTextDim = lipgloss.Color("#666666")
	ColorTextMid = lipgloss.Color("#AAAAAA")

	// Status colors
	ColorStatusOk      = lipgloss.Color("#44BB44")
	ColorStatusWarn    = lipgloss.Color("#E8890C")
	ColorStatusError   = lipgloss.Color("#CC2200")
	ColorStatusWaiting = lipgloss.Color("#00FFFF")
)

// ---- Layout styles ----

var (
	// Root app container
	StyleApp = lipgloss.NewStyle().
		Background(ColorBg).
		Foreground(ColorText)

	// Top workspace rail
	StyleRail = lipgloss.NewStyle().
		Background(ColorBgPanel).
		BorderStyle(lipgloss.NormalBorder()).
		BorderBottom(true).
		BorderForeground(ColorOrange).
		PaddingLeft(1).PaddingRight(1)

	// Workspace pill — inactive
	StyleWorkspacePill = lipgloss.NewStyle().
		Foreground(ColorTextDim).
		Padding(0, 2).
		MarginRight(1)

	// Workspace pill — active (Dark Phoenix accent)
	StyleWorkspacePillActive = lipgloss.NewStyle().
		Foreground(ColorBg).
		Background(ColorOrange).
		Bold(true).
		Padding(0, 2).
		MarginRight(1)

	// Workspace pill — needs attention
	StyleWorkspacePillAlert = lipgloss.NewStyle().
		Foreground(ColorBg).
		Background(ColorRed).
		Bold(true).
		Padding(0, 2).
		MarginRight(1)

	// Left sidebar
	StyleSidebar = lipgloss.NewStyle().
		Background(ColorBgPanel).
		BorderStyle(lipgloss.NormalBorder()).
		BorderRight(true).
		BorderForeground(ColorOrange).
		Padding(1, 1)

	// Sidebar title
	StyleSidebarTitle = lipgloss.NewStyle().
		Foreground(ColorOrange).
		Bold(true).
		MarginBottom(1)

	// Agent row — normal
	StyleAgentRow = lipgloss.NewStyle().
		Foreground(ColorText).
		Padding(0, 1)

	// Agent row — selected
	StyleAgentRowSelected = lipgloss.NewStyle().
		Foreground(ColorGold).
		Background(ColorBgHover).
		Bold(true).
		Padding(0, 1)

	// Activity feed panel
	StyleFeed = lipgloss.NewStyle().
		Background(ColorBg).
		Padding(1, 2)

	// Feed title
	StyleFeedTitle = lipgloss.NewStyle().
		Foreground(ColorGold).
		Bold(true).
		MarginBottom(1)

	// Feed event row
	StyleFeedRow = lipgloss.NewStyle().
		Foreground(ColorTextMid)

	// Feed event — needs attention highlight
	StyleFeedRowAlert = lipgloss.NewStyle().
		Foreground(ColorCyan).
		Bold(true)

	// Status bar (bottom)
	StyleStatusBar = lipgloss.NewStyle().
		Background(ColorBgPanel).
		Foreground(ColorTextDim).
		BorderStyle(lipgloss.NormalBorder()).
		BorderTop(true).
		BorderForeground(ColorOrange).
		PaddingLeft(2).PaddingRight(2)

	// Command palette overlay
	StylePalette = lipgloss.NewStyle().
		Background(ColorBgPanel).
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(ColorCyan).
		Padding(1, 2)

	// Palette title
	StylePaletteTitle = lipgloss.NewStyle().
		Foreground(ColorCyan).
		Bold(true).
		MarginBottom(1)

	// Palette entry — normal
	StylePaletteEntry = lipgloss.NewStyle().
		Foreground(ColorText)

	// Palette entry — selected
	StylePaletteEntrySelected = lipgloss.NewStyle().
		Foreground(ColorOrange).
		Bold(true)

	// Glitch accent — tiny separator line
	StyleGlitchLine = lipgloss.NewStyle().
		Foreground(ColorCyanDim).
		Bold(false)
)

// StatusColor returns the appropriate Lip Gloss color for an agent status.
func StatusColor(status string) lipgloss.Color {
	switch status {
	case "running", "busy":
		return ColorStatusOk
	case "waiting_input":
		return ColorStatusWaiting
	case "error":
		return ColorStatusError
	default:
		return ColorTextDim
	}
}
