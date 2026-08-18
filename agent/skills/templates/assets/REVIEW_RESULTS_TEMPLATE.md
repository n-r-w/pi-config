<review_report_template>
<!--
<review_guidelines>
  1. Use these guidelines for reviewing both code and documentation.
  2. Do not invent evidence. Every finding must cite a specific artifact location (file, section, line, or link), or be explicitly marked "General" with rationale.
  3. Use only these severities: Blocker, Major, Minor, Suggestion.
  4. All items must follow <id_guidelines> and have unique IDs.
  5. Recommendations must be actionable, specific, and unambiguous. Avoid vague wording such as "consider" unless concrete options are provided.
  6. If required information is missing, write "TBD" and add the corresponding question to "Open Questions".
  7. Do not include code snippets longer than 10 lines. Prefer describing the required change over pasting code. Do not include full file dumps.
  8. Conclusions must remain logically consistent: if any Blocker exists, the outcome cannot be "Approve".
  9. Report issues only. Do not include strengths or positive observations.
  10. The review must be fully self-contained and understandable without external context. Links may support details, but must not replace explanation.
  11. Any ambiguities must be either resolved with decisions or moved to Open Questions. No "ifs" outside of Open Questions.
  12. Always critically assess the realism of identified issues. Do not waste time fixing problems that are highly unlikely to occur.
    If the scenario is related to API calls, make sure it is achievable through the API and not just artificially reproduced through unit tests.
</review_guidelines>
  13. ALL sections in Findings are mandatory: Location, Issue, Impact, Scenario, Assessing Realism, Recommendation, Verification.
-->

# Review Results: {Artifact / Change Set Name}
{scope; out of scope; intent}

## Key definitions and abbreviations

- {id}: {Definition}. {Explanation of the term or abbreviation.}
- ...

## Summary
Outcome: ✅ Approve | 🟠 Approve with changes | 🔄 Request changes | 🟥 Reject

Overall assessment (1–3 bullets):
- <id> {most important observation}
- ...

## Issues Overview
<!-- Explanation in simple and understandable language, aimed at an analyst or business user, not a programmer.
This includes what's broken in the business logic, what logic/formulas should exist, and which are actually hardcoded, etc. -->
- **{level indicator} {id}**: {brief description of the issue}
- ...

## Findings <!-- Actionable items; Sort by severity, then by priority -->

### Blocker
#### ⛔ {id}
  - Location: {file/section/line/link}
  - Issue: {description of the problem}
  - Impact: {description of the consequences}
  - Scenario: {scenario demonstrating the process of problem occurrence}
  - Assessing Realism: probability - {High/Medium/Low}; {a description of the feasibility of the scenario, with justification for its likelihood of occurrence, including any assumptions made.}
  - Recommendation: {specific, actionable recommendation to fix the issue}
  - Verification: {how to verify the fix, including any necessary steps or tests}

### Major
#### ⚠️ {id}
  - Location: {file/section/line/link}
  - ...

### Minor
#### ℹ️ {id}
  - Location: {file/section/line/link}
  - ...

### Suggestion
#### 💬 {id}
  - Location: {file/section/line/link}
  - Issue: {description of the problem}
  - Recommendation: {specific, actionable recommendation to fix the issue}
  - Why: {why the recommendation is being made}
  - Verification: {how to verify the fix, including any necessary steps or tests}

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

## Next Steps
- {id}: {action linked to finding ids}
- ...

## References
- {id}: {reference (file/URL/standard, etc.)} - {one-line description}
- ...
</review_report_template>