<change_request_template>
<!--
<change_request_guidelines>
  1. Universal: use for requested changes in code, documentation, configuration, data contracts, content, or operational artifacts.
  2. Focus on WHAT must change and WHY. Do not include implementation steps unless they are hard constraints.
  3. The request must be self-contained and understandable without external context. Links may support details, but must not replace explanation.
  4. Follow <id_guidelines> for all items. Every item must have a unique ID.
  5. Separate requested changes from affected areas. This keeps scope clear and avoids hidden work.
  6. Do not include rollout, rollback, or backward compatibility requirements unless they are explicitly requested.
  7. If information is missing, write "TBD" and record it in Open Questions.
  8. Any ambiguities must be either resolved with decisions or moved to Open Questions. No "ifs" outside of Open Questions.
</change_request_guidelines>
-->

# Change Request: {Name}

{Brief description of the requested change and why it is needed now.}

## Key definitions and abbreviations

- {id}: {Definition}. {Explanation of the term or abbreviation.}
- ...

## Scope
In scope:
- {id}: {what must change}
- ...

Out of scope:
- {id}: {what will not change}
- ...

## Requested Changes
<!--
This section MUST contain ENOUGH information to execute this document.
Otherwise, the executor will not have enough data to understand what exactly needs to be done.
Set of SPECIFIC changes that must be made including:
  - Detailed steps
  - Expected outcomes
  - Code snippets, configuration examples, etc.
-->
- {id}: {requested change; expected outcome}
- ...

## Affected Areas
- {id}: {artifact, system, document, or workflow} - {why it is affected}
- ...

## Overengineering and Overspecification Considerations
{Justification for why this technical solution does not introduce unnecessary complexity, overengineering or specify details that are not essential for understanding and implementation. KISS and YAGNI principles must be followed.}

## Constraints and Risks
- {id}: {constraint or risk; impact}
- ...

## Assumptions
- {id}: {justification (why); verification (how/when)}
- ...

## Open Questions
<!--
CRITICAL: These are questions that relate DIRECTLY TO THIS DOCUMENT. The presence of entries here means that the document is not ready for execution until these questions are resolved.
If the CHANGE REQUEST contains a request to ADD open questions to other documents, then that is `Requested Changes`, not open questions in this section.
-->

### {id}: {brief description}
  - {impact}
  - {what should answer look like}
  - {what has been done to find answer}
  - {how/when it will be resolved}

### ...

## References
- {id}: {reference (file/URL/standard, etc.)} - {one-line description}
- ...
</change_request_template>