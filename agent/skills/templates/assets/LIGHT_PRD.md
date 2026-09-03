<!--
<light_prd_guidelines>
  1. Prefer bullets over long paragraphs; keep each bullet to 1–2 sentences.
  2. Do not invent facts. If ambiguity cannot be resolved, write "TBD" and add an item to Open Questions.
  3. Focus on WHAT and WHY, not HOW. Do not include implementation plans, component-level architecture, or code.
  4. The document must be self-contained and understandable without external context. Links may provide extra detail, but must not replace essential explanation.
  5. Be specific. Avoid vague placeholders such as "if", "may", "optional", "not confirmed", or "mandatory model", unless the item is an explicit open question.
  6. `<id_guidelines>` MUST NOT be used.
  7. Any ambiguities must be either resolved with decisions or moved to Open Questions. No "ifs" outside of Open Questions.
  8. This is not a full specification, but only a light version for quick idea evaluation.
  9. Requirements without a brief justification are not accepted.
</light_prd_guidelines>
-->

<light_prd_report_template>
# Idea: {Initiative / Feature Name}

## Definitions
<!-- Key definitions and abbreviations -->

## Context and Problem
<!-- What is the situation now? What problem are we trying to solve? -->
<!-- If there is an external document, then instead of a description here, you need to provide a link to the document -->

## Goal
<!-- What is the desired outcome? -->

## Scenarios
<!-- What are the key scenarios that will be impacted by this change? -->

## Scope and Non-Scope
<!-- What is in scope and what is out of scope for this change? -->

## Requirements
<!--
Key requirements for the change. Each requirement should be atomic, concrete, and unambiguous.
Do not try to capture ALL details. Focus on the minimally necessary set of requirements.
MUST provide a BRIEF justification for each requirement.
Recommended format:
  - [Requirement]
    Justification: [Brief explanation of why this requirement is necessary]
-->

## Open Questions
<!-- List any open questions that need to be resolved before the change can be implemented -->

## Technical Supplement
<!--
If technical details need to be recorded during the specification creation process, they MUST be ONLY here.
The other sections of the specification should contain only business-oriented details.

IMPORTANT:
 1. In most cases, the specification should EXCLUDE technical details, so adding them to the specification should be an exception, not the rule.
 2. It is PROHIBITED to reference items from the technical supplement from other sections of the specification. If connections are required, then only the technical supplement can refer to other sections of the specification, but not vice versa.
 3. Choose the structure and format of the technical supplement that best suits the specific case.
-->

## References
<!-- List any relevant references, such as related documentation, web resources, etc. -->

</light_prd_report_template>