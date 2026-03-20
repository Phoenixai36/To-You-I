<!-- T.U.I — To Unify Imagination -->
<!-- Dark Phoenix · Glitch Professional · by Phoenixai36 -->

<div align="center">

# 🔥 To:You&I

### *T.U.I — To Unify Imagination*

> *"To: You & I — a shared space where you and your agents think together."*

[![Go Version](https://img.shields.io/badge/Go-1.22+-00ADD8?style=for-the-badge&logo=go&logoColor=white)](https://golang.org)
[![Bubble Tea](https://img.shields.io/badge/Bubble_Tea-TUI_Framework-FF6E6E?style=for-the-badge)](https://github.com/charmbracelet/bubbletea)
[![License: MIT](https://img.shields.io/badge/License-MIT-E8890C?style=for-the-badge)](LICENSE)
[![Status](https://img.shields.io/badge/Status-WIP_🔥-FF4500?style=for-the-badge)](#)
[![Platform](https://img.shields.io/badge/Platform-Windows%20%7C%20Linux%20%7C%20macOS-8B0000?style=for-the-badge)](#)

</div>

---

## What is T.U.I?

**To:You&I** is a glitch-professional **terminal HUD** for orchestrating AI coding agents, workspaces and shells — all from a single, unified cockpit inside your terminal.

Built with **Go + Bubble Tea + Lip Gloss**, designed for developers who run multiple AI agents (Claude, Goose, Aider, OpenCode…) across multiple projects simultaneously and need a clean, fast, keyboard-driven control room.

```
┌──────────────────────────────────────────────────────────────┐
│ ● Backend API ⟳ tests │ ○ Frontend Web │ ! Infra │
├───────────────┬──────────────────────────────────────────────┤
│ AGENTS │ ACTIVITY FEED │
│ │ │
│ ▶ Claude ✦ │ 01:23 Claude needs_input: confirm refactor │
│ ● Linter ● │ 01:22 Tests passed (54 suites, 0 failures) │
│ ○ Tests ○ │ 01:21 Linter: 3 warnings in /internal/foo │
│ ! Deploy ! │ 01:20 Deploy: error — staging unreachable │
│ │ │
├───────────────┴──────────────────────────────────────────────┤
│ Ctrl+P Commands · F1 Help · Esc Back · [backend-api] │
└──────────────────────────────────────────────────────────────┘
```

---

## Features

- **Workspace Rail** — switch between project workspaces instantly (`← →`)
- **Agent Sidebar** — see all agents and their live status at a glance
- **Activity Feed** — real-time timeline of agent events and decisions
- **Agent Detail Panel** — review and respond to agent prompts without leaving the terminal
- **Command Palette** — `Ctrl+P` to run tasks, open shells, ask agents
- **Shell Agnostic** — works with PowerShell, pwsh, zsh, bash, fish
- **WezTerm Native** — first-class support for WezTerm workspaces and panes
- **Dark Phoenix Glitch Aesthetic** — because your terminal should look as good as your code

---

## Architecture

```
To-You-I/
├── cmd/
│   └── tuicortex/          # Binary entrypoint
│       └── main.go
├── internal/
│   ├── model/              # Core domain types
│   │   ├── agent.go        # Agent, AgentStatus, AgentEvent
│   │   ├── workspace.go    # Workspace, WorkspaceEvent
│   │   └── command.go      # UICommand types
│   ├── ui/                 # Bubble Tea views
│   │   ├── app.go          # Root model (tea.Model)
│   │   ├── rail.go         # Workspace rail (top bar)
│   │   ├── sidebar.go      # Agent list (left panel)
│   │   ├── feed.go         # Activity feed (right panel)
│   │   ├── detail.go       # Agent detail panel
│   │   ├── palette.go      # Command palette (Ctrl+P)
│   │   └── styles.go       # Lip Gloss Dark Phoenix theme
│   ├── agent/              # Agent adapters
│   │   ├── adapter.go      # Interface: AgentAdapter
│   │   ├── claude.go       # Claude/OpenCode adapter
│   │   ├── goose.go        # Goose adapter
│   │   └── aider.go        # Aider adapter
│   ├── workspace/          # Workspace management
│   │   ├── manager.go      # WorkspaceManager
│   │   └── wezterm.go      # WezTerm integration
│   └── server/             # Local event server
│       ├── server.go       # HTTP/WebSocket event bus
│       └── handlers.go     # REST handlers for agent events
├── docs/
│   ├── ARCHITECTURE.md     # Deep-dive architecture doc
│   ├── ADAPTERS.md         # How to write your own agent adapter
│   └── KEYBINDINGS.md      # Full keyboard reference
├── assets/
│   └── banner.svg          # Dark Phoenix banner
├── .env.example            # Config template
├── Makefile                # Build, run, test targets
├── go.mod
└── README.md
```

---

## Keybindings

| Key | Action |
|-----|--------|
| `← →` | Switch workspace |
| `↑ ↓` | Select agent |
| `Enter` | Open agent detail / focus workspace |
| `i` | Send input to agent |
| `l` | View agent logs |
| `r` | Restart agent task |
| `Ctrl+P` | Open command palette |
| `Ctrl+W` | Close current panel |
| `F1` | Help |
| `q` / `Esc` | Back / Quit |

---

## Stack

| Layer | Technology |
|-------|------------|
| Language | Go 1.22+ |
| TUI Framework | [Bubble Tea](https://github.com/charmbracelet/bubbletea) |
| Styling | [Lip Gloss](https://github.com/charmbracelet/lipgloss) |
| Components | [Bubbles](https://github.com/charmbracelet/bubbles) |
| Terminal Host | [WezTerm](https://wezfurlong.org/wezterm/) |
| Shell Support | pwsh, zsh, bash, fish |
| Agent Protocol | HTTP/WebSocket (local) |

---

## Quick Start

```bash
# Clone
git clone https://github.com/Phoenixai36/To-You-I.git
cd To-You-I

# Install dependencies
go mod tidy

# Run
make run

# Or directly
go run ./cmd/tuicortex/
```

---

## Configuration

Copy `.env.example` to `.env` and adjust:

```env
# Server
TUI_PORT=7331
TUI_HOST=localhost

# Default workspace
TUI_DEFAULT_WORKSPACE=main

# WezTerm integration
TUI_WEZTERM=true

# Theme
TUI_THEME=dark-phoenix
```

---

## Roadmap

- [x] Repo structure & architecture design
- [ ] Core domain model (Agent, Workspace, Event)
- [ ] Bubble Tea root app + layout
- [ ] Workspace rail component
- [ ] Agent sidebar component
- [ ] Activity feed component
- [ ] Agent detail + input panel
- [ ] Command palette
- [ ] Local HTTP event server
- [ ] Claude/OpenCode adapter
- [ ] Goose adapter
- [ ] Aider adapter
- [ ] WezTerm integration
- [ ] Dark Phoenix Lip Gloss theme
- [ ] Docs: ARCHITECTURE.md, ADAPTERS.md, KEYBINDINGS.md
- [ ] v0.1.0 release

---

## Contributing

PRs welcome. Please read [docs/ARCHITECTURE.md](docs/ARCHITECTURE.md) first to understand the design principles.

---

## Author

**Phoenixai36** — AI · Music · Code · Barcelona

> *T.U.I — To: You & I*
> *"a shared space where you and your agents think together."*

---

<div align="center">

[![GitHub](https://img.shields.io/badge/GitHub-Phoenixai36-E8890C?style=for-the-badge&logo=github)](https://github.com/Phoenixai36)

*Built with 🔥 in Barcelona*

</div>
