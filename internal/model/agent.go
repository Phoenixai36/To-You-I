// T.U.I — To Unify Imagination
// internal/model/agent.go — Core agent domain types
package model

import "time"

// AgentStatus represents the current state of an agent.
type AgentStatus string

const (
	StatusIdle         AgentStatus = "idle"
	StatusRunning      AgentStatus = "running"
	StatusBusy         AgentStatus = "busy"
	StatusWaitingInput AgentStatus = "waiting_input"
	StatusError        AgentStatus = "error"
	StatusStopped      AgentStatus = "stopped"
)

// AgentEventType classifies what happened.
type AgentEventType string

const (
	EventStarted      AgentEventType = "started"
	EventStopped      AgentEventType = "stopped"
	EventNeedsInput   AgentEventType = "needs_input"
	EventOutputLine   AgentEventType = "output_line"
	EventTaskComplete AgentEventType = "task_complete"
	EventError        AgentEventType = "error"
	EventLog          AgentEventType = "log"
)

// AgentEvent is an immutable event emitted by an agent.
type AgentEvent struct {
	ID          string         `json:"id"`
	AgentID     string         `json:"agent_id"`
	WorkspaceID string         `json:"workspace_id"`
	Type        AgentEventType `json:"type"`
	Message     string         `json:"message"`
	Payload     map[string]any `json:"payload,omitempty"`
	Timestamp   time.Time      `json:"timestamp"`
}

// Agent represents a single AI coding agent instance.
type Agent struct {
	ID          string         `json:"id"`
	Name        string         `json:"name"`
	WorkspaceID string         `json:"workspace_id"`
	Kind        string         `json:"kind"` // claude, goose, aider, custom
	Status      AgentStatus    `json:"status"`
	LastEvent   *AgentEvent    `json:"last_event,omitempty"`
	Events      []AgentEvent   `json:"-"`
	Meta        map[string]any `json:"meta,omitempty"`
}

// Icon returns a single-character status icon for the agent.
func (a *Agent) Icon() string {
	switch a.Status {
	case StatusRunning, StatusBusy:
		return "●"
	case StatusWaitingInput:
		return "▶"
	case StatusIdle:
		return "○"
	case StatusError:
		return "!"
	case StatusStopped:
		return "■"
	default:
		return "?"
	}
}

// ShortStatus returns a short human-readable status string.
func (a *Agent) ShortStatus() string {
	if a.LastEvent != nil && a.LastEvent.Message != "" {
		msg := a.LastEvent.Message
		if len(msg) > 40 {
			msg = msg[:37] + "..."
		}
		return msg
	}
	return string(a.Status)
}
