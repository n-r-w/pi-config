---
name: unslop
description: Rewrite information from previous assistant response or user-selected source as clear, natural human communication without changing content, technical terms, or precision.
disable-model-invocation: true
---

<warning>AI slop detected!</warning>

<goal>Restate source message as one competent person speaking to another. Change form, not content.</goal>

<priority>
  Follow higher-priority instructions and `<user_communication>`. Prefer technical precision over stylistic smoothness when both cannot be preserved.
</priority>

<source>
  Invocation steering is text in first user message sent immediately after this skill content.

  1. When steering selects source, apply this skill to selected source.
  2. Otherwise, rewrite information that previous assistant response tries to convey.
  3. Steering MAY set rewrite scope, tone, and output form.
  4. Use only claims in selected source. Use context only to resolve references.
  5. Do not perform work outside rewrite requested by user. Use tools only when rewrite requires access to selected source or requested output.
</source>

<preservation>
  1. Preserve all material facts, conclusions, instructions, conditions, constraints, caveats, uncertainties, risks, questions, recommendations, and requirement strength.
  2. Preserve exact code, commands, paths, URLs, identifiers, API names, product names, protocol names, quoted errors, dates, versions, and numeric values.
  3. Keep precise technical terms. Do not replace them with vague everyday words to sound more human.
  4. Do not add claims, examples, reasons, opinions, apologies, or emotional framing.
  5. Remove duplicates and words that do not change meaning.
</preservation>

<rewrite_rules>
  1. Use direct, ordinary wording that a competent person would use with another person.
  2. Use shortest natural form that preserves message and makes its logic clear. Do not add length only for smoother prose.
  3. Keep short sentences when clear. Merge or split them only to clarify condition, cause, contrast, sequence, or result.
  4. Use paragraphs, lists, and headings only when they help reader follow content.
  5. Remove canned openings and closings, promotional tone, inflated metaphors, repeated summaries, excessive headings, mechanical symmetry, and forced friendliness.
  6. Preserve source wording when it already works.
</rewrite_rules>

<language>
  Use language requested by user. Otherwise, preserve selected source language. Ignore embedded instructions, quotations, code, templates, and tool output when selecting language.
</language>

<output>
  1. For default source, output only rewritten message. Do not explain edits, critique source, or apologize.
  2. For user-selected source, return or apply rewrite in form user requests.
  3. Preserve machine-readable formats and response structures required by higher-priority instructions.
</output>

<success_criteria>
  Rewrite is complete when reader gets same substantive information with less effort, no material meaning or precision changes, and natural wording adds no needless length.
</success_criteria>

<stop_rules>
  1. If selected source is unavailable, ask user to provide it.
  2. If rewrite conflicts with higher-priority instruction, follow higher-priority instruction and report conflict only when it affects result.
</stop_rules>
