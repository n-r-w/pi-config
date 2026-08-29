<!--
<technical_solution_guidelines>
1. Document is intended to describe an arbitrary set of technical solutions.
2. It can be used as:
    1) An appendix to a specification, architecture description, implementation plan
    2) Preparation for creating a change request or ticket/task
    3) As a standalone document.
3. Prefer bullets over long paragraphs; keep each bullet to 1–2 sentences.
4. Document must be self-contained and understandable without external context. Links may provide extra detail, but must not replace essential explanation.
5. Be specific. Avoid vague placeholders such as "if", "may", "optional", "not confirmed", or "mandatory model", unless item is an explicit open question with its own ID.
6. Follow <id_guidelines> for all items, and ensure every item has a unique ID.
7. Any ambiguities must be either resolved with decisions or moved to Open Questions. No "ifs" outside of Open Questions.
</technical_solution_guidelines>
-->

<technical_solution_template>
# Technical Solution: {Name}

## Problem Statement
{If there is an external document, then instead of a description here, you need to provide a link to document}
{A concise statement of problem being solved, including any relevant context or constraints.}
- {id}: {Definition of problem or challenge}
- ...

## Proposed Solution
{A brief and structured description of proposed technical solution, including any relevant design decisions, trade-offs, and implementation details. Use markdown headings, bullet points, diagrams, and code snippets as needed to clearly communicate solution.}

## Overengineering and Overspecification Considerations
{Justification for why this technical solution does not introduce unnecessary complexity, overengineering or specify details that are not essential for understanding and implementation. KISS and YAGNI principles must be followed.}

## Open Questions

### {id}: {brief description}
  - {impact}
  - {what should answer look like}
  - {what has been done to find answer}
  - {how/when it will be resolved}

### ...

## References <!-- if any -->
- {id}: {reference (file/URL/standard, etc.)} - {one-line description}
- ...
</technical_solution_template>