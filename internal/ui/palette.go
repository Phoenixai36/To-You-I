// Package ui provides Terminal UI components for T.U.I (To Unify Imagination).
// palette.go — Fuzzy command palette powered by Bubble Tea + Lip Gloss.
// Activated with ⌘/Ctrl+P; filters UICommands in real-time.
package ui

import (
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/Phoenixai36/To-You-I/internal/model"
)

// PaletteOpenMsg opens the command palette.
type PaletteOpenMsg struct{}

// PaletteCloseMsg closes the command palette.
type PaletteCloseMsg struct{}

// PaletteSelectMsg is sent when a command is confirmed.
type PaletteSelectMsg struct {
	Command model.UICommand
}

// PaletteModel is the Bubble Tea model for the command palette overlay.
type PaletteModel struct {
	input    textinput.Model
	commands []model.UICommand
	filtered []model.UICommand
	cursor   int
	visible  bool
	styles   PaletteStyles
}

// PaletteStyles holds all Lip Gloss styles for the palette.
type PaletteStyles struct {
	Overlay    lipgloss.Style
	Input      lipgloss.Style
	Item       lipgloss.Style
	Selected   lipgloss.Style
	Category   lipgloss.Style
	Dimmed     lipgloss.Style
}

// defaultPaletteStyles returns Dark Phoenix – themed styles.
func defaultPaletteStyles() PaletteStyles {
	return PaletteStyles{
		Overlay: lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#FF2E63")).
			Padding(1, 2).
			Background(lipgloss.Color("#0D0D0D")).
			Width(60),
		Input: lipgloss.NewStyle().
			Foreground(lipgloss.Color("#EAEAEA")).
			Background(lipgloss.Color("#1A1A2E")).
			Padding(0, 1),
		Item: lipgloss.NewStyle().
			Foreground(lipgloss.Color("#EAEAEA")).
			PaddingLeft(2),
		Selected: lipgloss.NewStyle().
			Foreground(lipgloss.Color("#0D0D0D")).
			Background(lipgloss.Color("#08D9D6")).
			PaddingLeft(2).
			Bold(true),
		Category: lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FF2E63")).
			Bold(true).
			PaddingLeft(1),
		Dimmed: lipgloss.NewStyle().
			Foreground(lipgloss.Color("#444466")).
			PaddingLeft(4),
	}
}

// NewPaletteModel creates a new PaletteModel with the given commands.
func NewPaletteModel(cmds []model.UICommand) PaletteModel {
	ti := textinput.New()
	ti.Placeholder = "Search commands…"
	ti.CharLimit = 64
	ti.Width = 54

	return PaletteModel{
		input:    ti,
		commands: cmds,
		filtered: cmds,
		styles:   defaultPaletteStyles(),
	}
}

// Init implements tea.Model.
func (p PaletteModel) Init() tea.Cmd {
	return textinput.Blink
}

// Update implements tea.Model.
func (p PaletteModel) Update(msg tea.Msg) (PaletteModel, tea.Cmd) {
	if !p.visible {
		switch msg.(type) {
		case PaletteOpenMsg:
			p.visible = true
			p.input.Focus()
			p.input.SetValue("")
			p.filter("")
			return p, textinput.Blink
		}
		return p, nil
	}

	switch m := msg.(type) {
	case tea.KeyMsg:
		switch m.Type {
		case tea.KeyEsc:
			p.visible = false
			p.input.Blur()
			return p, func() tea.Msg { return PaletteCloseMsg{} }
		case tea.KeyEnter:
			if len(p.filtered) > 0 {
				cmd := p.filtered[p.cursor]
				p.visible = false
				p.input.Blur()
				return p, func() tea.Msg { return PaletteSelectMsg{Command: cmd} }
			}
		case tea.KeyUp:
			if p.cursor > 0 {
				p.cursor--
			}
		case tea.KeyDown:
			if p.cursor < len(p.filtered)-1 {
				p.cursor++
			}
		}
	}

	var cmd tea.Cmd
	p.input, cmd = p.input.Update(msg)
	p.filter(p.input.Value())
	return p, cmd
}

// filter updates the filtered list based on the query string.
func (p *PaletteModel) filter(query string) {
	p.cursor = 0
	if query == "" {
		p.filtered = p.commands
		return
	}
	q := strings.ToLower(query)
	out := p.filtered[:0]
	for _, c := range p.commands {
		if strings.Contains(strings.ToLower(c.Title), q) ||
			strings.Contains(strings.ToLower(c.Category), q) {
			out = append(out, c)
		}
	}
	p.filtered = out
}

// View implements tea.Model.
func (p PaletteModel) View() string {
	if !p.visible {
		return ""
	}

	var b strings.Builder
	b.WriteString(p.styles.Input.Render(p.input.View()))
	b.WriteString("\n\n")

	if len(p.filtered) == 0 {
		b.WriteString(p.styles.Dimmed.Render("No commands match…"))
	} else {
		for i, cmd := range p.filtered {
			line := cmd.Icon + "  " + cmd.Title
			if cmd.Shortcut != "" {
				line += "  " + p.styles.Dimmed.Render(cmd.Shortcut)
			}
			if i == p.cursor {
				b.WriteString(p.styles.Selected.Render(line))
			} else {
				b.WriteString(p.styles.Item.Render(line))
			}
			b.WriteString("\n")
		}
	}

	return p.styles.Overlay.Render(b.String())
}

// Visible returns true when the palette is open.
func (p PaletteModel) Visible() bool { return p.visible }
