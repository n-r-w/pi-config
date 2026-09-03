<!--
<specification_guidelines>
  1. Prefer bullets over long paragraphs; keep each bullet to 1–2 sentences.
  2. Do not invent facts. If ambiguity cannot be resolved, write "TBD" and add an item to Open Questions.
  3. Focus on WHAT and WHY, not HOW. Do not include implementation plans, component-level architecture, or code.
  4. The document must be self-contained and understandable without external context. Links may provide extra detail, but must not replace essential explanation.
  5. Be specific. Avoid vague placeholders such as "if", "may", "optional", "not confirmed", or "mandatory model", unless the item is an explicit open question with its own ID.
  6. Use IDs only when required by <id_guidelines>. Do not infer this need from the item type or template section.
  7. Every assigned ID must be unique within the document.
  8. Functional requirements, non-functional requirements, and decisions must use stable IDs because downstream plans, tasks, tests, and reviews reference them.
  9. Any ambiguities must be either resolved with decisions or moved to Open Questions. No "ifs" outside of Open Questions.
</specification_guidelines>
-->

<specification_report_template>
# Specification: {Initiative / Feature Name}

## Key definitions and abbreviations

- {Definition}. {Explanation of the term or abbreviation.}
- ...

## Context and Problem
Current state: {1–3 sentences}

Problems:
- {What is wrong and why it matters}
- ...

Why now:
- {Trigger/opportunity/risk/debt}
- ...

## Goal (Outcome)
Goal: {one sentence; measurable}

Success metrics:
- {Target value or direction}
- ...

Non-goals:
- {Explicitly out of scope item}
- ...

## Scenarios
Actors:
- {Who: users, systems, etc.}
- ...

Top scenarios (3–5):
- {Scenario + expected outcome}
- ...

Operational/UX constraints: {if any}
- {Constraint description}
- ...

## Scope of Change
In scope:
- {What will change / be added}
- ...

 Out of scope:
- {What will NOT be done}
- ...

System/domain boundaries:
- {Which systems are affected; which are not}
- ...

## High-Level Requirements
Functional:
- {FRQ-id}: {requirement description; rationale}
- ...

Non-functional:
- {NFQ-id}: {latency/availability/security/compliance/etc}
- ...

## Overengineering and Overspecification Considerations
{Justification for why this technical solution does not introduce unnecessary complexity, overengineering or specify details that are not essential for understanding and implementation. KISS and YAGNI principles must be followed.}

## Risks
Risks:
- {What could go wrong; how to reduce likelihood/impact}
- ...

## Assumptions
- {Justification (why); verification (how/when)}
- ...

## Open Questions

### {QST-id}: {brief description}
  - {impact}
  - {what should answer look like}
  - {what has been done to find answer}
  - {how/when it will be resolved}

### ...

## Decisions
- {DEC-id}: {decision description; rationale; how/when it should be resolved}
- ...

## Standards Deviations
{Describe any deviations from established coding/design/testing/documentation skills/standards/rules, with reasoning and tradeoffs considered.}

## Technical Supplement
<!--
If technical details need to be recorded during the specification creation process, they MUST be ONLY here.
The other sections of the specification should contain only business-oriented details.

IMPORTANT:
 1. In most cases, the specification should EXCLUDE technical details, so adding them to the specification should be an exception, not the rule.
 2. It is PROHIBITED to reference items from the technical supplement from other sections of the specification. If connections are required, then only the technical supplement can refer to other sections of the specification, but not vice versa.
 3. Choose the structure and format of the technical supplement that best suits the specific case.
-->

## References
- {Reference (file/URL/standard, etc.)} - {one-line description}
- ...
</specification_report_template>