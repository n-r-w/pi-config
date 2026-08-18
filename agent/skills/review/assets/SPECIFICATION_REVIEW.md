<specification_review name="Additional Requirements For Specification Review">
  <guidelines>
  Verify that document meets the following criteria:
    1. [ ] Goal is one sentence and measurable (Outcome).
    2. [ ] Scope is clearly split into “in” and “out”.
    3. [ ] All unknowns are captured as explicit questions in one place.
    4. [ ] Acceptance criteria are verifiable (test/metric/observation).
    5. [ ] No implementation plan, no “how-to build”, no code.
    6. [ ] Assumptions are stated with justification and verification.
    7. [ ] Open questions are listed with impact and resolution path.
    8. [ ] Document is self-contained and understandable without external context.
    9. [ ] All items have unique IDs.
    10. [ ] Goals and metrics are realistic and achievable. Verification is done based on documentation, code, web resources, and other relevant
    sources.
    11. [ ] Specification does not introduce unnecessary complexity or overengineering. KISS and YAGNI principles are followed.
    12. [ ] No overspecification. Keep only critical requirements.
  </guidelines>
  <clarity_rules>
    1. **GOLDEN RULE**: No ambiguities. Any ambiguity and vagueness in the formulations leads to a violation in implementation.
    2. Each statement is recorded in a single instance, without repetitions. Instead of repetitions, references to a single place of description are used.
    3. There are no statements that can be interpreted in two ways. No "if", "maybe", "sometimes", etc.
    4. Any mentions of constants, enums, etc. must contain a list of all possible values.
    5. Any abstraction must be disclosed to the level of content.
  </clarity_rules>
  <to_be_contract_rules critical="true">
    1. WHEN the task is to specify or review a target API or data contract, THE agent SHALL derive the contract from the target boundary of responsibility, not from the union of legacy consumer branches.
    2. WHEN logic is extracted into a dedicated service, THE agent SHALL treat the extracted service boundary as authoritative. Fields required by that boundary SHALL be mandatory in the target contract even if some current consumers do not always read them.
    3. WHEN writing contract or data-model text, THE agent SHALL avoid hedge wording such as "if", "may", "optional", "not confirmed", or "mandatory model" unless the item is an explicit open question with its own ID.
    4. WHEN ambiguity cannot be removed from a contract using the declared boundary, verified code, and existing requirements, THE agent SHALL ask one focused question with concrete options and a recommendation.
  </to_be_contract_rules>
</specification_review>