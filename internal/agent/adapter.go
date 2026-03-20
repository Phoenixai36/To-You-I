// T.U.I — To Unify Imagination
// internal/agent/adapter.go — Agent adapter interface
package agent

import (
	"context"

	"github.com/Phoenixai36/To-You-I/internal/model"
)

// AgentAdapter is the interface every agent integration must implement.
// Each adapter wraps a specific AI agent (Claude, Goose, Aider, custom) and
// translates T.U.I UICommands into the agent's native protocol.
type AgentAdapter interface {
	// Info returns the static metadata for this agent.
	Info() *model.Agent

	// Start boots the agent process / connection.
	Start(ctx context.Context) error

	// Stop gracefully shuts down the agent.
	Stop(ctx context.Context) error

	// SendInput delivers a user input string to the agent.
	SendInput(ctx context.Context, input string) error

	// Events returns a read-only channel of AgentEvent.
	// The adapter pushes events here asynchronously.
	Events() <-chan model.AgentEvent

	// Dispatch executes a UICommand (run_task, open_shell, etc.).
	Dispatch(ctx context.Context, cmd model.UICommand) error
}

// Registry holds all registered adapters keyed by agent ID.
type Registry struct {
	adapters map[string]AgentAdapter
}

// NewRegistry creates an empty adapter registry.
func NewRegistry() *Registry {
	return &Registry{adapters: make(map[string]AgentAdapter)}
}

// Register adds an adapter to the registry.
func (r *Registry) Register(a AgentAdapter) {
	r.adapters[a.Info().ID] = a
}

// Get returns the adapter for a given agent ID.
func (r *Registry) Get(id string) (AgentAdapter, bool) {
	a, ok := r.adapters[id]
	return a, ok
}

// All returns all registered adapters.
func (r *Registry) All() []AgentAdapter {
	out := make([]AgentAdapter, 0, len(r.adapters))
	for _, a := range r.adapters {
		out = append(out, a)
	}
	return out
}
