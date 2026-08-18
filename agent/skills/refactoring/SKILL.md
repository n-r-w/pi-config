---
name: refactoring
description: How to refactor code effectively, including techniques and best practices.
---

<refactoring>
  <big_picture_principles>
    1. Focus on the goal, not the details. Always evaluate your actions from the perspective of the end goal.
    2. Present the user with the big picture first, and only delve into details if they want to.
    3. CRITICAL: Remember that despite the need to focus on the goal and the big picture, you MUST study the details, as without understanding the details it is impossible to develop quality ideas and solutions.
  </big_picture_principles>

  <guidelines>
    1. **No Workarounds**: Avoid forwarding functions, type aliases, or adapter layers to preserve old call sites.
    2. **High Level Vision**: Any local change should be part of a global strategy.
    3. **Fast Does Not Mean Good**: Quick fixes can lead to technical debt and architectural violations.
      Always assess the consequences. It's better to change more code than to leave dirty workarounds or local solutions.
    4. **No Workarounds**: Do not leave temporary solutions.
  </guidelines>
</refactoring>