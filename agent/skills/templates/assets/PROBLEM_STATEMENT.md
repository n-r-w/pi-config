<!--
<problem_statement_guidelines>
1. Describe problem, not solution.
    Phrases like “build a service / dashboard / API / button” are not a problem statement. First, you need to capture what exactly is not working, for whom, in what context, and with what consequences.
2. Write verifiable statements.
    Any significant statement must be based on data, observations, incidents, user feedback, metrics, or explicitly marked as a hypothesis.
3. Separate facts, hypotheses, and opinions.
    Facts are confirmed. Hypotheses require verification. Opinions are acceptable only as input signals, but not as a basis for starting development.
4. Capture impact, not just inconvenience.
    Problem should be associated with measurable damage: loss of time, money, users, quality, reliability, security, delivery speed, or process manageability.
5. Use specific context.
    Do not write “users are uncomfortable” or “the system is slow.” You need to specify segment, scenario, frequency, scale, affected systems, and conditions under which problem manifests.
6. Do not expand scope unnecessarily.
    Document should narrow area of discussion. If problem concerns one scenario, region, segment, or service, do not turn it into a general platform modernization.
7. Keep desired state non-normative.
    Describe in one short paragraph what becomes true for affected audience when problem no longer exists. Do not list system behaviors or checks.
8. Leave contentious issues open, rather than masking them.
    If cause of problem is unknown, data is insufficient, or there are multiple interpretations, this should be explicitly recorded as an open question, rather than forcing document into a ready-made conclusion.
9. Write for decision-making.
    Document should help answer: is it worth continuing, what needs to be checked, who is affected, how important problem is, and what next rational step is.
10. Keep document short.
    A Problem Statement is not a PRD or a technical design. If document grows to include UX, API, architecture, acceptance criteria, and roadmap, it has already gone beyond its purpose.
11. `<id_guidelines>` MUST NOT be used.
12. Stop before requirements.
    A Problem Statement describes an observed problem and its impact. It MUST NOT define required system behavior, scenarios, acceptance criteria, implementation scope, architecture, components, APIs, files, tests, or delivery tasks.
13. Do not copy proposed solutions from source materials.
    When an input document contains requirements or implementation ideas, use them only to understand context. Defer them to PRD or Technical Solution.
14. Define problem boundary, not implementation scope.
    State which observed situation, audience, and consequences belong to problem. Do not decide which changes are in scope or out of scope.
15. Omit unsupported sections.
    Do not invent metrics, evidence, impact, assumptions, or questions to fill template.
16. Apply acceptance-criterion test.
    If a statement can be used unchanged as an acceptance criterion or implementation task, it does not belong in Problem Statement.
</problem_statement_guidelines>
-->

<problem_statement_template>
# Problem Statement

## Context
<!--Brief background needed to understand observed situation-->

## Observed Problem
<!--One concise statement of what is wrong, for whom, and under which conditions-->

## Affected Audience
<!--People, teams, or systems that experience problem-->

## Evidence
<!--Observations, incidents, data, or source references that show problem exists-->

## Impact
<!--Consequences caused by problem-->

## Current State
<!--Relevant behavior as it exists now, without proposed changes-->

## Desired State
<!--One short description of state in which problem no longer exists-->

## Problem Boundary
<!--Situation and audience covered by this problem. This is not implementation scope-->

## Assumptions
<!--Unverified facts that affect whether problem exists or how serious it is-->

## Open Questions
<!--Missing facts needed to understand problem, not decisions about a solution-->
</problem_statement_template>