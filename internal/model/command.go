// T.U.I — To Unify Imagination
// internal/model/command.go — UI command types
package model

// UICommandType enumerates all actions the HUD can dispatch.
type UICommandType string

const (
	// Workspace commands
	CmdOpenWorkspace  UICommandType = "open_workspace"
	CmdCloseWorkspace UICommandType = "close_workspace"

	// Agent commands
	CmdOpenAgentUI   UICommandType = "open_agent_ui"
	CmdSendInput     UICommandType = "send_input"
	CmdRestartAgent  UICommandType = "restart_agent"
	CmdStopAgent     UICommandType = "stop_agent"

	// Task commands
	CmdRunTask UICommandType = "run_task"

	// Shell commands
	CmdOpenShell  UICommandType = "open_shell"
	CmdRunInShell UICommandType = "run_in_shell"

	// UI commands
	CmdTogglePalette UICommandType = "toggle_palette"
	CmdFocusPanel    UICommandType = "focus_panel"
	CmdQuit          UICommandType = "quit"
)

// UICommand is a typed action dispatched from the HUD to adapters or the shell.
type UICommand struct {
	Type        UICommandType  `json:"type"`
	WorkspaceID string         `json:"workspace_id,omitempty"`
	AgentID     string         `json:"agent_id,omitempty"`
	Payload     map[string]any `json:"payload,omitempty"`
}

// PaletteEntry is a single entry in the command palette.
type PaletteEntry struct {
	Label       string
	Description string
	Command     UICommand
}

// DefaultPaletteEntries returns the built-in command palette entries for a workspace.
func DefaultPaletteEntries(ws *Workspace) []PaletteEntry {
	entries := []PaletteEntry{
		{
			Label:       "Run tests",
			Description: "Run test suite in " + ws.Name,
			Command:     UICommand{Type: CmdRunTask, WorkspaceID: ws.ID, Payload: map[string]any{"task": "test"}},
		},
		{
			Label:       "Open shell",
			Description: "Open " + ws.Shell + " in " + ws.Name,
			Command:     UICommand{Type: CmdOpenShell, WorkspaceID: ws.ID},
		},
		{
			Label:       "Restart all agents",
			Description: "Restart all agents in " + ws.Name,
			Command:     UICommand{Type: CmdRestartAgent, WorkspaceID: ws.ID},
		},
	}
	for _, a := range ws.Agents {
		a := a
		entries = append(entries, PaletteEntry{
			Label:       "Ask " + a.Name,
			Description: "Open " + a.Name + " agent UI",
			Command:     UICommand{Type: CmdOpenAgentUI, WorkspaceID: ws.ID, AgentID: a.ID},
		})
	}
	return entries
}
