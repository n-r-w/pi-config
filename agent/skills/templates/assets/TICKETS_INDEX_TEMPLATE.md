<tickets_index_template> <!-- MUST NOT follow <id_guidelines> for this template -->
# Tickets Index: {name}

## Tickets List
<!--
One ticket MUST contains EXACTLY ONE vertical slice:
1. Each slice cuts a narrow but COMPLETE path through every layer (schema, API, UI, tests): vertical, NOT a horizontal slice of one layer
2. A completed slice is demoable or verifiable on its own
3. Each slice is sized to fit in a single fresh context window
4. Any prefactoring should be done first
-->

### {Unique ID}. {Name of ticket}
- Owner: {Service in charge}
- Result: {Expected outcome of ticket}
- Dependencies: {List of dependencies, if any}
- Blockers: {List of blockers, if any}
- File: {Path to ticket file}

## Execution Order
{List of tickets in the order they should be executed}

## References
- {reference (file/URL/standard, etc.)} - {one-line description}
- ...
</tickets_index_template>