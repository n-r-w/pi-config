<!--
<problem_statement_guidelines>
1. Describe the problem, not the solution.
    Phrases like “build a service / dashboard / API / button” are not a problem statement. First, you need to capture what exactly is not working, for whom, in what context, and with what consequences.
2. Write verifiable statements.
    Any significant statement must be based on data, observations, incidents, user feedback, metrics, or explicitly marked as a hypothesis.
3. Separate facts, hypotheses, and opinions.
    Facts are confirmed. Hypotheses require verification. Opinions are acceptable only as input signals, but not as a basis for starting development.
4. Capture impact, not just inconvenience.
    The problem should be associated with measurable damage: loss of time, money, users, quality, reliability, security, delivery speed, or process manageability.
5. Use specific context.
    Do not write “users are uncomfortable” or “the system is slow.” You need to specify the segment, scenario, frequency, scale, affected systems, and conditions under which the problem manifests.
6. Do not expand the scope unnecessarily.
    The document should narrow the area of discussion. If the problem concerns one scenario, region, segment, or service, do not turn it into a general platform modernization.
7. Define the desired outcome without tying it to implementation.
    A good outcome describes a new state: what will become possible, faster, more reliable, or cheaper. It should not dictate architecture or a specific feature.
8. Define success criteria in advance.
    Before development, you need to understand by which metrics it will be clear that the problem has indeed been solved. Without this, it is impossible to distinguish between “we did something” and “we solved the problem.”
9. Explicitly capture constraints.
    Constraints on deadlines, compatibility, security, data, regulations, resources, and dependencies should be visible before choosing a solution.
10. Leave contentious issues open, rather than masking them.
    If the cause of the problem is unknown, data is insufficient, or there are multiple interpretations, this should be explicitly recorded as an open question, rather than forcing the document into a ready-made conclusion.
11. Write for decision-making.
    The document should help answer: is it worth continuing, what needs to be checked, who is affected, how important the problem is, and what the next rational step is.
12. Keep the document short.
    A Problem Statement is not a PRD or a technical design. If the document grows to include UX, API, architecture, acceptance criteria, and roadmap, it has already gone beyond its purpose.
13. `<id_guidelines>` MUST NOT be used.
</problem_statement_guidelines>
-->

<problem_statement_template>
# Problem Statement

## Context
<!--Brief background: where and why the topic arose-->

## Problem Statement
<!--One statement of the problem without a solution-->

## Who is affected
<!--Users, teams, clients, systems-->

## Evidence
<!--Metrics, tickets, interviews, incidents, logs, observations-->

## Impact
<!--User / business / operational / technical impact-->

## Reproduction Steps
<!--How to reproduce the problem, if applicable-->

## Current State
<!--How it works now-->

## Desired Outcome
<!--What state we want to achieve-->

## Success Metrics
<!--How we will measure that the problem is solved-->

## Scope
<!--What is included-->

## Out of Scope / Non-Goals
<!--What is not included and what we are not trying to achieve-->

## Constraints
<!--Technical, organizational, legal, and time constraints-->

## Assumptions
<!--Unproven but important assumptions-->

## Open Questions
<!--What needs to be clarified-->
</problem_statement_template>