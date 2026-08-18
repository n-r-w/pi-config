---
name: perplexity
description: How to use `Ask Perplexity` tool to search the web and get answers to user questions.
---

<perplexity_guidelines name="How to use `Ask Perplexity` tool">
    1. Use two roles (system and user) for consistent quality. Example:
        1) role: "system", content: "You are a principal software engineer with deep expertise in designing scalable, maintainable, and robust software systems"
        2) role: "user", content: "What is the best practice for clean architecture in Golang?"
    2. Choose `system` role based on the topic
    3. Provide specific, detailed question in `user` role
    4. ALWAYS retry on `Request timed out` errors
    5. Perplexity KNOWS NOTHING about the internal context of the project and internal libraries, so ALWAYS provide relevant information from the code and documentation for queries related to the project.
    6. Include in the query strict requirement to provide links that can be accessed by automated agents, without captcha and authorization (e.g., raw.githubusercontent.com) to avoid access issues.
    7. DON'T use `country` parameter.
</perplexity_guidelines>