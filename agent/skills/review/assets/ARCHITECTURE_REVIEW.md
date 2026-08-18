<architecture_review name="Additional Requirements For Architecture Review">
  <guidelines>
  Verify that document meets the following criteria:
    1. [ ] System boundaries and surrounding context are clear (consumers + neighbor systems).
    2. [ ] `One screen` overview exists: core idea + primary data flow.
    3. [ ] At least one failure mode is described with expected behavior.
    4. [ ] Key components are listed.
    5. [ ] Interfaces/contracts are named (APIs/resources/events) without unnecessary field detail.
    6. [ ] Idempotency/ordering/deduplication decision is stated (if applicable).
    7. [ ] NFR covers reliability, scaling, security, observability.
    8. [ ] At least 2 ADR-lite decisions with alternatives and rationale are included.
    9. [ ] Architecture risks/trade-offs are listed; no implementation plan is present.
    10. [ ] Components are designed according to clean architecture principles.
    11. [ ] Interfaces/Contracts are designed according to SOLID principles.
    12. [ ] The same component is not both provider and consumer of the same interface/contract.
    13. [ ] Interfaces defined in the place of usage, not in the place of implementation.
    14. [ ] No phases, timelines, step-by-step build instructions, or code.
    15. [ ] Assumptions are stated with justification and verification.
    16. [ ] Open questions are listed with impact and resolution path.
    17. [ ] Document is self-contained and understandable without external context.
    18. [ ] All items have unique IDs.
    19. [ ] Architecture is realistic and implementable. Verification is done based on documentation, code, web resources, and other relevant sources.
    20. [ ] Architecture does not introduce unnecessary complexity or overengineering. KISS and YAGNI principles are followed.
    21. [ ] No overspecification. Keep only critical requirements.
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
</architecture_review>