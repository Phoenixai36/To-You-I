// Package ui provides Terminal UI components for T.U.I (To Unify Imagination).
// glitch.go — Dark Phoenix glitch effect renderer.
// Produces scanline corruption, chroma-shift, and pixel-drop animations
// using Lip Gloss styled strings and Bubble Tea tick commands.
package ui

import (
	"math/rand"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// GlitchMsg is sent on every glitch tick.
type GlitchMsg struct{}

// GlitchTickCmd returns a Bubble Tea command that fires a GlitchMsg after d.
func GlitchTickCmd(d time.Duration) tea.Cmd {
	return tea.Tick(d, func(t time.Time) tea.Msg {
		return GlitchMsg{}
	})
}

// GlitchConfig controls glitch intensity parameters.
type GlitchConfig struct {
	// Probability (0-100) that a line gets corrupted on each tick.
	LineProbability int
	// Maximum horizontal shift in characters.
	MaxShift int
	// ChromaChars are the substitution characters used during corruption.
	ChromaChars string
	// TickInterval is how frequently the glitch refreshes.
	TickInterval time.Duration
}

// DefaultGlitchConfig returns a sensible default for Dark Phoenix aesthetic.
func DefaultGlitchConfig() GlitchConfig {
	return GlitchConfig{
		LineProbability: 12,
		MaxShift:        4,
		ChromaChars:     "▓░▒█▄▀■□▪▫◆◇●○",
		TickInterval:    120 * time.Millisecond,
	}
}

// GlitchRenderer applies glitch effects to a block of text.
type GlitchRenderer struct {
	cfg    GlitchConfig
	rng    *rand.Rand
	active bool
}

// NewGlitchRenderer creates a renderer seeded from current time.
func NewGlitchRenderer(cfg GlitchConfig) *GlitchRenderer {
	return &GlitchRenderer{
		cfg: cfg,
		rng: rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

// SetActive toggles the glitch effect on or off.
func (g *GlitchRenderer) SetActive(v bool) { g.active = v }

// Active returns whether the renderer is currently applying effects.
func (g *GlitchRenderer) Active() bool { return g.active }

// Render applies glitch corruption to src and returns styled output.
// When inactive, Render returns src unmodified.
func (g *GlitchRenderer) Render(src string) string {
	if !g.active {
		return src
	}
	lines := strings.Split(src, "\n")
	out := make([]string, len(lines))
	for i, line := range lines {
		if g.rng.Intn(100) < g.cfg.LineProbability {
			out[i] = g.corruptLine(line)
		} else {
			out[i] = line
		}
	}
	return strings.Join(out, "\n")
}

// corruptLine applies one of three glitch effects to a single line.
func (g *GlitchRenderer) corruptLine(line string) string {
	switch g.rng.Intn(3) {
	case 0:
		return g.chromaShift(line)
	case 1:
		return g.horizontalShift(line)
	default:
		return g.scanlineCorrupt(line)
	}
}

// chromaShift replaces a random slice of runes with glitch characters.
func (g *GlitchRenderer) chromaShift(line string) string {
	runes := []rune(line)
	if len(runes) == 0 {
		return line
	}
	chroma := []rune(g.cfg.ChromaChars)
	start := g.rng.Intn(len(runes))
	length := g.rng.Intn(max(1, len(runes)-start)) + 1
	for j := start; j < start+length && j < len(runes); j++ {
		runes[j] = chroma[g.rng.Intn(len(chroma))]
	}
	glitchStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#FF2E63")).
		Bold(true)
	corrupt := glitchStyle.Render(string(runes[start : start+length]))
	var b strings.Builder
	b.WriteString(string(runes[:start]))
	b.WriteString(corrupt)
	b.WriteString(string(runes[start+length:]))
	return b.String()
}

// horizontalShift prepends spaces to simulate pixel-row displacement.
func (g *GlitchRenderer) horizontalShift(line string) string {
	shift := g.rng.Intn(g.cfg.MaxShift + 1)
	pad := strings.Repeat(" ", shift)
	if len(line) > shift {
		return pad + line[:len(line)-shift]
	}
	return pad
}

// scanlineCorrupt randomly inverts brightness styling on the whole line.
func (g *GlitchRenderer) scanlineCorrupt(line string) string {
	colors := []string{"#08D9D6", "#FF2E63", "#252A34", "#EAEAEA"}
	fg := colors[g.rng.Intn(len(colors))]
	s := lipgloss.NewStyle().Foreground(lipgloss.Color(fg)).Faint(true)
	return s.Render(line)
}

// max returns the larger of two ints (backport for Go < 1.21).
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
