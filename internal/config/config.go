// Package config manages T.U.I runtime configuration.
// config.go — loads, validates, and provides defaults for all tuneable knobs.
// Config is read from ~/.config/tui/config.toml (XDG) or overridden by env vars.
package config

import (
	"errors"
	"os"
	"path/filepath"
	"time"
)

// Defaults
const (
	DefaultServerAddr    = "127.0.0.1:7788"
	DefaultLogLevel      = "info"
	DefaultGlitchEnabled = true
	DefaultTheme         = "dark-phoenix"
	DefaultWorkspaceDir  = "~/.tui/workspaces"
	DefaultSSEPath       = "/events"
	DefaultReadTimeout   = 5 * time.Second
	DefaultWriteTimeout  = 10 * time.Second
)

// Config is the top-level configuration for T.U.I.
type Config struct {
	// Server configures the local SSE HTTP server.
	Server ServerConfig `toml:"server"`

	// UI configures visual behaviour.
	UI UIConfig `toml:"ui"`

	// Workspaces configures workspace persistence.
	Workspaces WorkspacesConfig `toml:"workspaces"`

	// Log configures logging verbosity.
	Log LogConfig `toml:"log"`
}

// ServerConfig holds parameters for the embedded HTTP/SSE server.
type ServerConfig struct {
	// Addr is the listen address (host:port).
	Addr string `toml:"addr" env:"TUI_SERVER_ADDR"`

	// SSEPath is the HTTP path for the SSE stream.
	SSEPath string `toml:"sse_path"`

	// ReadTimeout for inbound HTTP requests.
	ReadTimeout time.Duration `toml:"read_timeout"`

	// WriteTimeout for SSE stream writes.
	WriteTimeout time.Duration `toml:"write_timeout"`
}

// UIConfig controls TUI appearance and effects.
type UIConfig struct {
	// Theme selects the Lip Gloss colour palette.
	// Supported: "dark-phoenix" (default).
	Theme string `toml:"theme" env:"TUI_THEME"`

	// GlitchEnabled toggles the scanline glitch renderer.
	GlitchEnabled bool `toml:"glitch_enabled" env:"TUI_GLITCH"`

	// GlitchIntensity is the line-corruption probability (0–100).
	GlitchIntensity int `toml:"glitch_intensity"`

	// GlitchInterval is how frequently the glitch refreshes.
	GlitchInterval time.Duration `toml:"glitch_interval"`

	// WorkspacePanelWidth is the pixel width of the left rail.
	WorkspacePanelWidth int `toml:"workspace_panel_width"`
}

// WorkspacesConfig controls workspace persistence.
type WorkspacesConfig struct {
	// Dir is the directory where workspace state is stored.
	Dir string `toml:"dir" env:"TUI_WORKSPACES_DIR"`

	// AutoSave enables periodic workspace state snapshots.
	AutoSave bool `toml:"auto_save"`

	// AutoSaveInterval is the save cadence when AutoSave is true.
	AutoSaveInterval time.Duration `toml:"auto_save_interval"`
}

// LogConfig controls log output.
type LogConfig struct {
	// Level is one of: debug, info, warn, error.
	Level string `toml:"level" env:"TUI_LOG_LEVEL"`

	// File is an optional path to write logs. Stdout is used when empty.
	File string `toml:"file" env:"TUI_LOG_FILE"`
}

// Default returns a Config populated with sensible defaults.
func Default() Config {
	return Config{
		Server: ServerConfig{
			Addr:         DefaultServerAddr,
			SSEPath:      DefaultSSEPath,
			ReadTimeout:  DefaultReadTimeout,
			WriteTimeout: DefaultWriteTimeout,
		},
		UI: UIConfig{
			Theme:               DefaultTheme,
			GlitchEnabled:      DefaultGlitchEnabled,
			GlitchIntensity:    12,
			GlitchInterval:     120 * time.Millisecond,
			WorkspacePanelWidth: 22,
		},
		Workspaces: WorkspacesConfig{
			Dir:              expandHome(DefaultWorkspaceDir),
			AutoSave:         true,
			AutoSaveInterval: 30 * time.Second,
		},
		Log: LogConfig{
			Level: DefaultLogLevel,
		},
	}
}

// Validate checks that the Config values are consistent and returns an error
// describing the first violation found.
func (c *Config) Validate() error {
	if c.Server.Addr == "" {
		return errors.New("config: server.addr must not be empty")
	}
	if c.UI.WorkspacePanelWidth < 10 {
		return errors.New("config: ui.workspace_panel_width must be >= 10")
	}
	if c.UI.GlitchIntensity < 0 || c.UI.GlitchIntensity > 100 {
		return errors.New("config: ui.glitch_intensity must be in [0, 100]")
	}
	validLevels := map[string]bool{"debug": true, "info": true, "warn": true, "error": true}
	if !validLevels[c.Log.Level] {
		return errors.New("config: log.level must be one of debug/info/warn/error")
	}
	return nil
}

// ConfigDir returns the XDG-compatible directory for T.U.I config files.
func ConfigDir() string {
	if xdg := os.Getenv("XDG_CONFIG_HOME"); xdg != "" {
		return filepath.Join(xdg, "tui")
	}
	return filepath.Join(expandHome("~"), ".config", "tui")
}

// DefaultConfigPath returns the canonical path of the TOML config file.
func DefaultConfigPath() string {
	return filepath.Join(ConfigDir(), "config.toml")
}

// expandHome replaces a leading ~ with the user home directory.
func expandHome(p string) string {
	if len(p) == 0 || p[0] != '~' {
		return p
	}
	home, err := os.UserHomeDir()
	if err != nil {
		return p
	}
	return filepath.Join(home, p[1:])
}
