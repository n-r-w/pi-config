---
name: pi-tui-rendering
description: Pi agent TUI rendering instructions for custom tools
---

<pi_tui_rendering>

  <core_rules>
    1. Custom tools MUST use Pi's default tool shell unless the tool needs full control over framing, padding, or background.
    2. Do not set `renderShell: "self"` only to work around text wrapping, Unicode width, or BiDi display bugs.
    3. Use Pi TUI components for text rendering when the tool output is plain text or Markdown:
      1) `Text` for wrapped text.
      2) `Markdown` for Markdown output.
      3) `Box(1, 1)` behavior as the default tool-shell width contract.
    4. Do not implement custom Unicode, RTL, or BiDi sanitization in tool renderers.
    5. Do not remove, rewrite, or normalize user-visible Unicode content to avoid terminal rendering bugs.
    6. If a bug reproduces in a Pi built-in tool such as `bash`, treat it as a Pi/TUI issue. Keep local code aligned with Pi unless the custom tool has different product semantics.
  </core_rules>

  <render_call_rules>
    1. `renderCall` SHOULD stay compact.
    2. Long tool arguments MUST NOT consume collapsed result space.
    3. If a tool-call header contains arbitrary user input, render it as one width-bounded row.
    4. If Pi `Text` wrapping would expand a tool-call header into multiple rows, use a narrow single-line clipping helper.
    5. Single-line clipping helpers MUST be reset-free when the line is rendered inside a colored Pi tool box.
  </render_call_rules>

  <render_result_rules>
    1. Match rendering strategy to tool semantics.
    2. For final-output tools, collapsed output SHOULD use Pi visual text rendering. Examples: LLM answers, command output, model summaries.
    3. For event-stream tools, collapsed output MAY keep custom row selection when events are semantic units.
    4. Expanded output SHOULD use Pi `Text` or `Markdown` unless the expanded view has a structured layout that Pi components cannot express.
    5. Hidden-content hints SHOULD use Pi keybinding helpers and Pi width utilities where they do not break parent styling.
  </render_result_rules>

  <final_output_tool_pattern>
    Use this pattern for tools that return one answer after execution.

    1. Keep the default Pi tool shell.
    2. Keep `renderCall` compact and width-bounded.
    3. Build collapsed answer text with the tool label and answer content.
    4. Delegate collapsed answer wrapping to Pi `Text.render(width)`.
    5. Apply the collapsed line budget after Pi returns visual lines.
    6. Preserve the product semantics of the preview:
      1) Use first visual lines when the answer starts with the most useful summary.
      2) Use last visual lines only when the tool behaves like terminal output.
    7. Use `Markdown` for expanded answer content when Markdown is valid output.
  </final_output_tool_pattern>

  <event_stream_tool_pattern>
    Use this pattern for tools that stream logical progress events.

    1. Keep the default Pi tool shell.
    2. Treat each progress event as a semantic unit.
    3. Collapsed view SHOULD prioritize the latest events when recent progress matters most.
    4. One collapsed event SHOULD render as one row unless the tool explicitly supports multi-row events.
    5. Do not replace event-row rendering with `Text.render()` when that would let one event consume multiple collapsed rows.
    6. Keep tool-owned event-row clipping narrow:
      1) clip one already-selected row;
      2) preserve grapheme boundaries;
      3) avoid `\x1b[0m` inside colored box content;
      4) do not add custom BiDi logic.
    7. Expanded view SHOULD use Pi `Text` or `Markdown` for full event text, stderr, and final output.
  </event_stream_tool_pattern>

  <width_and_ansi_rules>
    1. Every rendered line MUST satisfy `visibleWidth(line) <= width` for the width passed to `render(width)`.
    2. Remember that Pi's default tool shell uses `Box(1, 1)`: child renderers receive `width - 2`, then the box adds horizontal padding and background.
    3. Do not test only `renderResult().render(width)` when production rendering uses the default Pi tool shell.
    4. Test the shell contract with public Pi TUI components, such as `Box(1, 1)`, when the shell affects behavior.
    5. Avoid `truncateToWidth()` in pre-styled single-line rows when the active Pi version emits SGR reset sequences that break parent styling.
    6. If a tool-owned reset-free clipping helper is needed, keep it small, plain-text-only, and scoped to the renderer path that needs it.
  </width_and_ansi_rules>

  <stop_rules>
    1. Stop and ask the user before adding custom BiDi handling, Unicode sanitization, tool-rendering compatibility wrappers, fallback renderers, or `renderShell: "self"`.
    2. Stop and inspect Pi source before changing behavior that Pi built-in tools already implement.
    3. Stop and report the evidence when the same rendering bug reproduces in a Pi built-in tool.
  </stop_rules>

</pi_tui_rendering>