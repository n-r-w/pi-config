---
name: review
description: Review standards and best practices. Use EXACTLY BEFORE performing any kind of review (code review, review of implementation plan, etc.). DON'T need to read it BEFORE delegating review to subagents, as they will have access to it.
---

<review_standards>
  <scope name="Scope of review">
    <initial_review_scope>
      1. If not otherwise specified, code review MUST cover only changes introduced in current pull request, commit, staged or unstaged changes.
      2. DON'T review code that not modified unless it directly impacts changes being reviewed.
      3. DON'T review generated code: protobufs, ORM models, API clients, etc.
      4. Choose only `<review_areas>` relevant to specific task.
      5. Focus on potential problems, not positive changes.
    </initial_review_scope>
    <re_review_scope>
      1. If review is a re-review, scope MUST be set to `<initial_review_scope>`.
      2. It is FORBIDDEN to narrow scope of re-review to only fixes that were made after previous comments. Each re-review should be as thorough as first one and SHOULD NOT be limited to only fixed parts.
      3. It is FORBIDDEN to inform reviewer about what has already been fixed or checked, as this may negatively impact quality of re-review.
    </re_review_scope>
    <project_phase_scope>
      When conducting a review, you MUST limit review to current phase of project design based on `complex-project-design` skill. Otherwise, review may go into identifying problems that are NOT RELATED to current task and thereby break entire process.
    </project_phase_scope>
  </scope>

  <iron_law>
    **Maximum Depth:** You must engage in exhaustive, deep-level reasoning.
    **Multi-Dimensional Review:** Analyze request through every lens
    **Prohibition:** NEVER use surface-level logic. If reasoning feels easy, dig deeper until logic is irrefutable.
    **Edge Case Review:** What could go wrong and how we prevented it.
    **Backward Compatibility, Fallback & Deprecation:** allowed ONLY if they are EXPLICITLY stated in requirements. In all other cases, it is an ERROR.
    **No Overengineering:** Reject any solution that try to solve edge cases that are unlikely to occur or do not exist at all in real scenarios. Follow KISS and YAGNI principles!
    **Zero Tolerance Policy:** DON'T allow issues of level medium and above, as well as technical debt. Highly recommend fixes for low level issues.
    **Code Slop Hunting:** IMMEDIATELY REJECT any findings that appear to be `<code_slop>` from `review` skill.
    **No Deviation:** Any deviation from standards/skills that is not explicitly permitted by user MUST lead to rejection!
    **Subagents Utilization:** Utilize subagents usage and parallelize subagent tasks whenever possible to maximize speed and efficiency.
  </iron_law>

  <non_negotiables name="Non-negotiable rules">
    1. DON'T treat local git state (untracked/unstaged/uncommitted files) as a review issue or a blocker
    2. Verify that new code uses pre-existing project patterns in testing, error handling, logging, configuration management, etc. Prevent "reinventing wheel".
    3. Check compliance with all relevant skills, project rules and best practices. If skills are not loaded, load them first.
    4. REJECT "quick fixes" without explicit user permission. Such fixes lead to technical debt and create problems in future.
    5. PREFER refactoring over dirty hacks.
    6. It is FORBIDDEN to implement temporary solutions with expectation of future refactoring without explicit user permission.
    7. Reject suppression of linter errors unless there is a compelling reason to do so.
    8. ALWAYS perform reviews in current repository. No weird actions like code copies to a temporary directory for clean checkout review, etc.
  </non_negotiables>

  <diff_analysis>
    1. ALWAYS perform an overall analysis of list of changed files to avoid missing important changes.
    2. ALWAYS compare changed files with existing file structure to identify code placement issues (e.g., constants in common files instead of specific constants files, etc.).
    3. Analyze ONLY diffs that relevant to user's request.
    4. REMEMBER that analyzing unrelated diffs wastes context and may lead to failure to complete task.
  </diff_analysis>

  <review_areas name="Common Areas to consider during any review">
    <area name="General Principles">
      1. Consistency
      2. Completeness
      3. KISS & YAGNI
    </area>
    <area name="Scope & Requirements">
      1. Scope
      2. Requirements Coverage
      3. Assumptions vs Verified Context
      4. Backward Compatibility
      5. Fallback Behavior
    </area>
    <area name="Architecture & Design">
      1. Architecture
      2. Design Coherence
      3. Modularity
      4. Project Structure
    </area>
    <area name="Implementation Planning">
      1. Implementation Plan Phases
      2. Plan Sequencing
    </area>
    <area name="Security">
      1. Security
      2. Data Protection
    </area>
    <area name="Performance & Reliability">
      1. Performance
      2. Resource Efficiency
      3. Reliability
      4. Fault Tolerance
      5. Concurrency
      6. No external I/O inside DB transaction
    </area>
    <area name="Product Surface">
      1. UX
      2. API Surface
    </area>
    <area name="Operations">
      1. Configuration
      2. Deployment
      3. Monitoring & Logging
    </area>
    <area name="Testing & Regression">
      1. Testing Strategy
      2. Test Driven Development (TDD)
      3. Regression Risk
      4. Test Coverage
    </area>
    <area name="Code Health & Maintainability">
      1. Consistency with Project Conventions (Style, Patterns, Practices)
      2. Refactoring Residue (Stubs, Wrappers, Dead Code, Deprecations)
      3. Unrequested New Tech Debt:
        1) TODO/FIXME. Reject ONLY if they DON'T POINT to real issues in code. Otherwise, THINK how to fix issue ORIGIN and suggest it in review comments.
        2) Unnecessary Files
      4. Huge Functions, that SHOULD be split into smaller ones.
    </area>
    <area name="Code Comments">
      1. Explaining "how" instead of "why"
      2. Outdated comments that don't reflect current code
      3. Functions and variables without comments (even if they are self-explanatory)
      4. Lack of comments in complex code sections
      5. REMEMBER: bad commenting equals bad maintainability. DON'T allow it!
    </area>
    <area name="Documentation">
      1. Documentation quality
      2. Prefer symbol/function names and search terms over hard-coded line numbers.
    </area>
    <area name="Dependencies">
      1. Dependencies
      2. Third-Party Integrations
    </area>
  </review_areas>

  <additional_guidelines>
    1. Before starting review, you MUST read additional review guidelines if they are relevant to your task.
    2. All guidelines are ACCESSIBLE and located in `assets` subdirectory RELATIVE to this file (no need to go fetch them from web!):
      1. [Specification/PRD](./assets/SPECIFICATION_REVIEW.md)
      2. [Architecture](./assets/ARCHITECTURE_REVIEW.md)
      3. [Plan](./assets/PLAN_REVIEW.md)
      4. [Code/Scripts](./assets/CODE_REVIEW.md)
      5. [Research/Analysis/Investigation/Brainstorming](./assets/RESEARCH_REVIEW.md)
      6. [Algorithm](./assets/ALGORITHM_REVIEW.md)
      7. [Ticket](./assets/TICKET_REVIEW.md)
      8. [Change Request](./assets/CHANGE_REQUEST_REVIEW.md)
    3. NEVER postpone guideline reading to later stages. If it fits - READ IT RIGHT NOW!
    4. Example: If you need to review an implementation plan, you MUST read `PLAN_REVIEW.md` guideline first.
  </additional_guidelines>

  <realism_assessment>
    1. MUST NOT only identify problems but also evaluate problem scenarios in terms of their realism and likelihood of occurrence.
    2. MUST NOT design a "spaceship" to solve a problem that might occur in one in a million cases.
    3. MUST ALWAYS rank problems by their degree of realism.
    4. If scenario is related to API calls, MUST ensure it is achievable through API and not just artificially reproduced through unit tests.
  </realism_assessment>
</review_standards>