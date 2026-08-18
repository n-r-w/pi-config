<test_scenario_template>
<!--
  <guidelines>
  1. Based on EARS (Easy Approach to Requirements Syntax):
  `WHILE {optional pre-condition}, WHEN {optional trigger}, THE {actor} SHALL {required behavior}`
  2. `Scenario` section MUST contain business logic, not specific test data.
  </guidelines>
-->

# Test scenarios for {project/task/etc.}

## {ID}: {Business rule name}

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