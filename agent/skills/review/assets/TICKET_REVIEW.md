<plan_review name="Additional Requirements For Ticket Review">
  <guidelines>
    Verify that document meets the following criteria:
    1. [ ] Ticket is realistic and executable. Verification is done based on documentation, code, web resources, and other relevant sources.
    2. [ ] Ticket follows TDD principles: has explicit RED-GREEN-REFACTOR steps or decision why not applicable.
    3. [ ] Ticket does not introduce unnecessary complexity or overengineering. KISS and YAGNI principles are followed.
  </guidelines>
  <clarity_rules>
    1. **GOLDEN RULE**: No ambiguities. Any ambiguity and vagueness in the formulations leads to a violation in implementation.
    2. Each statement is recorded in a single instance, without repetitions. Instead of repetitions, references to a single place of description are used.
    3. There are no statements that can be interpreted in two ways. No "if", "maybe", "sometimes", etc.
    4. Any mentions of constants, enums, etc. must contain a list of all possible values.
    5. Any abstraction must be disclosed to the level of content.
  </clarity_rules>
</plan_review>