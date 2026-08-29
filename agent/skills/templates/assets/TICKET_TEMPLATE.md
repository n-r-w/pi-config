<ticket_template>
<!--
<ticket_guidelines>
  1. Used to describe requirements for implementing a specific task. Unlike a task template, a ticket does not describe in detail how to implement task, but only what needs to be done.
  2. ticket must be self-contained and understandable without opening related documents.
  3. Follow <id_guidelines> for all items, and ensure every item has a unique ID.
  4. Make verification explicit. Each acceptance criterion must be checkable from ticket alone.
  5. Do not include rollout, rollback, or backward compatibility requirements unless they are explicitly required.
  6. If information is missing, write "TBD" and record it in Open Questions.
  7. Prefer work items that end in a verifiable artifact or observable behavior.
  8. Any ambiguities must be either resolved with decisions or moved to Open Questions. No "ifs" outside of Open Questions.
  9. One ticket MUST contains EXACTLY ONE vertical slice:
    1) Each slice cuts a narrow but COMPLETE path through every layer (schema, API, UI, tests): vertical, NOT a horizontal slice of one layer
    2) A completed slice is demoable or verifiable on its own
    3) Each slice is sized to fit in a single fresh context window
    4) Any prefactoring should be done first
</ticket_guidelines>
-->

# Ticket: {ID/Name of ticket}
{Brief description of ticket}

## Key definitions and abbreviations
- {id}: {Definition}. {Explanation of term or abbreviation.}
- ...

## Problem Statement
{If there is an external document, then instead of a description here, you need to provide a link to the document}
{Description of problem that this ticket is intended to solve.
Why is this ticket needed?
What is impact of not implementing it?}

## Target Picture
{Description of desired state after ticket is implemented.
What will be different after ticket is implemented?}

## Scenarios

### {id}: {Definition}
  - Actor: {who is involved in this scenario}
  - Pre-condition: {conditions that must be met before this scenario can occur}
  - Trigger: {event that initiates the scenario}
  - Required behavior: {expected behavior of the system}
  - Example input and expected output: {example input and expected output}

### ...

## Scope
In scope:
- {id}: {what this ticket will implement}
- ...

Out of scope:
- {id}: {what this ticket will not implement}
- ...

## Dependencies and Preconditions
- {id}: {dependency, prerequisite, or blocker; current status}
- ...

## Requirements <!-- remember that we specify requirements, not how to implement them -->

### Goals
- {id}: {goal to achieve; rationale}
- ...

### Functional Requirements
- {id}: {requirement; rationale}
- ...

### Non-Functional Requirements
- {id}: {requirement; rationale}
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

### {id}: {brief description}
  - {impact}
  - {what should answer look like}
  - {what has been done to find answer}
  - {how/when it will be resolved}

### ...

## Technical Supplement
<!--
If technical details need, they MUST be ONLY here.
The other sections of ticket should contain only business-oriented details.

IMPORTANT:
 1. In most cases, ticket should EXCLUDE technical details, so adding them to technical supplement should be an exception, not rule.
 2. It is PROHIBITED to reference items from technical supplement from other sections of ticket. If connections are required, then only technical supplement can refer to other sections of ticket, but not vice versa.
 3. Choose structure and format of technical supplement that best suits specific case.
-->

## References
- {id}: {reference (file/URL/standard, etc.)} - {one-line description}
- ...
</ticket_template>