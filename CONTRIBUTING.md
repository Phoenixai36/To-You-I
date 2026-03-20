# Contributing to T.U.I — To Unify Imagination

Thank you for spending time on T.U.I. This guide covers everything you need to
go from zero to your first merged pull request.

---

## Table of Contents

1. [Code of Conduct](#code-of-conduct)
2. [Getting Started](#getting-started)
3. [Development Workflow](#development-workflow)
4. [Commit Convention](#commit-convention)
5. [Code Style](#code-style)
6. [Testing](#testing)
7. [Opening a Pull Request](#opening-a-pull-request)
8. [Reporting Bugs](#reporting-bugs)
9. [Proposing Features](#proposing-features)

---

## Code of Conduct

Be respectful. Be curious. Be constructive.  
Harassment, gatekeeping, and bad-faith criticism are not welcome here.

---

## Getting Started

### Prerequisites

| Tool | Minimum version |
|------|----------------|
| Go | 1.22 |
| Git | 2.40 |
| golangci-lint | 1.57 |
| make | any |

### Fork & Clone

```bash
git clone https://github.com/<your-handle>/To-You-I.git
cd To-You-I
git remote add upstream https://github.com/Phoenixai36/To-You-I.git
```

### Install Dependencies

```bash
go mod download
```

### Build & Run

```bash
make build          # compile to ./dist/tuicortex
make run            # build and launch the TUI
make run-dev        # hot-reload via go run (no rebuild step)
```

---

## Development Workflow

```
main            ─ stable; protected; CI must pass
develop         ─ integration branch (optional)
feat/<name>     ─ new features branched from main
fix/<name>      ─ bug fixes
chore/<name>    ─ tooling, deps, refactors
```

1. Branch from `main`:  
   `git checkout -b feat/my-feature`
2. Make your changes.
3. Run `make test` and `make lint` locally.
4. Commit using the convention below.
5. Push and open a PR against `main`.

---

## Commit Convention

T.U.I follows [Conventional Commits](https://www.conventionalcommits.org/).

```
<type>(<scope>): <short summary>

[optional body]
[optional footer]
```

| Type | When to use |
|------|-------------|
| `feat` | New feature visible to users |
| `fix` | Bug fix |
| `docs` | Documentation only |
| `style` | Formatting, no logic change |
| `refactor` | Code restructure, no feature change |
| `perf` | Performance improvement |
| `test` | Adding or fixing tests |
| `chore` | Tooling, dependencies, CI |
| `ci` | CI/CD config changes |
| `build` | Build system changes |

Scopes: `ui`, `model`, `agent`, `server`, `config`, `cmd`, `docs`, `ci`

Examples:

```
feat(ui): add glitch renderer with chroma-shift
fix(server): close SSE connection on client disconnect
docs(readme): update installation section
```

---

## Code Style

- Run `gofmt` / `goimports` before committing.
- All exported symbols must have a doc-comment.
- Keep functions short (≤40 lines preferred).
- Error strings must be lowercase and not end with punctuation.
- Avoid `panic` outside of test helpers.
- Lip Gloss styles live in `internal/ui/styles.go`; don’t scatter inline style definitions.

```bash
make lint    # golangci-lint with project config
```

---

## Testing

```bash
make test              # run all tests with race detector
go test ./... -v       # verbose
go test ./internal/ui  # single package
```

- Every new package should have at least one `_test.go` file.
- Table-driven tests are preferred for functions with multiple cases.
- Use `t.Helper()` in test utilities.
- Mock external I/O with interfaces, not concrete types.

---

## Opening a Pull Request

1. Ensure `make test` and `make lint` pass with zero errors.
2. Fill in the PR template (title, description, related issues).
3. Link any relevant issue: `Closes #42`.
4. Request a review from a maintainer if unsure who to tag.
5. Be responsive to review feedback — mark conversations resolved after addressing them.

---

## Reporting Bugs

Open an issue and include:

- **T.U.I version** (`tuicortex --version`)
- **OS + terminal emulator** (e.g. Windows 11, WezTerm 20240203)
- **Steps to reproduce** (minimal)
- **Expected vs. actual behavior**
- **Relevant logs** (run with `TUI_LOG_LEVEL=debug`)

---

## Proposing Features

Open a GitHub Discussion or an issue tagged `enhancement` and describe:

- The problem you’re solving.
- Your proposed solution.
- Alternatives you’ve considered.

Large features should be discussed _before_ implementation to avoid wasted effort.

---

_Built with 🔥 by the Dark Phoenix collective._
