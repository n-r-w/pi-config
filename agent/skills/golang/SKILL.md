---
name: golang
description: Golang coding guidelines. MUST use BEFORE activity related to Golang code - analysis, review, planning, implementation.
---

<golang_guidelines>
1. **Ownership rule:**
   1) If a function's behavior semantically belongs to one locally defined type and uses that type's state or invariants, it MUST be a method on that type.
   2) MUST NOT create wrapper, projector, builder, or namespace types only to convert free functions into methods.
   3) Cross-package mapping, constructors, and logic that combines multiple types MAY remain free functions when no existing type owns behavior.
   4) MUST NOT move behavior to a type when that move would violate package ownership or dependency direction.
2. **Cross-package interface rule:**
   1) Go interface MUST belong to its consumer package.
   2) Every implementation MUST contain a compile-time interface assertion in implementation package.
   3) Before adding consumer import or assertion, MUST verify that consumer package does not depend transitively on implementation package.
   4) If required assertion creates an import cycle, dependency direction is invalid. MUST STOP and report complete cycle.
   5) MUST NOT move or remove assertion, move interface away from its consumer, suppress check, or add an adapter, alias, or shared contract package to hide cycle.
   6) MUST remove dependency edge that violates intended architecture before implementation continues.
</golang_guidelines>