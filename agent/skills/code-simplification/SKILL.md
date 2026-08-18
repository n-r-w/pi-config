---
name: code-simplification
description: Code simplification standards and best practices. Use BEFORE ANY activity related to code simplification.
---

<code_simplification_standard>
  <goal> Ensure that all code, scripts, makefiles, taskfiles meet the highest standards of elegance and maintainability while preserving its complete functionality</goal>
  <refinement_guidelines>
    <rule name="Apply Standards">
      Follow ALL the established coding standards including:
        1. General coding standards
        2. Programming language standards
        3. Project specific standards
    </rule>
    <rule name="Ensure Clarity">
      Simplify code, script, makefile, and taskfile structure by:
        1. Using modern language constructs and libraries
        2. Reducing unnecessary complexity and nesting
        3. Eliminating redundant code and abstractions
        4. Improving readability through clear variable and function names
        5. Consolidating related logic
        6. Split large functions into smaller, focused ones (avoid splitting atomic business logic)
        7. Align new code with existing code style and patterns
        8. Eliminating unnecessary exports/public visibility (e.g. DTOs)
        9. Remove redundant tests that don't add value
        10. Change comments to explain "why" instead of "how"
        11. Removing from comments:
          1) Non-keyboard symbols like `→`, `—`, etc.
          2) Unnecessary explanations of obvious things
          3) References to plans, tasks, code reviews, memories, etc. that are not related to the code context
        12. Choose clarity over brevity - explicit code is often better than overly compact code
    </rule>
    <rule name="Special Libraries">
      Analyze which well-known libraries and frameworks can help simplify the code, and use them if it does not contradict project standards.
    </rule>
  </refinement_guidelines>

  <constraints>
    1. All simplifications MUST preserve the original functionality and behavior of the code.
    2. MUST NOT simplify code that is not part of the current task scope (eg. current ticket, PR, or user request). Only refine code that has been recently modified or touched in the current session, unless explicitly instructed to review a broader scope.
    3. NEVER remove comments like "TODO", "FIXME", "NOTE", etc. until they are resolved. DON'T hide important information!
    4. Avoid over-simplification that could:
        1. Reduce code clarity or maintainability
        2. Create overly clever solutions that are hard to understand
        3. Combine too many concerns into single functions or components
        4. Remove helpful abstractions that improve code organization
        5. Prioritize "fewer lines" over readability (e.g., nested if, dense one-liners)
        6. Make the code harder to debug or extend
  </constraints>

  <golang_specific>
    <avoid_to_simplify>
      Struct initialization. E.g.: `Service{ID: 1, Data: ""}` -> `Service{ID: 1}` or `new(Service)`. Why: explicit initialization of all fields ensures that nothing is forgotten and does not trigger `exhaustruct` linter warnings. External SDK types can use `//nolint:exhaustruct` if they have many fields.
    </avoid_to_simplify>
    <recommend_to_simplify>
      1. `github.com/samber/lo` instead of custom helpers.
      2. `github.com/google/uuid` instead of custom UUID generation code.
      3. `github.com/caarlos0/env/v11` instead of custom environment variable parsing.
    </recommend_to_simplify>
  </golang_specific>
</code_simplification_standard>