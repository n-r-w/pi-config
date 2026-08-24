---
name: workflow-authoring
description: How to create and edit valid Pi workflow YAML catalogs. Use before creating, modifying, or reviewing workflow files under agent-suite/workflow/workflows/
disable-model-invocation: true
---

<context>
  <purpose>
    Define rules for authoring workflow YAML catalogs in `agent-suite/workflow/workflows/`.
  </purpose>

  <scope>
    1. Graph structure and validity.
    2. Prompt placement: root prompt vs stage prompts.
    3. Transition and return-policy design.
    4. Goal and task formulation.
    5. Style and verification.
  </scope>
</context>

<graph_constraints>
  1. Exactly one stage MUST be `initial`.
  2. At least one stage MUST be `final`.
  3. Stage `id` MUST be non-empty, single-line, free of whitespace.
  4. Stage `description` MUST be single-line trimmed text.
  5. Every stage MUST have a non-empty `prompt`.
  6. Stage IDs MUST be unique.
  7. Transition endpoints MUST reference existing stage IDs.
  8. Every non-final stage MUST have at least one outgoing `advance`.
  9. Final stages MUST NOT have outgoing `advance`.
  10. `advance` graph MUST be acyclic.
  11. Every stage MUST be reachable from the `initial` stage via `advance`.
  12. Every `rework` MUST target a strict `advance` ancestor of its source.
</graph_constraints>

<prompt_placement>
  1. Root `prompt` is active on every stage. Put only workflow-wide rules there: universal constraints, cross-cutting domain rules, safety guards, workflow gates.
  2. Stage `prompt` is active only when that stage is active. Put stage-specific rules there: goal, actions, completion criteria, stage constraints.
  3. A single-stage rule MUST live in that stage prompt, not in the root prompt.
  4. Cross-stage decision policy MUST live in root prompt. This includes policy that decides whether any stage stays active or uses `rework`.
  5. Stage-specific `Rework rules:` MUST define exact target conditions. They MUST NOT restate complete shared return policy.
  6. When workflow creates replacement workflow, copy every still-required cross-stage policy into generated root prompt. Source root prompt stops applying after replacement.
  7. Do not state a rule graph already enforces. Example: "do not skip stages" is meaningless because transitions prevent skipping.
  8. Do not repeat same rule in root prompt and stage prompt.
</prompt_placement>

<runtime_return_policy>
  1. Every workflow with `rework` transitions MUST include shared return policy in root `prompt`.
  2. Shared return policy controls runtime transition choice. It MUST NOT reduce or replace graph recovery paths.
  3. Workflow author MUST add every realistic `rework` edge even when shared policy makes edge use exceptional.
  4. Root `prompt` return policy MUST tell runtime model:
    1) Stay in current stage by default.
    2) Before return, identify exact output of target ancestor that became invalid or incomplete.
    3) Do not return when no ancestor output became invalid and current stage permits correction.
    4) Handle local errors, failed checks, implementation-detail changes, factual investigation, and clarifications in current stage when stage constraints permit.
    5) Treat working-detail change as local while goal, critical requirements, behavior, architecture, scope, dependencies, and material trade-offs remain valid.
    6) Return only when evidence materially invalidates ancestor output or current stage constraints prohibit required correction.
    7) Return to nearest ancestor that owns invalid output or prohibited correction.
    8) Do not return only because plan and implementation differ. Return to planning only when mismatch materially invalidates approved plan.
    9) Investigate evidence gap in current stage first. Return to collection only when required evidence cannot be gathered or refreshed there.
    10) Handle explanation, clarification, and allowed check rerun in current stage. Return on user feedback only when feedback invalidates ancestor output or requests prohibited change.
  5. Generated replacement workflow MUST receive this policy in its own root `prompt`.
</runtime_return_policy>

<transition_design>
  1. `advance` moves to a later stage; `rework` returns to an earlier stage.
  2. Design the `advance` chain first: each non-final stage advances to the next logical stage.
  3. Add `rework` only for realistic process return. Do not add `rework` only for stage coverage.
  4. Every stage with outgoing `rework` edges MUST contain one `Rework rules:` block. Stages without outgoing `rework` edges MUST NOT contain this block.
  5. Every outgoing `rework` edge MUST have at least one matching rule in source stage prompt.
  6. Use format `Return to TARGET_ID when EXPLICIT_CONDITION.` Wrap `TARGET_ID` in backticks in workflow prompt.
  7. Condition MUST identify concrete process event, evidence gap, invalidated decision, feedback, or failure that can require return. Generic conditions such as "when rework is needed", "when issues occur", or "go back and rework" are forbidden.
  8. Every target-specific return rule MUST have matching `rework` edge. Keep target-specific return instructions in `Rework rules:` block, not in `Actions`, constraints, or other stage sections. Keep shared runtime return policy in root prompt.
  9. Add full-restart `rework` edge from final stage to initial stage when realistic iterative cycles are expected.
  10. Add semantic skip-jump `rework` edges for every realistic return that must bypass immediate predecessor.
  11. Runtime return policy MUST NOT justify omission of realistic immediate or skip-jump recovery edge.
  12. Verify skip-jump `rework` targets are strict `advance` ancestors. Immediate predecessors always qualify; farther targets need manual check.
</transition_design>

<style_rules>
  1. Use text formatting, not markdown
  2. Use caveman style: no articles, short sentences, imperative mood.
  3. Structure stage prompts as `Goal`, `Actions`, optional `Subagents` and `Rework rules`, then optional `Constraints` or completion criteria.
  4. Use `Subagents: required | forbidden | if needed` or more complex custom rules, if subagents are allowed. Omit if not allowed.
  5. Reference stages by `id` when a prompt describes a return target.
  6. Omit filler, motivational text, and decorative phrasing.
</style_rules>

<subagents_rules>It is RECOMMENDED to implement using subagents and verify their work using main agent, and not other way around</subagents_rules>

<goal_writing>
  1. `Goal:` states the desired end state of the stage.
  2. Goal answers "what must be achieved". Task answers "what to do". They are different concepts.
  3. Goal MUST describe a verifiable outcome, not an action.
  4. Tasks MUST go into `Actions`, never into `Goal:`.
  5. Formulate goal as resulting state. Optionally add outcome purpose.
  6. Wrong and right pairs:
    - `Goal: Collect facts.` (action) vs `Goal: Build evidence base for analysis.` (outcome)
    - `Goal: Get user approval of plan.` vs `Goal: Obtain agreed plan.`
    - `Goal: Present results.` vs `Goal: Provide user with necessary information.`
</goal_writing>

<verification>
  1. Parse file as YAML.
  2. Check graph constraints programmatically.
  3. Check prompt placement manually: read root prompt and every stage prompt.
  4. When graph has `rework`, check root prompt contains complete policy defined in runtime return policy section.
  5. Check generated workflow root receives runtime return policy after source workflow replacement.
  6. Check runtime policy did not remove realistic recovery edges. Review failures, user feedback, stale evidence, invalidated decisions, and stage-prohibited corrections.
  7. Check every outgoing `rework` edge has matching explicit condition in source stage `Rework rules:` block.
  8. Check every return rule names exact target stage `id` and has matching `rework` edge.
  9. Check stage-specific return conditions work with root runtime gate without duplicating it.
  10. Check every realistic skip-jump recovery path exists and targets strict ancestor.
  11. Reject generic, missing, or ambiguous return conditions.
</verification>

<example>
```yaml
description: Example workflow
prompt: |-
  Shared workflow goal: Solve the user's problem.

  Shared rules:
  1. Ground every conclusion in evidence.

  Return policy:
  1. Stay in current stage for local errors, clarifications, and evidence available there.
  2. Before rework, identify exact target-stage output that became materially invalid.
  3. Return to nearest stage that owns invalid output or correction prohibited in current stage.

stages:
  - id: collect_facts
    description: Collect facts
    prompt: |-
      Goal: Gather foundation for creating plan

      Subagents:
      1. MUST read primary documentation (specs, plans, etc.) yourself, not trusting subagents.
      2. Use subagents to collect additional information on codebase or logically related non-primary resources.

      Actions:
      1. Identify information sources.
      2. Gather evidence.
    initial: true

  - id: approve_plan
    description: Approve plan
    prompt: |-
      Goal: Obtain agreed plan for solving problem

      Subagents: forbidden

      Actions:
      1. Present plan.
      2. Wait for explicit approval.

      Rework rules:
      1. Return to `collect_facts` when evidence essential for plan approval is missing or stale and cannot be gathered in current stage.

  - id: report
    description: Report results
    prompt: |-
      Goal: Provide user with necessary information

      Subagents: forbidden

      Actions: Present findings and evidence.

      Rework rules:
      1. Return to `approve_plan` when feedback materially invalidates approved plan but evidence base remains valid.
      2. Return to `collect_facts` when feedback materially invalidates evidence base and required evidence cannot be refreshed in current stage.
    final: true

transitions:
  - from: collect_facts
    to: approve_plan
    type: advance
  - from: approve_plan
    to: report
    type: advance
  - from: approve_plan
    to: collect_facts
    type: rework
  - from: report
    to: approve_plan
    type: rework
  - from: report
    to: collect_facts
    type: rework
```
</example>
