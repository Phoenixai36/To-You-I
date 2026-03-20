# Architecture — T.U.I (To Unify Imagination)

This document outlines the high-level architecture of the T.U.I project,
describing its components, their interactions, and design principles.

---

## Table of Contents

1. [Overview](#overview)
2. [Component Diagram](#component-diagram)
3. [Package Structure](#package-structure)
4. [Core Components](#core-components)
   - [cmd/tuicortex](#cmdtuicortex)
   - [internal/ui](#internalui)
   - [internal/model](#internalmodel)
   - [internal/agent](#internalagent)
   - [internal/server](#internalserver)
   - [internal/config](#internalconfig)
5. [Data Flow](#data-flow)
6. [Concurrency Model](#concurrency-model)
7. [Extension Points](#extension-points)

---

## Overview

T.U.I is a **Terminal User Interface (TUI)** designed for managing AI agents,
workspaces, and shells with a glitch-professional "Dark Phoenix" aesthetic.

It combines:
- **Bubble Tea** for reactive TUI rendering
- **Lip Gloss** for composable styling
- **HTTP + SSE** for real-time event streaming from external agents
- **Go channels** for internal event bus communication

---

## Component Diagram

```
┌────────────────────────────────────────────────────────┐
│                  cmd/tuicortex                         │
│  ┌──────────────────────────────────────────────────┐ │
│  │ main.go: parse flags, init config, start server  │ │
│  │ & launch TUI program                             │ │
│  └──────────────────────────────────────────────────┘ │
└────────────────────────────────────────────────────────┘
                         │
         ┌───────────────┼───────────────┐
         ▼               ▼               ▼
 ┌─────────────┐ ┌─────────────┐ ┌──────────────┐
 │ internal/ui │ │ internal/   │ │ internal/    │
 │             │ │ server      │ │ config       │
 │  Bubble Tea │ │ (HTTP+SSE)  │ │ (XDG TOML)   │
 │  App Model  │ │             │ │              │
 └──────┬──────┘ └──────┬──────┘ └──────┬───────┘
        │               │               │
        │  ┌────────────┴────────┐      │
        │  │    Event Bus        │◄─────┘
        │  │ (agent → TUI push)  │
        │  └────────────┬────────┘
        │               │
        ▼               ▼
 ┌──────────────────────────────┐
 │      internal/model          │
 │ (Workspace, Agent, Command)  │
 └──────────────────────────────┘
```

---

## Package Structure

```
To-You-I/
├─ cmd/tuicortex/
│  └─ main.go                    # entry point
├─ internal/
│  ├─ agent/
│  │  └─ adapter.go              # AgentAdapter interface + Registry
│  ├─ config/
│  │  └─ config.go               # XDG-aware config management
│  ├─ model/
│  │  ├─ agent.go                # Agent domain type
│  │  ├─ command.go              # UICommand type
│  │  └─ workspace.go            # Workspace domain type
│  ├─ server/
│  │  └─ server.go               # HTTP + SSE event stream
│  └─ ui/
│     ├─ app.go                  # root Bubble Tea App model
│     ├─ glitch.go               # glitch effect renderer
│     ├─ palette.go              # command palette overlay
│     ├─ styles.go               # Dark Phoenix Lip Gloss theme
│     └─ workspacepanel.go       # workspace switcher rail
├─ .github/workflows/ci.yml
├─ .gitignore
├─ CONTRIBUTING.md
├─ LICENSE
├─ Makefile
├─ README.md
├─ go.mod
└─ go.sum
```

---

## Core Components

### `cmd/tuicortex`

**Responsibility**: Parse CLI flags, load config, start HTTP server, launch TUI.

- Reads `~/.config/tui/config.toml` (or env overrides).
- Spawns the SSE server in a goroutine.
- Initializes the root Bubble Tea program (`ui.App`) and runs it.

### `internal/ui`

**Responsibility**: All TUI rendering logic using Bubble Tea + Lip Gloss.

- **`app.go`**: Root model that orchestrates keyboard navigation, workspace rail, and command palette.
- **`styles.go`**: Centralized Dark Phoenix theme (colors, borders, spacing).
- **`glitch.go`**: Optional scanline glitch effect (chroma-shift, horizontal displacement).
- **`palette.go`**: Fuzzy command palette overlay (Ctrl+P).
- **`workspacepanel.go`**: Left rail showing workspaces, agent counts, status dots.

### `internal/model`

**Responsibility**: Domain types shared across the application.

- **`Workspace`**: Groups agents and tasks. Tracks status (idle/busy/error).
- **`Agent`**: Represents an AI agent with name, status, events, icon.
- **`UICommand`**: An action invokable from the command palette.

### `internal/agent`

**Responsibility**: Agent adapter registry and lifecycle management.

- **`AgentAdapter`**: Interface for plugging in external agent frameworks (Aider, Goose, OpenCode, etc.).
- **`Registry`**: Maps agent IDs to running adapters; handles start/stop/query.

### `internal/server`

**Responsibility**: Expose HTTP + SSE endpoint for agent→TUI events.

- Listens on `127.0.0.1:7788` by default.
- Agents POST JSON events → server broadcasts them via SSE.
- The TUI subscribes to SSE and updates its model on incoming events.

### `internal/config`

**Responsibility**: Load and validate configuration.

- Reads TOML from XDG config directory or fallback.
- Merges environment variable overrides (`TUI_SERVER_ADDR`, `TUI_GLITCH`, etc.).
- Validates all fields (e.g., port in range, log level is valid).

---

## Data Flow

```
┌─────────────┐
│  External   │
│   Agent     │  (Aider, Goose, custom script)
└──────┬──────┘
       │ POST /events (JSON)
       ▼
┌────────────────┐
│  HTTP Server   │  (internal/server)
│  (SSE Broker)  │
└────────┬───────┘
         │ SSE stream
         ▼
  ┌─────────────┐
  │ Bubble Tea  │  (internal/ui/app.go)
  │  App.Update │  ← receives EventMsg
  └──────┬──────┘
         │ model update
         ▼
  ┌──────────────┐
  │ Workspace /  │  (internal/model)
  │ Agent state  │
  └──────────────┘
```

---

## Concurrency Model

- **Single TUI goroutine**: Bubble Tea's event loop runs on the main goroutine.
- **HTTP server goroutine**: Listens and handles requests concurrently.
- **SSE broadcaster**: Uses `sync.Mutex` around subscriber map to safely publish events.
- **Agent adapters**: Each adapter may spawn goroutines for I/O, but communicates via channels.

**Key invariant**: Only the TUI's `Update` method mutates model state; all external events are sent as `tea.Msg` via `p.Send(msg)`.

---

## Extension Points

To add a new agent type:

1. Implement `AgentAdapter` interface (`Start`, `Stop`, `GetEvents`).
2. Register it in `agent.Registry`.
3. POST events to `/events` with `{"agent_id": "...", "type": "...", "payload": {...}}`.

To add a new UI panel:

1. Create a new Bubble Tea model in `internal/ui/`.
2. Embed it in `App` model.
3. Update `App.Update` to route keyboard events.
4. Render it in `App.View`.

---

_Built with 🔥 by the Dark Phoenix collective._
