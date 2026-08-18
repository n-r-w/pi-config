<analysis_report_template>
<!--
<analysis_guidelines>
  1. Universal: works for investigating a bug, incident, performance issue, product idea, architectural concern, or any technical/business question.
  2. If the topic is broad, prioritize: problem framing → evidence → conclusion → next steps.
  3. Do not invent facts. Every claim must contain a reference to specific evidence.
  4. If information is missing, write "TBD" and add a question in "Open Questions".
  5. Separate observations from interpretations. Do not mix them in one bullet.
  6. Do not include code dumps. If necessary, include at most 10 lines, but prefer referencing locations (file/section/PR/link).
  7. Ensure conclusions follow from evidence; if uncertain, say so and propose how to validate.
  8. Document MUST be fully self-contained and understandable without external context. Use links for details, not as a substitute for explanation.
  9. MUST follow <id_guidelines> for all items in the document.
  10. All items have unique IDs.
  11. MUST NOT mix different languages in the same document (e.g. English and Russian). Translate section names if using a non-English language.
  12. Any ambiguities must be either resolved with decisions or moved to Open Questions. No "ifs" outside of Open Questions.
</analysis_guidelines>
-->

# Analysis Results: {Topic}

{why this analysis was conducted; what question it aimed to answer; what decision it aimed to inform}

## Scope
{scope; out of scope}

## Key definitions and abbreviations

- {id}: {Definition}. {Explanation of the term or abbreviation.}
- ...

## Executive Summary
{question/problem statement; conclusion; impact; recommended next step}

## Background and Context
{brief history; relevant technical/product context; why this question is important}


## Method and Data Sources
{what you examined and how; methods; data sources with links/locations; limitations of data}

## Observations

- {id}: {fact; what was observed; what data was examined}
- ...

## Analysis and Interpretations

- {id}: {interpretation of the observation; what it might mean; how it relates to the question}
- ...

## Hypotheses and Tests <!--If not fully confirmed-->

- {id}: {testable claim; what it means if true; what it means if false; how to test; expected signal; falsification signal; effort/cost}
- ...

## Options and Trade-offs

- {id}: {pros; cons; risks; cost/effort; when to choose}
- ...

## Recommendation
{recommended options; rationale; preconditions; release/rollback/exit strategy if applicable and explicitly requested by the user; next steps}

## Action Plan <!--Prioritized-->

- {id}: {priority; action; success criteria; links to tickets/PRs/docs}
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
</analysis_report_template>