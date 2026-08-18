<plan_review name="Additional Requirements For Plan Review">
  <guidelines>
    Verify that document meets the following criteria:
    1. [ ] Plan is split into phases; each has goal, work, deliverables, exit criteria, risks.
    2. [ ] Plan follows TDD principles: each phase has explicit RED-GREEN-REFACTOR steps or decision why not applicable.
    3. [ ] Each phase delivers a verifiable increment or reduces a top risk.
    4. [ ] Delivery/release strategy is defined (flags/canary/dark launch/dual-run where applicable).
    5. [ ] Release and rollback plan created ONLY if explicitly requested by user.
    6. [ ] Test strategy includes unit, integration, E2E, load/stability with pass criteria.
    7. [ ] Data migration approach is stated (if applicable) with risks.
    8. [ ] Dependencies/blockers are listed; no invented dates or staffing.
    9. [ ] Definition of Done covers docs, observability, operational readiness, SLO/stability.
    10. [ ] No code samples or low-level implementation steps are included.
    11. [ ] TBDs are explicit and localized (not hidden).
    12. [ ] Document is self-contained and understandable without external context.
    13. [ ] Tests provide REAL value, and not just be a chore to increase coverage.
    14. [ ] All items have unique IDs.
    15. [ ] Plan is realistic and executable. Verification is done based on documentation, code, web resources, and other relevant sources.
    16. [ ] Critical risk: Making every phase testable does not lead to inventing unnecessary logic and tests that will not provide real value when looking at the task as a whole. Each item in the `Exit criteria` section is analyzed from this point of view and excessive verification requirements are eliminated, or phases are combined to avoid such situations. Plan MUST provide justification in `Decomposition Justification` section.
    17. [ ] Plan does not introduce unnecessary complexity or overengineering. KISS and YAGNI principles are followed.
    18. [ ] No overspecification. Keep only critical requirements.
  </guidelines>
  <clarity_rules>
    1. **GOLDEN RULE**: No ambiguities. Any ambiguity and vagueness in the formulations leads to a violation in implementation.
    2. Each statement is recorded in a single instance, without repetitions. Instead of repetitions, references to a single place of description are used.
    3. There are no statements that can be interpreted in two ways. No "if", "maybe", "sometimes", etc.
    4. Any mentions of constants, enums, etc. must contain a list of all possible values.
    5. Any abstraction must be disclosed to the level of content.
  </clarity_rules>
</plan_review>