---
name: mcp-go
description: How to create an MCP (Model Context Protocol) server in Golang using the official Go SDK.
---

<mcp_server_go name="How to create an Model Context Protocol server in Go">
  <goal>How to build an MCP server in Go using the official Go SDK (github.com/modelcontextprotocol/go-sdk)</goal>

  <reference_implementation name="Recommended Reference layout">
    1. cmd/{app}/main.go
      1) process entrypoint
      2) logging initialization
      3) config loading
      4) dependency graph wiring
      5) server start (transport selection)
    2. internal/server
      1) MCP server wrapper (construction, shared helpers)
      2) interfaces for tool registrators
    3. internal/tools/{feature}
      1) tool registration (mcp.AddTool)
      2) tool input/output DTOs for schema inference
      3) request validation and mapping
    4. internal/adapters/{upstream}
      1) HTTP clients or other integrations
      2) translation between upstream DTOs and internal domain models
    5. internal/domain
      1) cross-layer models and options
      2) cross-layer errors (including upstream error wrappers)
    Notes:
      1. Keep DTOs unexported (lowercase names) unless you have a strong reason to publish them.
      2. Avoid "utils" packages without ownership; prefer small helper packages with clear scope.
  </reference_implementation>

  <server_bootstrap name="Server bootstrap">
    <implementation_identity>
      1. Set server implementation identity:
        1) Name: stable identifier (example: "xxxx-mcp")
        2) Version: build-time version string
        3) Title: short human-friendly name
      2. SHOULD provide a flag to print version info (example: -version).
      3. SHOULD inject build metadata via ldflags (version, commit, date, builtBy).
    </implementation_identity>
    <server_creation>
      Recommended pattern:
        1. Create the server via mcp.NewServer(implementation, options)
        2. Provide a short, self-contained Instructions/system prompt via ServerOptions.Instructions
        3. Register all tools BEFORE running the server (capabilities are inferred from registered features)
    </server_creation>
    <shutdown>
      1. Use context cancellation for graceful shutdown.
      2. SHOULD use signal.NotifyContext for SIGINT/SIGTERM.
    </shutdown>
    <transport>
      1. Stdio transport (recommended for CLI-style MCP servers): `server.Run(ctx, &mcp.StdioTransport{})`
    </transport>
  </server_bootstrap>

  <tool_design name="Tool design">
    <tool_naming>
      1. Use stable, explicit tool names.
      2. Use consistent prefixes per feature/service (example: wiki_* and tracker_*).
      3. Keep names and descriptions short and self-contained.
      4. If a field has predefined values, document ALL possible values.
      5. DON'T use vague wording like "etc.", "and so on", "and more", "such as".
    </tool_naming>
    <tool_registration>
      Prefer typed tool registration (mcp.AddTool / mcp.AddTool[In,Out]) over raw Server.AddTool.
      Recommended pattern: define a per-feature "registrator" that receives dependencies and registers tools.
      Benefits:
        1. keeps server wiring clean
        2. supports per-feature enablement
        3. makes tests straightforward
    </tool_registration>
    <typed_handlers name="Typed tool handlers (preferred)">
      Why:
        1. automatic JSON Schema inference for input/output
        2. argument validation
        3. marshaling/unmarshaling handled by SDK
        4. consistent error packaging
      Preferred handler shape (from go-sdk docs): `func(ctx context.Context, req *mcp.CallToolRequest, in In) (*mcp.CallToolResult, Out, error)`
      Practical guidance:
        1. under normal circumstances you can ignore req and return a nil *mcp.CallToolResult
        2. keep In/Out as structs (or pointers to structs) with json tags
        3. use jsonschema struct tags for argument descriptions
    </typed_handlers>
    <dto_schema name="DTOs and schema inference">
      Input DTO rules:
        1. Use a struct with json tags
        2. SHOULD include jsonschema tags for field descriptions
        3. Make required fields non-omitempty (and validate in code)
        4. Validate ranges and bounds in code (page_size, limits, etc.)
      Output DTO rules:
        1. Be JSON-marshalable
        2. SHOULD omit optional data via omitempty
        3. SHOULD prefer explicit sub-structures over map[string]any for stable outputs
    </dto_schema>
    <input_validation name="Input validation">
      1. Validate required strings (non-empty)
      2. Validate numeric bounds (>= 0, max limits)
      3. Validate enums when the tool exposes a closed set of options
      4. Return a safe error message for invalid inputs
    </input_validation>
  </tool_design>

  <mcp_features name="Other MCP features (optional)">
    <prompts>
      1. Prompts are reusable message templates exposed by the server.
      2. WHEN you add at least one prompt before connecting a client, THE server SHALL advertise the prompts capability.
      3. Use Server.AddPrompt(prompt, handler).
    </prompts>
    <resources>
      1. Resources are data addressed by URI.
      2. Use Server.AddResource and Server.AddResourceTemplate.
      3. Return ResourceNotFound (or an equivalent safe error) when a resource URI does not exist.
      4. SHOULD support list-changed notifications only when you truly add/remove resources dynamically.
    </resources>
    <utilities>
      1. Completion: enable by setting ServerOptions.CompletionHandler.
      2. Logging to client: use mcp.NewLoggingHandler(serverSession, opts) to obtain an slog.Handler.
      3. Cancellation: tool handlers MUST respect ctx.Done() and terminate promptly.
      4. Pagination: server-side pagination is enabled by default; you MAY tune ServerOptions.PageSize.
    </utilities>
  </mcp_features>

  <error_handling name="Error handling">
    <principles>
      1. DON'T return raw upstream errors directly.
      2. Sanitize upstream response bodies before logging.
      3. Return safe, short errors suitable for end-users/LLMs.
    </principles>
    <recommended_pattern>
      1. Adapters create a typed upstream error with:
         1) service name
         2) operation name
         3) HTTP status
         4) safe short message
         5) optional sanitized details for logs
      2. Tools convert errors to safe tool errors and log them.
    </recommended_pattern>
  </error_handling>

  <config name="Configuration">
    1. Load config in one place (package config) and validate it.
    2. Document env vars in .env.example.
    3. Keep env vars consistent across: .env.example, .env, Taskfile.yml, code, and documentation.
  </config>

  <logging name="Logging">
    1. Use log/slog structured logging.
    2. Set a default logger at startup.
    3. DON'T log secrets.

    Suggested levels:
      1. Info: startup, shutdown, major lifecycle events
      2. Error: failures; include operation context
  </logging>

  <testing name="Testing">
    1. Test that tools are registered (names are present).
    2. Test tool handlers with adapter mocks.
    3. Use go.uber.org/mock for mocks (go:generate).
    4. Use in-memory transports for MCP integration-like tests.
  </testing>

  <documentation name="Documentation">
    1. Keep tool documentation in docs/ and keep it aligned with code.
    2. DON'T use tables in user-facing docs; prefer sections and lists.
  </documentation>

  <security name="Security">
    1. Treat all inbound tool inputs as untrusted.
    2. Avoid echoing back sensitive values.
    3. If exposing an HTTP transport publicly:
      1) Add authentication (bearer token / OAuth)
      2) SHOULD follow go-sdk auth middleware guidance (RequireBearerToken)
  </security>

  <folder_structure name="Recommended folder structure">
```txt
├── Taskfile.yml # use taskfile for common commands: lint, test, fmt, vet, generate, build, etc.
├── .env.example # document all env vars here; keep in sync with .env and docs
├── README.md
├── go.mod
├── cmd
│   └── {name}-mcp
│       └── main.go
├── internal
│   ├── adapters # upstreams, wrapper for filesystem, etc.
│   ├── appinit # dependency building
│   ├── config # config loading and validation
│   ├── domain # cross-cutting models and errors
│   ├── server
│   │   ├── consts.go # some internal constants
│   │   ├── doc.go # package docs
│   │   ├── dto.go # tools input/output DTOs
│   │   ├── interfaces.go # interfaces, that implemented by usecase
│   │   ├── interfaces_mock.go # generated by go:generate, used in server tests
│   │   ├── models.go # some internal models
│   │   ├── service.go # mcp server construction and tools registration
│   │   └── service_test.go # server tests
│   └── usecases
│       ├── {some usecase}
│       │    ├── consts.go # some internal constants
│       │    ├── doc.go # package docs
│       │    ├── interfaces.go # interfaces, that implemented by adapters
│       │    ├── interfaces_mock.go # generated by go:generate, used in usecase tests
│       │    ├── models.go # some internal models
│       │    ├── service.go # usecase service implementation
│       │    └── service_test.go # usecase service tests
│       └── ...
```
  </folder_structure>

  <examples name="Examples">
    <example name="Minimal stdio server (typed tool + graceful shutdown)">
      ```go
      package main

      import (
        "context"
        "errors"
        "log"
        "os/signal"
        "syscall"

        "github.com/modelcontextprotocol/go-sdk/mcp"
      )

      type GreetIn struct {
        Name string `json:"name" jsonschema:"name to greet. Required"`
      }

      type GreetOut struct {
        Greeting string `json:"greeting" jsonschema:"greeting text"`
      }

      func greet(_ context.Context, _ *mcp.CallToolRequest, in GreetIn) (*mcp.CallToolResult, GreetOut, error) {
        if in.Name == "" {
          return nil, GreetOut{}, errors.New("name is required")
        }
        return nil, GreetOut{Greeting: "Hi " + in.Name}, nil
      }

      func main() {
        srv := mcp.NewServer(&mcp.Implementation{Name: "example-mcp", Version: "v0.1.0"}, nil)
        mcp.AddTool(srv, &mcp.Tool{Name: "greet", Description: "Greets the user"}, greet)

        ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
        defer cancel()

        if err := srv.Run(ctx, &mcp.StdioTransport{}); err != nil {
          log.Fatal(err)
        }
      }
      ```
    </example>

    <example name="Typed DTOs with jsonschema tags">
      ```go
      type SearchIn struct {
        Query   string `json:"query" jsonschema:"search query. Required"`
        Page    int    `json:"page,omitempty" jsonschema:"page number. Minimum 1"`
        PerPage int    `json:"per_page,omitempty" jsonschema:"items per page. Valid range: 1-50"`
      }

      type SearchOut struct {
        Items      []any  `json:"items"`
        NextCursor string `json:"next_cursor,omitempty"`
      }
      ```
    </example>

    <example name="Registrator pattern (feature module tool registration)">
      ```go
      // file: internal/server/interfaces.go (consumer package)
      package server

      import "github.com/modelcontextprotocol/go-sdk/mcp"

      // IToolsRegistrator is defined where it is consumed (server wiring).
      type IToolsRegistrator interface {
        Register(srv *mcp.Server) error
      }
      ```

      ```go
      // file: internal/tools/feature/service.go (implementation package)
      package feature

      import (
        "context"

        "github.com/modelcontextprotocol/go-sdk/mcp"

        // import path depends on your module name
        "path/to/internal/server"
      )

      type Registrator struct {
        // deps...
      }

      var _ server.IToolsRegistrator = (*Registrator)(nil)

      func (r *Registrator) Register(srv *mcp.Server) error {
        mcp.AddTool(srv, &mcp.Tool{Name: "feature_do", Description: "Does the thing"}, r.doTool)
        return nil
      }

      func (r *Registrator) doTool(ctx context.Context, _ *mcp.CallToolRequest, in SearchIn) (*mcp.CallToolResult, SearchOut, error) {
        // validate input; call deps; map errors
        return nil, SearchOut{}, nil
      }
      ```
    </example>

    <example name="Input validation helper">
      ```go
      // imports: errors, fmt, strings
      func validateSearchIn(in SearchIn) error {
        if strings.TrimSpace(in.Query) == "" {
          return errors.New("query is required")
        }
        if in.Page < 0 {
          return errors.New("page must be non-negative")
        }
        if in.PerPage < 0 || in.PerPage > 50 {
          return fmt.Errorf("per_page must be in range 0-50")
        }
        return nil
      }
      ```
    </example>

    <example name="Safe upstream error mapping (no secret leakage)">
      ```go
      // imports: errors, fmt
      type UpstreamError struct {
        Service    string
        Operation  string
        HTTPStatus int
        Message    string
      }

      func (e UpstreamError) Error() string {
        return fmt.Sprintf("%s %s: %s (HTTP %d)", e.Service, e.Operation, e.Message, e.HTTPStatus)
      }

      func toSafeError(err error) error {
        var up UpstreamError
        if errors.As(err, &up) {
          return up
        }
        return errors.New("internal error")
      }
      ```
    </example>

    <example name="In-memory transport test (tools are registered)">
      ```go
      // imports: context, testing, github.com/modelcontextprotocol/go-sdk/mcp, github.com/stretchr/testify/require
      func TestToolsRegistered(t *testing.T) {
        t.Parallel()

        srv := mcp.NewServer(&mcp.Implementation{Name: "server", Version: "v0.1.0"}, nil)
        mcp.AddTool(srv, &mcp.Tool{Name: "greet"}, func(ctx context.Context, _ *mcp.CallToolRequest, in GreetIn) (*mcp.CallToolResult, GreetOut, error) {
          return nil, GreetOut{Greeting: "Hi " + in.Name}, nil
        })

        client := mcp.NewClient(&mcp.Implementation{Name: "client", Version: "v0.1.0"}, nil)
        t1, t2 := mcp.NewInMemoryTransports()

        _, err := srv.Connect(context.Background(), t1, nil)
        require.NoError(t, err)

        session, err := client.Connect(context.Background(), t2, nil)
        require.NoError(t, err)
        t.Cleanup(func() { _ = session.Close() })

        var names []string
        for tool, err := range session.Tools(context.Background(), nil) {
          require.NoError(t, err)
          names = append(names, tool.Name)
        }
        require.Contains(t, names, "greet")
      }
      ```
    </example>

    <example name="Streamable HTTP transport (optional)">
      ```go
      // imports: log, net/http, github.com/modelcontextprotocol/go-sdk/mcp
      server := mcp.NewServer(&mcp.Implementation{Name: "server", Version: "v0.1.0"}, nil)
      handler := mcp.NewStreamableHTTPHandler(
        func(r *http.Request) *mcp.Server { return server },
        &mcp.StreamableHTTPOptions{},
      )

      httpServer := &http.Server{Addr: ":8080", Handler: handler}
      if err := httpServer.ListenAndServe(); err != nil {
        log.Fatal(err)
      }
      ```
    </example>
  </examples>

  <checklist name="Definition of done">
    1. The server starts and runs over the chosen transport.
    2. Tools are registered and listed correctly.
    3. Tool schemas are inferred (or provided) and inputs are validated.
    4. Errors are safe for end-users and detailed context is logged server-side.
  </checklist>
</mcp_server_go>