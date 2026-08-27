---
name: mcp-yandex
description: Guidelines for using Yandex Wiki and Yandex Tracker via MCP. CAN be read and followed ONLY when Yandex MCP tools are available in current session. MUST NOT be read or applied otherwise.
metadata:
  uri: https://github.com/n-r-w/yandex-mcp
---

<yandex_mcp_guidelines>
  <critical_notice>
    🚨 IMPORTANT: DON'T USE regular web fetch tools for yandex wiki and tracker. This will result in an authentication error.
  </critical_notice>
  <document_analysis>
    1. After receiving any document or task:
      1) Analyze all links in text.
      2) Select links that are necessary for solving current task.
      3) For selected links:
        * If it is a link to Yandex Wiki or Tracker, use MCP for access.
        * If it is an external link, use available tools for downloading.
    2. Use redirect options if documents are not accessible directly.
  </document_analysis>
  <document_referencing>
    🚨 NEVER shorten links to Yandex Wiki and Yandex Tracker page/slug in documentation, even if they are long. Otherwise, link will be invalid for user.
  </document_referencing>
  <tracker_tasks>
    <statuses>
      1. Available statuses (`russian key` -> `english key`):
        1) `Бэклог` -> `backlog`
        2) `Открыт` -> `open`
        3) `Уточнение требований` -> `utocnenieTrebovanij`
        4) `Согласование` -> `soglasovanie`
        5) `PBR` -> `pbr`
        6) `В работе` -> `inProgress`
        7) `Готово к ревью` -> `gotovoKRevu`
        8) `В ревью` -> `inReview`
        9) `Можно тестировать` -> `readyForTest`
        10) `Тестируется` -> `testing`
        11) `Верификация` -> `verifikacia`
        12) `Готово к деплою` -> `gotovoKDeplou`
        13) `Готово к производству` -> `gotovoKProizvodstvu`
        14) `Закрыт` -> `closed`
        15) `Отменено` -> `cancelled`
        16) `Остановлено` -> `ostanovleno`
        17) `Воронка идей` -> `voronkaIdej`
        18) `Анализ эффекта` -> `analizEffekta`
      2. API supports both russian and english keys for statuses.
    </statuses>
    <users>
      1. Available fields for user-based search:
        1) `assignee`: task assignee
        2) `created_by`: task author
        3) `updated_by`: last user who updated task
        4) `followers`: task followers
      2. Special values:
        1) `me()`: current authorized user
        2) `empty()`: absence of a user (e.g., unfilled assignee field)
    </users>
    <types>
      1. Available types:
        1) `bug`
        2) `epic`
        3) `story`
        4) `techdebt`
        5) `development`
        6) `testing`
        7) `analytics`
        8) `enabler`
      2. API supports ONLY english keys for types.
    </types>
    <examples>
      <example condition="Search tasks for user ivanov@email.com with status open">
        "filter": {"assignee": "ivanov@email.com", "status": "open"}
      </example>
      <example condition="Search for my tasks with status ready for deploy">
        "filter": {"assignee": "me()", "status": "gotovoKDeplou"}
        "filter": {"assignee": "me()", "status": "Готово к деплою"}
        "query": 'Assignee: me() AND Status: gotovoKDeplou'
        "query": 'Assignee: me() AND Status: "Готово к деплою"'
      </example>
      <example condition="Search for all tasks with status backlog in queue CP">
        "query": 'Queue: CP AND Status: backlog'
        "filter": {"queue": "CP", "status": "backlog"}
      </example>
      <example condition="Search for all tasks updated in last week with status ready for review">
        "query": 'Status: gotovoKRevu AND Updated: >today()-7d'
        "query": 'Status: "Готово к ревью" AND Updated: >today()-1w'
      </example>
      <example condition="Search for tasks in multiple statuses">
        "filter": {"status": "open,inProgress,backlog"}
        "query": '(Status: open OR Status: inProgress OR Status: backlog)'
      </example>
      <example condition="Search for critical priority bugs">
        "filter": {"priority": "critical", "type": "bug"}
        "query": 'Priority: Critical AND Type: Bug'
      </example>
      <example condition="Search in multiple queues">
        "query": '(Queue: CP OR Queue: BB) AND Status: open'
        "query": 'Queue: CP OR Queue: BB'
      </example>
      <example condition="Search for tasks without assignee">
        "query": 'Assignee: empty() AND Status: open'
        "filter": {"assignee": "empty()"}
      </example>
      <example condition="Search for tasks with resolution">
        "query": 'Resolution: empty()'
        "query": 'Resolution: !empty() AND Status: closed'
      </example>
      <example condition="Search excluding closed tasks">
        "query": 'Assignee: me() AND NOT Status: closed'
        "query": 'Queue: CP AND NOT (Status: closed OR Status: cancelled)'
      </example>
      <example condition="Search where I am author but not assignee">
        "query": 'Author: me() AND NOT Assignee: me()'
      </example>
      <example condition="Search for my tasks updated in last 30 days">
        "query": 'Assignee: me() AND Updated: >today()-30d'
      </example>
      <example condition="Search for tasks created in specific date range">
        "query": 'Created: >=2025-12-01 AND Created: <=2025-12-31'
        "query": 'Assignee: me() AND Created: >=2025-01-01'
      </example>
      <example condition="Search for cancelled or stopped tasks in last quarter">
        "query": '(Status: cancelled OR Status: ostanovleno) AND Updated: >today()-90d'
      </example>
      <example condition="Search for stories and epics in progress">
        "query": '(Type: Story OR Type: Epic) AND Status: inProgress'
        "query": 'Type: Story AND Status: inProgress'
      </example>
      <example condition="Search for technical debt tasks">
        "query": 'Type: techdebt AND Created: >=2025-12-01'
        "query": 'Type: techdebt AND Queue: CP'
      </example>
      <example condition="Search for critical tasks in production-ready statuses">
        "query": 'Priority: Critical AND (Status: gotovoKDeplou OR Status: gotovoKProizvodstvu)'
        "query": 'Priority: Critical AND Queue: CP'
      </example>
      <example condition="Search for tasks in testing statuses updated recently">
        "query": '(Status: testing OR Status: readyForTest) AND Updated: >=today()-3d'
      </example>
      <example condition="Search for open tasks without assignee in specific queue">
        "query": 'Assignee: empty() AND NOT Status: closed AND Queue: CP'
      </example>
      <example condition="Search for tasks by specific author in review statuses">
        "query": 'Author: a.iskritsky@bronevik.com AND (Status: gotovoKRevu OR Status: inReview)'
      </example>
      <example condition="Search for epic type tasks that are open">
        "query": 'Type: Epic AND Status: open'
      </example>
    </examples>
  </tracker_tasks>
</yandex_mcp_guidelines>