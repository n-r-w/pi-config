<test_scenario_template>
<!--
  <guidelines>
  1. Based on EARS (Easy Approach to Requirements Syntax):
  `WHILE {optional pre-condition}, WHEN {optional trigger}, THE {actor} SHALL {required behavior}`
  2. `Scenario` section MUST contain business logic, not specific test data.
  3. Use IDs only when required by <id_guidelines>. Do not infer this need from the item type or template section.
  4. Every assigned ID must be unique within the document.
  5. Assign a SCN ID or reuse the source requirement ID only when traceability or independent scenario execution requires a stable reference.
  </guidelines>
-->

# Test scenarios for {project/task/etc.}

## {Business rule name}
<!-- When a stable reference is required, use `## {SCN-id or source requirement ID}: {Business rule name}`. -->

{Brief business rule description}

### Scenario

**Actor:** {Entity responsible for performing the required behavior. E.g.: system, user}

**Pre-conditions:**
- {State or pre-condition that must already be true for the requirement to apply. E.g.: "WHILE the user is authenticated"}
- ...

**Triggers:**
- {Event, input, command, signal, or condition transition that causes the required behavior to be executed. E.g. "WHEN the user submits the payment form"}
- ...

**Shall:**
- {Indicates a mandatory requirement. E.g.: "API SHALL return HTTP 403""}
- ...

### How to test
- {Steps to verify the requirements }
- ...

### Examples
{Input and expected output parameters}

</test_scenario_template>