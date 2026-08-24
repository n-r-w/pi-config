---
name: commit-message
description: How to generate accurate and descriptive commit messages.
---

<commit_message_guidelines>
  1. Use conventional commit message format
  2. Your commit messages should be concise, informative, and follow best practices for writing effective commit messages.
  3. Choose message language according project rules. If not specified:
    1) Analyze changes in codebase and determine the most appropriate language from code comments.
    2) If you cannot determine the language, default to ASD-STE100 (Simplified Technical English).
  4. Determine scope of changes to analyze:
    1) If user provides a description of changes, use that as the basis for your commit message.
    2) Analyze diff between current branch and changes.
  5. Present the generated commit message to the user in format:
  {type}({scope}): {subject}
  {BLANK LINE}
  {body}
</commit_message_guidelines>