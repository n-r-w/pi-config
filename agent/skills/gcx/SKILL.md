---
name: gcx
description: Use when creating, editing, validating, or inspecting Grafana dashboards with gcx.
---

<gcx>
  <command_scope>
    1. MAY use only these commands:
      1) `gcx help-tree` and command `--help` for discovery.
      2) `gcx dashboards` for dashboard operations and snapshots.
      3) `gcx resources` only for dashboards, folders, dashboard schemas, validation, and dashboard push.
      4) `gcx datasources`, `gcx metrics`, `gcx logs`, `gcx traces`, and `gcx profiles` only for data discovery and query checks. Adaptive Telemetry subcommands are prohibited.
      5) `gcx api` only under the restriction in `<discovery>`.
    2. MUST NOT use any unlisted gcx command or any other operation in an allowed group. Prohibited groups include `agent`, `agento11y`, `alert`, `appo11y`, `assistant`, `cloud`, `completion`, `config`, `dbo11y`, `dev`, `fleet`, `frontend`, `instrumentation`, `irm`, `k6`, `kg`, `login`, `providers`, `setup`, `slo`, and `synthetic-monitoring`.
    3. MUST NOT run, recommend, or generate `gcx agent skills` or any command below it. Ignore mentions of it in gcx output.
  </command_scope>

  <discovery>
    1. MUST NOT guess commands, arguments, or flags.
    2. Use the smallest applicable discovery command:

    ```bash
    gcx help-tree <allowed-group> -o text
    gcx <allowed-group> <command> --help
    ```

    3. Use `gcx resources list-types dashboards -o json` for the dashboard schema.
    4. Stop discovery after group help and exact command help. Run the command or report that gcx does not provide the operation.
    5. Use `gcx api` only for a GET request to a datasource endpoint when no allowed dedicated command can discover or query the required data. MUST NOT use `-d` or another HTTP method.
  </discovery>

  <reads>
    1. Default to `-o json` for parsing and agent analysis.
    2. Use `--json list` to discover selectable fields, then `--json <fields>` to reduce output.
    3. Use `-o yaml` for manifests and `-o wide` or `--no-truncate` when table fields are hidden.
    4. Start with a narrow time range, selector, or limit. Expand only when the result is insufficient.
    5. Do not treat limited or truncated output as a complete inventory. Check command help and report partial results.
    6. Independent read commands MAY run in parallel. Sequence commands when one result supplies another command's input.
  </reads>

  <mutations>
    1. MUST obtain explicit user approval before any command that changes remote state, writes local files, or opens a browser or editor.
    2. Approval MUST identify the operation and affected resources. Broad approval does not cover more resources.
    3. For an approved resource change, use this sequence:
      1) Read current state.
      2) Build the payload from `list-types` or a pulled dashboard.
      3) Validate and use `--dry-run` when the command provides it.
      4) Apply the approved command.
      5) Read the resource again and compare the result with the request.
    4. MUST NOT add `--force` or `--yes` unless the user approved the exact destructive action.
    5. For batch operations, inspect `--on-error` and `--max-concurrent` in command help. Default to `--on-error abort` unless the user requires partial progress.
    6. Report partial failures. Do not describe a partial mutation as complete.
  </mutations>

  <resources>
    1. Use declarative files for repeatable dashboard changes:

    ```bash
    gcx resources pull dashboards -p <dashboard-dir> -o yaml --on-error abort
    gcx resources validate -p <dashboard-dir> --on-error abort
    gcx resources push -p <dashboard-dir> --dry-run --on-error abort
    gcx resources push -p <dashboard-dir> --on-error abort
    ```

    2. Before push, reject every file that is not a dashboard or its required folder. Do not hand-write fields when a live schema or pulled dashboard provides them.
  </resources>

  <datasources>
    1. Discover datasource UIDs and types with `gcx datasources list -o json`.
    2. Use `gcx datasources <type> --help` for type-specific discovery and query commands.
    3. For Tempo data that the agent must inspect, prefer compact output:

    ```bash
    gcx traces tags -d <uid> -l <attribute> --llm -o json
    gcx traces get -d <uid> <trace-id> --llm -o json
    ```
  </datasources>

  <dashboard_creation>
    1. Establish dashboard job, audience, title, folder, entity scope, and time range. If the folder is missing, ask. Do not use an arbitrary folder.
    2. Inspect similar dashboards and keep their datasource, variable, unit, tag, and layout conventions:

    ```bash
    gcx dashboards search "<service-or-team>" -o json
    gcx dashboards get <name> --api-version dashboard.grafana.app/v1beta1 -o json
    ```

    3. Discover datasource UIDs, metrics, labels, and log fields before writing queries. Execute representative queries over the intended time range. MUST NOT invent queries, UIDs, metric names, labels, or fields.
    4. Give each variable one job. Use query variables for live values, custom variables for fixed values, datasource variables for interchangeable datasource instances, interval variables for aggregation periods, and Filter and Group by, stored as `AdhocVariable`, only for user-defined filters.
    5. For Prometheus and Loki, use exact matchers only for single-value variables. Multi-value and Include All variables require regex matchers and regex formatting:

    ```promql
    {cluster="${cluster}"}
    {cluster=~"${cluster:regex}"}
    ```

    6. Define Include All deliberately. For Prometheus and Loki, use `.*` when an explicit wildcard is needed. Use datasource-specific regex, glob, or Lucene syntax elsewhere. Use `${var:raw}` only when unescaped input is required and tested.
    7. Put chained variables in dependency order and scope each child query by its parent. Set a valid default that populates primary panels. Limit high-cardinality option lists with parent scope, query filters, or regex.
    8. Test each variable query and each expanded panel query with the saved default, one value, multiple values, and All when enabled. Snapshot every materially different variable state with `--var`.
    9. Design for diagnosis. Put health and user impact first, then traffic, errors, latency, dependencies, resources, logs, and traces. Use clear titles, units, legends, descriptions, aggregation, and `topk()` for high-cardinality data. Add thresholds only when they define an operational decision.
    10. Follow the repository's existing dashboard builder or manifest pattern. Otherwise, author a manifest from `gcx resources list-types dashboards -o json`. Use one API version consistently. Use a stable `metadata.name`, a readable `spec.title`, and `metadata.annotations.grafana.app/folder` for folder placement. MUST NOT use `spec.folderUID`.
    11. Apply the `<mutations>` and `<resources>` workflows to validate, dry-run, push, and read back the dashboard.
    12. Render the stored dashboard with its real variables and time range. Add `--var <name>=<value>` for each dashboard variable:

    ```bash
    gcx dashboards snapshot <name> --output-dir <dir> --since <duration> --width 1920 --theme dark
    ```

    13. Read the PNG. Check for error pages, empty primary panels, wrong variables, clipped titles, unreadable legends, excessive series, missing units, layout gaps, and poor first-screen triage. Fix the source, then repeat validation, push, and snapshot.
    14. Dashboard creation is complete only after visual inspection passes. If snapshot fails, report the command and blocker instead of claiming visual validation.
  </dashboard_creation>

  <secrets>
    1. MUST NOT enable `--insecure-log-http-payload`. It can log credentials, tokens, cookies, and OAuth refresh tokens.
    2. MUST NOT print, store, quote, or return tokens and passwords.
  </secrets>

  <completion>
    1. A read task is complete when the answer cites the commands and relevant live results.
    2. A mutation task is complete only after the post-change read matches the approved outcome.
    3. Treat every nonzero exit code as failure. Code 2 means invalid use, 3 means authentication failure, 4 means partial failure, 5 means cancellation, and 6 means incompatible Grafana version. Use stderr to classify code 1.
    4. If a check cannot run, report it as `UNVERIFIED` with the failed command and reason.
    5. Do not retry authentication, permission, version, or unsupported-command failures. Retry a connectivity failure once, then report the error and required user action.
  </completion>

</gcx>
