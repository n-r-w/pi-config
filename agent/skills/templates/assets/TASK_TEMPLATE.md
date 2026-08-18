<task_template>
<!--
<task_guidelines>
  1. Use for a detailed development task that implements one approved phase, stage, or work package from a plan.
  2. Keep task concrete and executable. Focus on exact work to be done in this task, not whole project.
  3. Task must be self-contained and understandable without opening original plan, though it may link back to it.
  4. Follow <id_guidelines> for all items. Every item must have a unique ID.
  5. Include only work that can be completed and verified within this task. Move unrelated or future work to Out of scope or Dependencies.
  6. Make verification explicit. Each acceptance criterion must be checkable from task alone.
  7. Do not include rollout, rollback, or backward compatibility requirements unless they are explicitly required.
  8. If information is missing, write "TBD" and record it in Open Questions.
  9. Prefer work items that end in a verifiable artifact or observable behavior.
  10. Any ambiguities must be either resolved with decisions or moved to Open Questions. No "ifs" outside of Open Questions.
</task_guidelines>
-->

# Development Task: {Name}

{Brief description of task, problem it solves, and phase or stage of plan it comes from.}

## Key definitions and abbreviations

- {id}: {Definition}. {Explanation of term or abbreviation.}
- ...

## Source Context
- {id}: {source plan / phase / stage / task reference}
- {id}: {why this task exists now}
- ...

## Scope
In scope:
- {id}: {what this task will implement}
- ...

Out of scope:
- {id}: {what this task will not implement}
- ...

## Dependencies and Preconditions
- {id}: {dependency, prerequisite, or blocker; current status}
- ...

## Overengineering and Overspecification Considerations
{Justification for why this technical solution does not introduce unnecessary complexity, overengineering or specify details that are not essential for understanding and implementation. KISS and YAGNI principles must be followed.}

## File Structure
{what directories and files will be created or modified as part of this task, with a brief description of their purpose}
- {folder/file path}: {description}

## Code Structure
{what classes, functions, modules, or other code structures will be created or modified as part of this task, with a brief description of their purpose}
- {code structure name}: {description}
- ...

## Work Items
<!--
This section MUST contain ENOUGH information (detailed steps, expected outcomes, code snippets, configuration examples, etc.) to execute this document with minimum external context.
Otherwise, executor will not have enough data to understand what exactly needs to be done.
MUST follows TDD principles: has explicit RED-GREEN-REFACTOR steps or decision why not applicable.
If sequence of steps has logical blocks, then explicitly highlight separate groups of steps with their own completion criteria.
-->
- {id}: {specific development task; expected result}
- ...

## Deliverables
- {id}: {artifact, code path, document, config, test, or other output}
- ...

## Acceptance Criteria
- {id}: {verifiable outcome}
- ...

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

## References
- {id}: {reference (file/URL/standard, etc.)} - {one-line description}
- ...
</task_template>