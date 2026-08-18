---
name: tdd
description: Test Driven Development (TDD) guidelines. Use BEFORE ANY activity related to coding (analysis, review, planning, implementation, etc.).
---

<tdd_rules title="Test Driven Development (TDD) Rules">
  <exceptions>TDD is mandatory unless the user explicitly requests an exception</exceptions>

  <iron_law>
    1. No production code without a failing test first except `<when_not_to_use>`.
    2. Always prove RED/GREEN by running tests (avoid relying on cached results).
    3. Prefer mocks over real code paths to isolate units; Use real/integration paths only when validating a contract or end-to-end behavior.
    4. Tests MUST provide REAL value, and not just be a chore to increase coverage. If you realize that testing is pointless in a given situation, DOCUMENT THAT. DON'T allow silent solutions.
    5. BEFORE planning new tests, check if relevant tests already exist and plan to update them if possible.
    6. For each new or updated test, EXPLICITLY specify:
      1) Purpose of the test
      2) Inputs and expected outputs (briefly)
      3) Edge cases being tested
      4) Dependencies on other tests or components
    7. If tests fail on RED phase due to timeout instead of expected assertion failure, this is NOT considered valid RED.
    8. MUST NEVER use TDD for writing tests that check for the absence of functionality! Tests MUST check actual behavior, not the absence of it.
  </iron_law>

  <why_tdd>
    1. **Goal:** Reduce uncertainty by turning requirements/bugs into executable checks.
    2. **Pros:**
      1) Faster feedback on behavior regressions.
      2) Safer refactors (tests become a harness).
      3) Better modularity (tests push clear boundaries).
    3. **Risks:**
      1) Risk of over-mocking (tests that pass but miss real behavior).
      2) Poorly chosen tests can lock in bad design.
  </why_tdd>

  <when_to_use>
    **Use TDD for:**
      1. New behavior
      2. Bug fixes
      3. Behavior changes
  </when_to_use>

  <when_not_to_use>
    CRITICAL: ALWAYS critically evaluate the need for tests first. Otherwise, your work won't be accepted, and you'll waste time.

    **Don't use TDD for:**
      1. Refactors with no behavior change.
      2. Documentation-only work.
      3. Formatting-only work.
      4. When tests provide no real value (e.g., checking constant values, trivial getters/setters, etc.).
      5. Scenarios where TDD is meaningless in principle (for example, writing configuration files).
      6. When removing functionality, it is meaningless to write tests that check that it is no longer present. This will be a garbage test that carries no value.
      7. Checks of mutable content. For example, web page content, etc. You need to check behavior, not the content itself.
  </when_not_to_use>

  <how_to_do_it>
    1. **RED:** Write/adjust a behavioral test; run it; confirm it fails for the expected reason.
    2. **GREEN:** Implement the minimum change to pass; rerun the same test(s).
    3. **REFACTOR:** Improve code structure without changing behavior; keep tests green.
    4. **VERIFY:** Run the broader relevant suite (fast + targeted + end-to-end as needed).
  </how_to_do_it>

   <red_compile_setup>
     1. RED tests MUST compile and run before their failure is accepted as RED.
     2. If a missing target symbol, file, or entry point prevents compilation, add the smallest compile-only stub before RED.
     3. Compile-only stubs MUST NOT implement behavior, copy legacy code, read runtime config, call external systems, or add compatibility paths.
     4. Missing dependencies, unresolved imports, invalid module resolution, and package setup issues are blockers, not valid RED failures.
     5. If resolving such a blocker requires changing package manifests, lock files, module resolution, or install state, stop and ask the user first.
     6. If the code does not compile, this is NOT CONSIDERED RED. In this case, add a minimal stub.
   </red_compile_setup>

  <bug_fixing>
    CRITICAL: NEVER write a fix BEFORE writing a test that proves the bug EXISTS. This is the most common mistake that breaks TDD and leads to wrong fixes.
      1. **RED:** Add a test asserting the buggy behavior is rejected; run the test and confirm it fails for the expected reason.
      2. **GREEN:** Implement the minimal fix; rerun the same test (pass).
      3. **REFACTOR:** Clean up if needed; keep the test green.
  </bug_fixing>

  <review_followup>
    If after a code review you receive feedback that requires changes to the code, you MUST follow TDD rules for those changes. This means:
      Ask yourself: "Why wasn't this issue caught by existing tests?". If it's technically possible to modify/add a test that would catch this issue, do that BEFORE fixing the code. Run the test and confirm it fails for the expected reason.
  </review_followup>

  <example>
    **Bug fix sketch:**
      1. **RED:** Add a test asserting the buggy behavior is rejected; run the test and confirm it fails for the expected reason.
      2. **GREEN:** Implement the minimal fix; rerun the same test (pass).
      3. **REFACTOR:** Clean up if needed; keep the test green.
  </example>о
</tdd_rules>
