# CLAUDE.md — agenticitems

## What This Project Is

A learning project building an agentic inventory management system. The dual goal is:
1. Learn idiomatic Go — not just syntax, but how Go architects things (composition over inheritance, errors as values, interfaces satisfied implicitly).
2. Build toward a 4-tier agentic mesh architecture (see `resources/overall-plan.md`).

The user is an experienced React/TypeScript/OOP developer. Frame Go explanations in relation to frontend/TypeScript analogues where helpful.

---

## Project Layout

```
agenticitems/
├── service/          # Go REST API (Tier 3 — "The Hands")
├── agents/           # Go Agent API (Tier 2 — "The Brain")
├── itemui/           # React + TypeScript UI (Tier 1 — "The Interface")
├── lets-go-further/  # HTML book/tutorial reference material
└── resources/        # Architecture plan and notes
```

### Service (`service/`)
- Module: `agenticitemsapi.arryn.net`, Go 1.26+
- Router: `httprouter` (lightweight, no framework)
- Entry: `cmd/api/main.go` — config flags, slog structured logging, net/http server
- Handlers: `cmd/api/items.go`, `healthcheck.go`
- Helpers: `cmd/api/helpers.go` — JSON read/write, envelope type, ID param extraction
- Models: `internal/data/` — Item struct, custom Cents type
- Validation: `internal/validator/`
- Port: 4000

### Agents (`agents/`)
- Module: `agenticitemagents.arryn.net`, Go 1.26+
- Dependency: `google.golang.org/adk` (Google Agents Development Kit)
- Currently a placeholder HTTP server on port 4001
- Will become the LLM orchestration layer

### ItemUI (`itemui/`)
- React 19, TypeScript, Vite
- Dev server port 5173, proxies `/api` to `localhost:4000`
- Axios for HTTP

### Infrastructure
- `docker-compose.yml` at root — PostgreSQL 16, API service, Agents service
- Each Go service has its own Dockerfile (multi-stage, Alpine-based)

---

## Learning Focus & How to Help

The user is learning Go through the "Let's Go Further" book (content in `lets-go-further/`). When helping:

- **Teach the why, not just the what.** Explain Go idioms in relation to the user's TypeScript/React background.
- **Socratic over prescriptive.** When the user is working through a concept, ask what they think before giving the answer.
- **Enforce idiomatic Go.** Push back on anti-patterns — OOP instincts (class hierarchies, global state, exceptions) don't translate. Go prefers flat composition, explicit errors, and simple interfaces.
- **Don't over-engineer.** Go favors simplicity. Resist adding abstractions, helpers, or packages that aren't needed yet.
- **Standard library first.** Avoid reaching for third-party packages when `net/http`, `encoding/json`, `database/sql`, etc. suffice.

Key mental model translations to reinforce:
| TypeScript/React | Go |
|---|---|
| `async/await`, Promises | Goroutines & channels |
| Classes, inheritance | Structs + receiver methods, composition |
| TypeScript interfaces (structural but explicit) | Go interfaces (implicit satisfaction) |
| `try/catch` | `if err != nil` — errors are values |
| Global state / context | Dependency injection via `application` struct |

---

## Current Build Phase

**Phase 2 — Go Backend Baseline** (from `resources/overall-plan.md`):
- [x] Basic REST API structure with `httprouter`
- [x] JSON read/write helpers, envelope pattern, error handling middleware
- [x] Custom types (Cents)
- [x] Input validation package
- [ ] SQLite/PostgreSQL integration via `database/sql`
- [ ] Middleware: logging, rate limiting
- [ ] Full CRUD for items

Next phases: AI primitives (Ollama, structured LLM output), MCP server, Agent API with ADK, containerization.

---

## Code Conventions in This Project

- Structured logging with `log/slog` — not `fmt.Println`
- JSON responses always wrapped in an `envelope` (`map[string]any`) — e.g. `{"item": {...}}`
- Money stored and handled as integer cents (`Cents` type), serialized as `"dollars.cents"` string
- Errors returned to clients via centralized helpers in `errors.go`
- Config passed via CLI flags into a `config` struct held on an `application` struct — no globals
- Routes registered in `routes.go`, handlers in dedicated files per resource

---

## What Not to Do

- Don't add packages or dependencies without discussing it — the learning value is in the standard library
- Don't refactor surrounding code when fixing a specific bug or adding a specific feature
- Don't add docstrings, comments, or type annotations to unchanged code
- Don't suggest framework replacements (e.g. Gin, Echo) — `net/http` + `httprouter` is intentional
- Don't generate URLs or external links unless confident they're valid and relevant
