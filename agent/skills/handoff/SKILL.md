---
name: handoff
description: Compact current conversation into a handoff document for another agent to pick up.
disable-model-invocation: true
---

1. Write a handoff document in ASD-STE100 (Simplified Technical English), summarising current conversation so a fresh agent can continue work. Save to temporary directory of user's OS - not current workspace.
2. Include a "suggested skills" section in document, naming which skills next agent should use.
3. MUST NOT include current workflow state, because another agent may use another workflow.
4. MUST NOT use specific output template.
5. Do not duplicate content already captured in other artifacts (specs, plans, ADRs, issues, commits, diffs, messages). Reference them by path, URL or ID instead.
6. Redact any sensitive information, such as API keys, passwords, or personally identifiable information.
7. Do not repeat contents of handoff in answer, just provide created file path.
8. MUST NOT talk about agent, only about task to be performed.
9. Handoff MUST be self-contained for a fresh session that has no conversation history.
10. Define every conversation-local identifier, abbreviation, implementation unit, phase label, and nickname before first use. Examples include U1, U2, TSK-01, and local branch names.
11. MUST NOT include workflow runtime data such as active workflow ID, active workflow stage, workflow transitions, or workflow session IDs. Include implementation plan needed to continue task, including unit definitions, execution order, dependencies, and completion status of each unit.
12. No-duplication rule does not permit replacing essential orientation with references. Include minimum context needed to understand each referenced artifact and its relevance.
13. Before saving, verify that a fresh reader can identify goal, completed work, current work, next action, dependencies, constraints, verification requirements, and every local identifier without access to prior conversation.
14. References MUST include enough context to explain why referenced artifact is needed. If a reference may be unavailable, include facts required to continue.
15. Include "user-facing language" section to specify language used in last substantive assistant response before this skill was invoked.
16. If user passed arguments, treat them as a description of what next session will focus on and tailor doc accordingly.