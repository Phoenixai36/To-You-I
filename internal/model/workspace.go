// T.U.I — To Unify Imagination
// internal/model/workspace.go — Workspace domain types
package model

import "time"

// WorkspaceStatus represents the overall health of a workspace.
type WorkspaceStatus string

const (
	WSIdle  WorkspaceStatus = "idle"
	WSBusy  WorkspaceStatus = "busy"
	WSError WorkspaceStatus = "error"
)

// WorkspaceEvent is a high-level event at workspace scope.
type WorkspaceEvent struct {
	ID          string         `json:"id"`
	WorkspaceID string         `json:"workspace_id"`
	AgentID     string         `json:"agent_id,omitempty"`
	Summary     string         `json:"summary"`
	Timestamp   time.Time      `json:"timestamp"`
	Meta        map[string]any `json:"meta,omitempty"`
}

// Workspace groups agents and events belonging to a single project.
type Workspace struct {
	ID     string          `json:"id"`
	Name   string          `json:"name"`
	Path   string          `json:"path"`
	Shell  string          `json:"shell"` // pwsh, zsh, bash, fish
	Status WorkspaceStatus `json:"status"`
	Agents []*Agent        `json:"agents"`
	Events []WorkspaceEvent `json:"events"`
	Meta   map[string]any  `json:"meta,omitempty"`
}

// ActiveAgents returns agents that are not idle or stopped.
func (w *Workspace) ActiveAgents() []*Agent {
	var active []*Agent
	for _, a := range w.Agents {
		if a.Status != StatusIdle && a.Status != StatusStopped {
			active = append(active, a)
		}
	}
	return active
}

// NeedsAttention returns true if any agent needs user input or has errored.
func (w *Workspace) NeedsAttention() bool {
	for _, a := range w.Agents {
		if a.Status == StatusWaitingInput || a.Status == StatusError {
			return true
		}
	}
	return false
}

// StatusIcon returns a display icon for the workspace status bar.
func (w *Workspace) StatusIcon() string {
	if w.NeedsAttention() {
		return "!"
	}
	switch w.Status {
	case WSBusy:
		return "●"
	case WSError:
		return "!"
	default:
		return "○"
	}
}
