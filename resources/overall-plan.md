# The Target Architecture: The 4-Tier Agentic Mesh

Before diving into the learning phases, keep this end-state in mind. Every step in this roadmap is building toward this specific, containerized structure.


| Tier                 | Component  | Tech Stack                    | Primary Responsibility                                                                                 |
| -------------------- | ---------- | ----------------------------- | ------------------------------------------------------------------------------------------------------ |
| **1. The Interface** | Web UI     | React, TypeScript, Tailwind   | Consumes JSON APIs and streaming text. Absolutely no AI logic.                                         |
| **2. The Brain**     | Agent API  | Go, Google ADK-Go (or Mastra) | Orchestrates prompts, routes tasks to specialized agents, manages context windows. No database access. |
| **3. The Hands**     | MCP Server | Go, MCP Go-SDK                | Securely executes business logic. Exposes specific functions (tools) to the Agent API via JSON-RPC.    |
| **4. The Memory**    | Data Layer | SQLite, Chromem-go            | Stores relational data (task queues, users) and semantic vector data (long-term agent memory).         |


---

## Phase 1: Rewiring for Go (Weeks 1-2)

Coming from a world of component lifecycles, global state management, and asynchronous promises, Go requires a strict mental shift. The goal here is to learn idiomatic Go to prevent AI coding assistants from injecting anti-patterns.

### Mental Model Translation


| Frontend Paradigm          | Go Equivalent              | Key Difference                                                                                      |
| -------------------------- | -------------------------- | --------------------------------------------------------------------------------------------------- |
| **Promises / Async/Await** | Goroutines & Channels      | Go uses lightweight threads and explicit message passing, rather than an event loop.                |
| **Classes / Prototypes**   | Structs & Receiver Methods | Go favors composition over inheritance. No "extend" keywords.                                       |
| **TypeScript Interfaces**  | Go Interfaces              | Go interfaces are satisfied *implicitly*. If a struct has the methods, it implements the interface. |
| **Try / Catch Blocks**     | Explicit Error Returns     | Errors are values (`if err != nil`). You must handle them immediately.                              |


**Action Items:**

1. **Read the Source:** Consume *Effective Go* and Rob Pike's talks on concurrency. Do not rely on LLMs to explain why Go is structured the way it is.
2. **Deploy the Socratic AI:** Set up Cursor or Claude Code with the "Strict Go Architect" system prompt discussed earlier.
3. **Build a CLI Tool:** Write a pure Go command-line application that reads a local JSON file, parses it into structs, and prints specific values. This forces you to handle strict typing and error checking without the distraction of a web server.

---

## Phase 2: The Go Backend Baseline (Weeks 3-4)

Agents need a fast, reliable network to communicate. Before introducing unpredictable LLMs, you must master Go's native networking capabilities. Do not use massive web frameworks; stick to the standard library.

**Action Items:**

1. **Master `net/http`:** Build a standard REST API that can receive a JSON payload, validate it, and return a 200 OK or 400 Bad Request.
2. **Integrate SQLite:** Connect your API to a local SQLite database using the standard `database/sql` package. Write raw SQL queries to insert and retrieve data.
3. **Implement Middleware:** Write a simple logging middleware function that wraps your HTTP handlers to track incoming requests and execution time. You will need this later to observe agent behavior.

---

## Phase 3: The AI & Memory Foundation (Weeks 5-6)

Now you introduce the AI primitives into your Go code, focusing purely on getting structured data in and out of an LLM.

**Action Items:**

1. **The LLM Connection:** Install Ollama locally. Write a Go function that sends a text string to your local Llama 3 model and prints the response.
2. **Enforce Structured Output:** Modify your prompt to demand a strict JSON response. Use Go's `encoding/json` to unmarshal the AI's response directly into a Go struct. If it fails, write logic that sends the error back to the LLM to fix it.
3. **Embed Vector Memory:** Import `chromem-go`. Take a paragraph of text, convert it into an embedding, and save it. Then, write a search function that queries the database with a slightly different phrase and successfully retrieves the original text.

---

## Phase 4: Building the "Hands" (MCP Server) (Week 7)

You are now building Tier 3 of your architecture. This service exists solely to execute tasks securely.

**Action Items:**

1. **Setup MCP Go-SDK:** Initialize a new Go module dedicated purely to the MCP Server.
2. **Define Your Tools:** Write a Go function that queries your SQLite database (from Phase 2). Expose this function as an "MCP Tool."
3. **Test the Server:** Run the server locally. Ensure it broadcasts its available tools and securely listens for JSON-RPC execution commands.

---

## Phase 5: Building the "Brain" (Agent API) (Week 8)

This is Tier 2. You will now orchestrate the intelligence and connect it to the tools you just built.

**Action Items:**

1. **Setup Google ADK-Go (or Mastra):** Initialize the framework in a new Go module.
2. **Define the Router:** Create a primary agent whose system prompt instructs it to evaluate user queries and determine which tools are needed.
3. **Connect the Mesh:** Configure your Agent API to connect to your running MCP Server as a client.
4. **Execute the Loop:** Send a complex request to the Agent API (e.g., "Summarize the data in my database"). Watch the Agent API request the data from the MCP Server, process it, and return the final answer.

---

## Phase 6: Infrastructure & Deployment (Week 9)

The code works locally; now it needs to be locked down and containerized so the components can talk to each other securely.

**Action Items:**

1. **Write Dockerfiles:** Create lightweight Alpine-based Dockerfiles for your UI, Agent API, and MCP Server. Because Go compiles to a single binary, these images will be incredibly small and fast.
2. **Configure Docker Compose:** Write a `docker-compose.yml` file that defines your four services.
3. **Lock Down the Network:** Configure internal Docker networks so that the Agent API can talk to the UI and the MCP Server, but the MCP Server and SQLite database are entirely blocked from the public host port.

