# Example Pi configuration for use with the [pi-agent-suite](https://github.com/n-r-w/pi-agent-suite) extension package

## Installation

1. Install [pi agent](https://pi.dev/)
2. Copy the contents of this repository's `agent` directory to `~/.pi/agent`.
3. Configure `~/.pi/agent/agent-suite/model-aliases/config.json` for the models you use. By default, the configuration uses `openai-codex/gpt-5.6-xxx` models. The `agent/agent-suite/model-aliases/examples/zai.json` file provides an example configuration for a GLM Coding Plan subscription (requires ZAI_API_KEY env var). See [Model alias recommendations](#model-alias-recommendations).
4. Install [pi-agent-suite](https://github.com/n-r-w/pi-agent-suite) by running `pi install npm:pi-agent-suite`.
5. Run `pi`.
6. Select a main agent with `/agents` or `Ctrl+Option+A`, for example `Auto`.
7. Ask the agent about the available tools and workflows.

## Additional MCP servers

Configure additional MCP servers in `~/.pi/agent/agent-suite/mcp-wrapper/config.json`. The following servers are recommended. The local executable paths in the source configuration are development-machine paths and must not be copied. Install each server by following its linked repository instructions, then configure the installed executable name or absolute path.

- **Perplexity** — provides current web search and research through the Sonar API. Create a Perplexity API key and expose it to the Pi process as `PERPLEXITY_API_KEY`. The `npx -y server-perplexity-ask` command downloads and starts the server when needed.
- **[Team MCP](https://github.com/n-r-w/team-mcp#installation)** — provides shared desks, topics, and messages for multi-agent coordination. Install it from a GitHub release, with `brew install n-r-w/homebrew-tap/team-mcp`, or from source as documented in the repository. Confirm that `team-mcp` is executable before enabling this entry. Configure `TEAM_MCP_MESSAGE_DIR` with a persistent absolute directory if collaboration data must survive operating-system temporary-directory cleanup.
- **[Asteria](https://github.com/n-r-w/asteria#asteria-installation)** — provides symbolic code search through language servers. Install it from a GitHub release, with `brew install --cask n-r-w/homebrew-tap/asteria-mcp`, or from source as documented in the repository. Confirm that `asteria-mcp` is executable before enabling this entry. Asteria does not install language servers. Install the servers required for the languages you use and ensure their executables are in the `PATH` inherited by Pi. Asteria supports `gopls`, `typescript-language-server`, `marksman`, `basedpyright-langserver`, `clangd`, `phpactor`, and `rust-analyzer`; see [Language server installation](https://github.com/n-r-w/asteria#language-server-installation).

Example `mcpServers` entries:

```json
{
  "mcpServers": {
    "websearch": {
      "command": "npx",
      "args": ["-y", "server-perplexity-ask"],
      "onDemand": {
        "name": "web-search-and-query",
        "description": "Search the web and get answers to user questions"
      }
    },
    "team": {
      "type": "stdio",
      "command": "team-mcp",
      "env": {
        "TEAM_MCP_MESSAGE_DIR": "/absolute/path/to/team-mcp-data"
      },
      "onDemand": {
        "name": "team-collaboration-desk",
        "description": "Common communication space between agents"
      }
    },
    "asteria": {
      "type": "stdio",
      "command": "asteria-mcp",
      "onDemand": {
        "name": "asteria-symbolic-tools",
        "description": "Symbolic code search through language servers"
      }
    }
  }
}
```

Use an absolute executable path in `command` when the executable is not in the `PATH` inherited by Pi. The MCP wrapper uses `command`, `env`, and other values literally; it does not expand `${VAR}` or `$env:VAR` placeholders.

MCP server configuration registers tools, but each agent also has a tool allowlist. Add the generated MCP tool names to the `tools` field of every agent that must use them under `~/.pi/agent/agent-suite/agent-selection/agents`. For the server names in the example above, use:

- `websearch_perplexity_ask` for Perplexity.
- `team_*` for the main orchestrator. For subagents, prefer `team_topic_list`, `team_message_create`, `team_message_get`, `team_message_list`, and `team_save_message_to_file`; omitting desk and topic creation avoids coordination conflicts.
- `asteria_*` for Asteria.
- `activate_toolset` for agents that use any server configured with `onDemand`. Without this tool, the agent cannot activate the deferred MCP toolset.

See the agent definitions in `agent/agent-suite/agent-selection/agents` for complete examples. Generated tool names depend on the MCP server key and tool name, so update the allowlists if you rename a server entry. Run `/mcp-refresh` after changing MCP server definitions to rebuild the tool catalog and reload Pi.

## Model alias recommendations

Use models with reliable instruction following and tool calling for every agent alias. Keep comparable context-window sizes across the regular, complex, and critical tiers so that selecting a stronger tier does not reduce the amount of task context available.

| Alias | Purpose | Recommended model profile |
| --- | --- | --- |
| `coding_regular` | Regular coding tasks | Use the fastest low-cost coding model that can reliably edit code, use tools, and run tests. Optimize for latency and cost. |
| `coding_complex` | Complex coding tasks | Use a stronger coding model with good repository-scale reasoning. Balance quality, latency, and cost. |
| `coding_critical` | Critical or high-risk coding tasks | Use the most capable coding model available. Optimize for correctness and deep reasoning; treat latency and cost as secondary. |
| `analysis_regular` | Regular review, research, planning, and other non-coding analysis | Use a fast, low-cost reasoning model with reliable structured-output and instruction-following behavior. |
| `analysis_complex` | Complex non-coding analysis | Use a stronger reasoning model that can synthesize large inputs and preserve requirements and technical details. |
| `analysis_critical` | Critical analysis, criticism, advisory calls, and council participants | Use the most capable reasoning model available. Optimize for correctness, independent judgment, and detection of subtle risks. |
| `analysis_user` | Main `Auto` agent | Use a top-tier general reasoning model with reliable tool use and delegation. This model selects workflows, coordinates subagents, and produces the user-facing result. |
| `regular_user` | Main `General` and `ReadOnly` agents | Use a strong general-purpose model with reliable tool use. Balance quality, latency, and cost for interactive work. |
| `simple_user` | Main `Simple` agent | Use a fast, low-cost general-purpose model for lightweight tasks. It still needs reliable instruction following and tool use. |
| `summary` | Summaries of individual tool results for context projection | Use the fastest and cheapest model that produces accurate summaries. Prefer low reasoning effort. Its context window must fit the largest tool result that should be summarized, plus the summary prompts and output reserve. A smaller window causes that result to fall back to omission or standard serialization. |
| `extract` | Raw information extraction, adaptive compaction, and knowledge extraction and merging | Use a fast model with high summarization fidelity, low hallucination risk, and reliable format compliance. Prefer a context window at least as large as the largest context window among `*_user`, `coding_*`, and `analysis_*`. A smaller window still works for adaptive compaction, but can require extra reduction calls and can reduce fidelity. |
| `subagent_query` | Focused questions over a saved subagent conversation; also used by `/ask` in this configuration | Use a fast, low-cost model with good question answering and synthesis. Its context window must fit the complete saved subagent branch, the question, the query system prompt, and injected knowledge. The operation does not truncate oversized branches, so prefer a context window larger than or equal to every `coding_*` and `analysis_*` model window, with additional room for prompt and answer tokens. |

### Context-window relationships

- `summary` processes one tool result at a time, not the complete agent conversation. Size it for the largest projected tool result.
- `extract` compacts sessions created by all main and subagent models. Matching the largest agent context window minimizes adaptive reduction work; this is a performance and fidelity recommendation rather than a hard runtime requirement.
- `subagent_query` sends the complete saved subagent branch without truncation. Its context-window requirement is strict: an oversized request fails with `query_failed`.
- Context-window comparisons must use the models' usable input limits after reserving tokens for system prompts, tools, the requested answer, and provider-specific limits. Equal advertised context-window sizes do not guarantee equal usable input capacity.
