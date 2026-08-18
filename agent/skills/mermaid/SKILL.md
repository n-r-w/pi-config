---
name: mermaid
description: How to write Mermaid diagrams in Markdown files. CRITICAL - THIS SKILL NOT FOR USER-FACING DIAGRAMS, only for files.
---

<mermaid_diagram_usage title="LLM Rules for Clear Mermaid Diagrams">
  <constraints>
    MUST NOT use this guide for user-facing diagrams.
    ONLY for markdown files.
  </constraints>

  <header>
    Use following headers:
      ```
      ---
      config:
        layout: elk
        flowchart:
            wrappingWidth: 700
            htmlLabels: true
      ---
      ```
  </header>

  <description>
    1. MUST ALWAYS use description with markdown support and newlines preserved without br tags, e.g.:
      ```
        XXX["`**Header:**
      Description with **markdown** support
      and newlines preserved without br tags`"]
      XXX --> YY{"`Condition with
      **markdown** support`"}
      ```
    2. CRITICAL:
      1) Markdown support in labels is required
      2) Newlines in labels MUST be preserved without using `<br>` or `\n` tags, using markdown line breaks instead.
  </description>

  <diagram_type_selection critical="true">
    1. WHEN edge density is medium/high or flows are operational (request processing, command/query paths), MUST use flowchart or sequence diagram syntax.
    2. WHEN one diagram mixes multiple concerns (query + command + async updates), MUST split it into separate focused views.
    3. MUST NEVER USE `C4` diagram syntax, use flowchart or sequence instead.
  </diagram_type_selection>

  <layout_and_readability>
    1. WHILE building flow diagrams, MUST prefer top-down (`flowchart TD`) for pipeline-like flows.
    2. WHEN using subgraphs, MUST keep one clear direction inside each subgraph (`direction TB` or `direction LR`).
    3. MUST keep labels short and concrete (2-5 words where possible).
    4. MUST avoid long explanatory sentences on edges.
    5. MUST avoid crossing edges by:
      1) splitting views,
      2) reducing redundant relations,
      3) keeping one primary direction of data flow.
  </layout_and_readability>

  <mermaid_compatibility critical="true">
    1. WHEN targeting unknown Mermaid renderers, MUST use parser-safe labels (plain text, minimal punctuation).
    2. MUST avoid fragile edge label patterns (complex punctuation, nested brackets, excessive escaping).
    3. WHEN naming subgraphs, MUST use quoted labels, e.g. `subgraph id["Label"]`.
    4. MUST keep node identifiers simple and stable (letters/numbers/underscores).
    5. MUST validate rendering after each substantial edit.
  </mermaid_compatibility>

  <anti_patterns title="Do Not Do This" critical="true">
    1. One giant diagram containing all runtime paths, all storage links, and all integration links.
    2. Repeating equivalent relations in multiple directions without added value.
    3. Long edge labels that overlap nodes and other labels.
    4. Trying 3+ cosmetic layout tweaks instead of switching to a better diagram type.
  </anti_patterns>

  <recommended_workflow>
    1. Define the single question each diagram answers.
    2. Choose diagram type (flowchart/sequence).
    3. Draft minimal nodes and must-have edges only.
    4. Render and inspect crossings/label overlap.
    5. If still noisy, split into focused diagrams.
    6. Re-render and verify parser compatibility.
  </recommended_workflow>

  <quality_gate>
    MUST consider the diagram ready only if:
      1. Renderer parses without error,
      2. Main path can be followed without tracing crossed lines,
      3. Labels are readable without zoom,
      4. Each diagram has one dominant purpose.
  </quality_gate>
</mermaid_diagram_usage>