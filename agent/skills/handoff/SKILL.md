---
name: handoff
description: Compact current conversation into a handoff document for another agent to pick up.
disable-model-invocation: true
---

1. Write a handoff document in ASD-STE100 (Simplified Technical English), summarising current conversation so a fresh agent can continue work. Save to temporary directory of user's OS - not current workspace.
2. Include a "suggested skills" section in document, naming which skills next agent should call Skill tool for.
3. MUST NOT include current workflow state, because another agent may use another workflow.
4. MUST NOT use specific output template.
5. Do not duplicate content already captured in other artifacts (specs, plans, ADRs, issues, commits, diffs, messages). Reference them by path, URL or ID instead.
6. Redact any sensitive information, such as API keys, passwords, or personally identifiable information.
7. If user passed arguments, treat them as a description of what next session will focus on and tailor doc accordingly.
8. Do not repeat contents of handoff in the answer, just provide created file path.
