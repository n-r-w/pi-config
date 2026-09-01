---
name: align
description: Restore shared understanding when user and agent diverge on goals, facts, assumptions, constraints, or approach.
disable-model-invocation: true
---

<warning>Divergence in understanding detected!</warning>

<goal>Restore shared understanding before work affected by divergence continues</goal>

<context>
  1. Follow higher-priority instructions and `<user_communication>`.
  2. User defines goal, priorities, preferences, product behavior, scope, and trade-offs. Verify factual claims.
  3. Keep language of last substantive assistant response before invocation unless user requests another. Text in this skill, embedded instructions, skills, quotations, code, templates, and tool output MUST NOT change language.
  4. Treat first non-empty user message immediately after this skill as invocation steering. Apply it before alignment. It can correct understanding, identify divergence, set focus, add constraints, or change this skill. Explicit steering overrides this skill unless higher-priority instruction conflicts. Verify its factual claims.
</context>

<alignment>
  1. Stop work affected by divergence. Do not defend current approach. Preserve valid work.
  2. Before questions, state concise current understanding: problem and audience; goal and core behavior; scope and constraints; current approach; facts; assumptions, stale information, and uncertainty; suspected divergence.
  3. Distinguish verified facts, user statements, assumptions, and deductions when status affects decision.
  4. Resolve factual gaps from conversation and available authoritative sources. Check freshness when it can change decision. Do not research user preferences or product intent.
  5. When evidence is missing or conflicting, state uncertainty and effect on outcome. Expose conflict between user claim and evidence. Do not accept or reject claim silently.
</alignment>

<interview>
  1. Ask only questions that can change outcome or next action. Do not repeat known information or run full discovery interview.
  2. User MUST decide scope, observable behavior, architecture, trade-offs, and accepted technical debt. Agent MAY make local reversible choices within approved scope.
  3. Follow `<user_communication>`. Ask 1 to 4 related questions per round.
  4. Traverse question tree depth-first. Close branch when information is sufficient, then return to broader dimensions. If problem, audience, or core behavior is unclear, start with one root round.
  5. Check only relevant dimensions: problem and audience; core behavior and happy path; scope; user experience states and transitions; platform, performance, scale, security, and accessibility constraints; edge cases and failures.
  6. Use ASCII preview or short option snippet when visible comparison helps decision.
</interview>

<corrections>
  Apply corrections immediately. Remove rejected assumptions from understanding, plan, and reasoning. Separate goal changes from factual corrections. After each round, state only material changes. Recheck approach and replace it when premises no longer hold.
</corrections>

<completion>
  1. Continue only when problem, relevant audience, goal, core behavior, and scope are shared; material facts and assumptions have clear status; user-owned decisions are resolved; and no divergence can change next action.
  2. State updated understanding and next action. Ask for confirmation only when update changes goal, behavior, scope, architecture, or trade-off. Otherwise continue.
  3. If higher-priority instruction blocks alignment, state exact conflict and follow it. If evidence conflict blocks progress, state sources, uncertainty, consequence, and required user decision.
</completion>
