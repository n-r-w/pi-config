<extraction_report_template>
<!--
<extraction_guidelines>
   1. If the topic is broad, prioritize: problem framing → evidence → conclusion → next steps.
   2. Use IDs only when required by <id_guidelines>. Do not infer this need from the item type or template section.
   3. Every assigned ID must be unique within the document.
   4. All relevant information must be extracted with minimal interpretation, ensuring a comprehensive and unbiased report.
</extraction_guidelines>
-->

# {Report Title}
{Briefly restate the user's request and define the scope of the analysis.}

## Summary
{Provide a concise summary of the key findings from the information extraction process (2-3 sentences)}

## Findings
- {Finding name}
   - Location: {file path and line numbers; URL; etc.}
   - Description: {A brief description of the finding, including any relevant details that were extracted.}
   - Dependencies: {List any dependencies or related findings that are relevant to this finding.}
   - Patterns: {List any patterns or commonalities observed across findings, if applicable.}
   - Confidence Level: {Provide a confidence level in percentage based on the completeness and reliability of the extracted information.}

## Assumptions
- {Justification (why); verification (how/when)}
- ...

## Open Questions

### {QST-id}: {brief description}
  - {impact}
  - {what should answer look like}
  - {what has been done to find answer}
  - {how/when it will be resolved}

### ...
</extraction_report_template>