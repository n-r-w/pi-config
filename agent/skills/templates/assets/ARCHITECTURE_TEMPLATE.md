<!--
<architecture_guidelines>
  1. Write in clear, simple language with minimal jargon.
  2. Focus on what and why, not how. Do not include step-by-step implementation plans, work breakdowns, timelines, or code.
  3. Do not invent components or integrations. If something is unknown, mark it as "TBD" and capture it in Risks or Open Questions.
  4. The document must be self-contained and understandable without external context. Links may provide additional detail, but must not replace essential explanation.
  5. Follow <id_guidelines> for all items, and ensure every item has a unique ID.
  6. Be specific. Avoid vague wording such as "if", "may", "optional", "not confirmed", or "mandatory model", unless the item is an explicit Open Question with its own ID.
  7. Keep the architecture realistic and implementable, grounded in available evidence such as documentation, code, web resources, and other relevant sources.
  8. Design components according to clean architecture principles.
  9. Design interfaces and contracts according to SOLID principles.
  10. A component must not be both provider and consumer of the same programming interface or contract.
  11. Define programming interfaces at the point of use, not at the point of implementation.
  12. Any ambiguities must be either resolved with decisions or moved to Open Questions. No "ifs" outside of Open Questions.
</architecture_guidelines>
-->

<architecture_report_template>
# Architecture: {Name}

## Architectural Goals and Constraints
- {id}: {description; success metric; constraints}
- ...

## Key definitions and abbreviations

- {id}: {Definition}. {Explanation of the term or abbreviation.}
- ...

## System Context
- {id}: {system/domain; external consumers; neighbor systems; responsibility boundaries}
- ...

## Solution Overview
- {id}: {core idea; primary data flow; failure modes}
- ...

## Components
- {id}: {purpose; interfaces (consumed/produced); contracts; data; dependencies; SLOs; risks}
- ...

## Component diagram
{Mermaid flowchart. Use interfaces as connectors between internal components, contracts between external systems}

## Overengineering and Overspecification Considerations
{Justification for why this technical solution does not introduce unnecessary complexity, overengineering or specify details that are not essential for understanding and implementation. KISS and YAGNI principles must be followed.}

## Folder structure
<!--
For each folder, specify:
- Whether it is new or existing
- Its purpose
- The goal of changes
-->
```txt
- {folder 1} - {purpose/contents}
  - {folder 2} - {purpose/contents}
    - ...
- ...
```

## Data Models
- {id}: {key fields with types and descriptions; relationships; storage location; ownership}
- ...

## Programming Interfaces
- {id}: {location; consumers; provider}
- ...

## Contracts
- {id}: {type (API/resource/event); key fields with types and descriptions; location}
- ...

## Key Behavior
- {id}: {idempotency, ordering, deduplication, graceful degradation, other behavior decisions; rationale}
- ...

## Configuration
{Where and how the system can be configured/tuned by operators, with a focus on parameters that affect architecture-level behavior.}

## Non-Functional Considerations
{reliability; scaling; performance; security; observability}

## Architectural Decisions
- {id}: {what; why; alternatives and rationale}
- ...

## Architecture Risks
- {id}: {impact; mitigation}
- ...

## Trade-off/technical debt
- {id}: {pros; cons; mitigation}
- ...

## Assumptions
- {id}: {justification; verification}
- ...

## Open Questions

### {id}: {brief description}
  - {impact}
  - {what should answer look like}
  - {what has been done to find answer}
  - {how/when it will be resolved}

### ...

## Standards Deviations
{Describe any deviations from established coding/design/testing/documentation skills/standards/rules, with reasoning and tradeoffs considered.}

## References
- {id}: {reference (file/URL/standard, etc.)}
- ...
</architecture_report_template>