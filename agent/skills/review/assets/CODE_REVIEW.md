<code_review name="Additional Requirements For Code Review">
  1. Verify that no redundant code was added during development and debugging. Code cleanliness and maintainability are crucial.
  2. Ensure that ALL temporary fixes, stubs, or debugging constructs have been removed from final code, if they not part of future implementation plan.
  3. Insist to split huge functions into smaller, more manageable ones.
  4. STRICTLY verify `<commenting>` rules from `coding-rules` skill.
  5. `Code slop` review area is a collection of low-value, redundant, or structurally harmful code that does not add functionality. `Code Slop` review area is CRITICAL and BLOCKING because such issues will be REJECTED by user
  6. It is important to check for presence of comments for functions, classes, variables, and complex logical blocks. Comments should be informative, up-to-date, and follow established commenting rules in project.
  7. Unit tests are NEVER a criterion of truth, as they can reinforce invalid behavior.
  8. MUST STRICTLY FOLLOW `<realism_assessment>` rules.
  9. Verify that code does not introduce unnecessary complexity or overengineering. KISS and YAGNI principles are followed.
  10. Check that during implementation, parts of code that are not related to task were not affected. Any changes in code must be justified and documented (but should not contain references to ticket numbers, plan stages, etc. unrelated to code itself).
</code_review>