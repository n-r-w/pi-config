<algorithm_report_template>
<!--
<algorithm_guidelines>
  1. No document size limit, but aim for conciseness, clarity and no duplication. Avoid unnecessary details that do not contribute to understanding the algorithm.
  2. Focus on the key components and stages of the algorithm.
  3. Include diagrams or flowcharts only when they materially improve understanding.
  4. Code examples must stay narrowly focused on the algorithm. Do not include unrelated concerns such as logging, error handling, or helper utilities unless they are essential to understanding it.
  5. Use code examples only when they help explain the algorithm. Avoid showing data structures or other code that is not central to the explanation.
  6. Add comments to code examples when they clarify the algorithm’s intent or steps. Prefer writing new explanatory comments over copying existing ones.
  7. Verify that all referenced files and functions are current.
  8. Path references must be relative to the current directory of the documentation file. Do not include line numbers in path references.
  9. Do not add sections beyond the required document structure.
  10. All items must follow <id_guidelines> and have unique IDs.
</algorithm_guidelines>
-->

# {Algorithm Name}

## Key definitions and abbreviations

- {id}: {Definition}. {Explanation of the term or abbreviation.}
- ...

## Overview

{Description of the algorithm and its purpose.}

## Key Steps
- {id}: Step Name. Brief description of this step's role.
...

## Main Components
{Here we describe the main components of the algorithm, their roles, and locations in the codebase.
A component can be: a part of a system (a module, an external service, logically related blocks of code, etc.). A component that is not: a function, a structure, etc.}

### {Component Name}

#### Role in the algorithm
{Brief description of this component's role in the overall algorithm.}

#### Location
{Relative path to file(s) and function(s), in format [path/to/file.go](path/to/file.go), functions: Function1, Function2.}

### ...
...

## Diagram
{Mermaid diagram for visually representing the algorithm's flow and components.}

## Operation Description

### {id}: {Step Name}

#### Business Context
{Business-level explanation of this step for non-technical readers (analysts, product managers). State what the step means in domain terms and why it works this way. MUST NOT contain code references, function names, or internal identifiers. MUST NOT duplicate content from "Role in the algorithm" or "Detailed Description". Be brief and concise.}

#### Role in the algorithm
{Brief description of this step's role.}

#### Location
{Relative path to file(s) and function(s), in format [path/to/file.go](path/to/file.go), functions: Function1, Function2.}

#### Detailed Description
{Detailed explanation of what happens at this step, including any important details or nuances.
MUST NOT include code snippets except when absolutely necessary for clarity. ALWAYS better to explain in words.
Use headings, bullet points, or numbered lists if it helps to structure the explanation better.}

### ...
</algorithm_report_template>