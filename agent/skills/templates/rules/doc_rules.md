<documents_slop>
  <guidelines>
    1. Document Slop is a collection of low-value, redundant, or structurally harmful content in documentation that does not add value but increases complexity, cognitive load, and technical debt.
    2. This section contains list of document slop categories in format "issue ==> **how to fix**"
  </guidelines>
  <categories>
    <category name="Temporary References">
      References to any temporary artifacts (e.g. collaboration desk ids, temporary files, etc.) ==> **REMOVE**
    </category>
    <category name="Statements without Evidence">
      Phrases like "confirmed", "verified", "supported", "validated", "proven", etc. DO NOT INCREASE the value of the document, but only create meaningless NOISE.
      Instead of such phrases, you MUST add specific evidence (e.g. links to documents, code, etc.) that support these statements. ==> **REMOVE such phrases and add specific evidence if possible**
    </category>
    <category name="Duplicate Information">
      The same information presented in multiple places in the document (even if in different wording) ==> **REMOVE duplicates and keep only one version of the information. If necessary, add links to the main description instead of duplication**
     </category>
    <meta_information>
      Information that is not related to the content of the document, but rather to its formatting, style, structure, etc. For example, an explanation next to a diagram about why a particular style was chosen ==> **REMOVE such meta-information**
    </meta_information>
  </categories>
</documents_slop>

<anti_ambiguity>
  <usage>
    1. **Authoring:** apply while creating or updating specifications, architecture, plans, tickets, ADRs, contracts, research, glossaries, and similar artifacts.
    2. **Review:** apply before completion, delivery, approval, or implementation handoff. Fix ambiguity when modification is authorized; otherwise propose exact replacements.
  </usage>

  <rules>
    1. Review normative statements and explanatory text that affects interpretation of scope, requirements, behavior, ownership, contracts, errors, constraints, or verification.
    2. Wording is ambiguous when it permits materially different interpretations or has no defined pass/fail observation. Missing requirements, contradictions, unresolved decisions, evidence gaps, and factual errors are not ambiguity and MUST be reported separately.
    3. MUST NOT invent behavior or change approved scope, requirements, decisions, business behavior, terminology, or requirement strength to make wording specific. If evidence cannot determine one replacement, report the actual blocker.
    4. Each normative statement MUST define one meaning. Include the subject, applicable condition, required behavior, observable outcome, and closed value set when each affects interpretation.
    5. Replace `current`, `existing`, `same`, `unchanged`, and equivalent references with the owning contract, symbol, fields, validations, and observable outcomes that must be preserved.
    6. Replace broad qualifiers such as `supported`, `relevant`, `complete`, `valid`, or `appropriate` with a closed list, decision rule, or one named owning definition.
    7. For errors, define source condition, incoming status, outgoing status, external result, retry behavior, and allowed follow-up calls. When errors cross several boundaries, keep one mapping with stable IDs and define behavior for unlisted statuses.
    8. For non-functional requirements, define an observable signal, observation point, allowed values, threshold or expected result, and verification method. For logging or telemetry restrictions, list prohibited data and sinks. Remove a qualitative requirement when another requirement already makes the behavior verifiable.
    9. Shared behavior MUST have one owning definition. Group repeated violations by root cause and apply one correction pattern to every affected artifact.
    10. State essential behavior in the artifact. Code references and links are evidence, not substitutes for required behavior.
    11. Verify referenced terminology, symbols, fields, statuses, enum values, files, sections, and links against current evidence.
    12. Completion requires no unresolved ambiguity in scope and no change to approved meaning. MUST follow KISS and YAGNI; style-only edits and unnecessary implementation detail are prohibited.
  </rules>

  <workflow>
    <authoring>
      1. Define terms and normative owners.
      2. Draft content using rules 4–10.
      3. Verify references using rule 11, scan for ambiguity, and apply rule 12.
    </authoring>
    <review>
      1. Determine the reviewed artifacts, affected references, and modification permission.
      2. Detect wording whose meaning depends on reader inference; keywords alone are not findings when nearby text or one named owner defines exact meaning.
      3. Classify findings using rules 2–3 and group repeated violations using rule 9.
      4. Apply rules 4–10 when modification is authorized. Otherwise provide the original wording, materially different interpretations, exact replacement, and evidence or blocker.
      5. Verify references using rule 11, confirm approved meaning did not change, and repeat the ambiguity scan until rule 12 passes.
    </review>
  </workflow>
</anti_ambiguity>