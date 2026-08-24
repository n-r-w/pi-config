<plan_report_template>
<!--
<plan_guidelines>
  1. Do not include code samples, low-level build steps, timelines, invented dates, or invented staffing. Use "TBD" for unknowns and record assumptions and dependencies explicitly.
  2. Do not assume file contents. Read relevant code, documentation, and other evidence before making decisions. Plans must be based on verified evidence from codebase, docs, web research, and other relevant sources.
  3. document must be understandable without external context. Links may support details, but must not replace explanation.
  4. All document items must follow <id_guidelines>. Every item must have a unique ID.
  5. Do not plan backward compatibility, fallbacks, release plan, or rollback plan unless explicitly required by user.
  6. Phase management:
    1) Follow TDD principles in planning: RED -> GREEN -> REFACTOR. Phases should preserve testability whenever practical.
    2) Keep phases meaningful, testable, and easy to review. Avoid phases that are so small they add process overhead or so large they hide risk.
    3) Prefer phase boundaries that preserve a working, verifiable state after each phase.
    4) If temporary non-verifiable intermediate states are unavoidable, make that explicit and define steps required to regain verifiability in next phase.
    5) Combine setup and verification work when separating them would leave system in a hard-to-test or misleading intermediate state.
    6) Include a final independent verification phase.
    7) Include a final cleanup phase that removes temporary code, workarounds, debugging artifacts, and other implementation residue.
    8) Critical risk: trying to make every phase testable can easily lead to inventing unnecessary logic and tests that will not provide real value when looking at task as a whole. You MUST analyze EACH item in `Exit criteria` section from this point of view and eliminate excessive verification requirements, or combine phases to avoid such situations. MUST provide justification in `Decomposition Justification` section.
    9) MUST break work into vertical slices:
      - Each slice cuts a narrow but COMPLETE path through every layer (schema, API, UI, tests): vertical, NOT a horizontal slice of one layer
      - A completed slice is demoable or verifiable on its own
      - Each slice is sized to fit in a single fresh context window
      - Any prefactoring should be done first
  7. Every referenced function, class, file, and module name must link to its location in codebase, but without line numbers or other context that may change over time.
  8. Any ambiguities must be either resolved with decisions or moved to Open Questions. No "ifs" outside of Open Questions.
</plan_guidelines>
-->

# Delivery Plan: {Name}

{Brief description of plan goal and scope.}

## Key definitions and abbreviations

- {id}: {Definition}. {Explanation of term or abbreviation.}
- ...

## Delivery Strategy
{release model; feature flags/config; data migration strategy}

## Main Changes
{Brief description of main changes that will be made to system as part of plan.}

## Entities and Invariants
{List of main entities and invariants that will be used or affected by plan}

## New Folders and Components
{List of new folders and components that will be created as part of plan}

## Backward Compatibility
{Describe any backward compatibility considerations, if applicable. If none, state "No backward compatibility."}

## Phased Plan
{Each phase must end with verifiable outcomes and explicit exit criteria. Break work into vertical slices}

### Phase Tree
{Mermaid flowchart. Use dependencies as connectors between phases.
Take into account possibility of parallel work on phases.}

### Decomposition Justification
{Justification for why this particular phase division was chosen.
 Proof that breaking down into phases does not generate unnecessary logic and tests, in order to achieve principle of "each phase is testable"}

## Overengineering and Overspecification Considerations
{Justification for why this technical solution does not introduce unnecessary complexity, overengineering or specify details that are not essential for understanding and implementation. KISS and YAGNI principles must be followed.}

### Phase {id} - {Brief Name}

#### Goal
{Goal of phase}

#### Work
{Which tasks will be done in this phase; how they will be verified}

#### Deliverables
{What will be delivered at end of phase}

#### Exit criteria
{What conditions must be met to consider phase complete and move to next one}

#### Risks
{What are risks associated with this phase; how will they be mitigated}

### Phase ...
...

## Test Strategy
{Unit, integration, E2E, load/stability tests; what is covered by each; how to verify}

## Release and Rollback <!-- only if explicitly requested by user -->
{Release steps; monitoring strategy; rollback conditions and actions; post-release validation}

## Dependencies and Resourcing
{External dependencies and other resources required}

## Project Definition of Done
{Documentation, observability, operational readiness, SLO/stability targets, etc.}

## Assumptions
- {id}: {Justification (why); Verification (how/when)}
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
- {id}: {Reference (file/URL/standard, etc.)} - {one-line description}
- ...
</plan_report_template>