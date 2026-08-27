---
name: templates
description: 'Documentation templates: Ticket, Change Request, Algorithm, Specification, Architecture, Plan, Review, Research/Analysis/Investigation/Brainstorming, Information Extraction, and other. MUST read and follow BEFORE writing any type of documentation. Use this skill for artifacts, NOT FOR COMMUNICATION'
---

<templates>
   <guidelines>
      1. All templates are ACCESSIBLE and located in `assets` subdirectory RELATIVE to this file (no need to go fetch them from web!)
      2. CRITICAL: Content Purity Rule.
         In any generated document, statements MUST describe document’s subject matter as defined by selected template (e.g., product scope, architecture, plan, review findings, research conclusions, extracted facts).
         Statements MUST NOT include instructions about how to write, structure, or limit documents.
      3. Meta Guidance Placement Rule.
         Process/meta guidance is allowed only in sections explicitly designated by template for meta content (if any). It MUST NOT appear in domain/content sections.
      4. Strictly maintain consistency of document in identifiers.
      5. Follow chosen template RIGOROUSLY when writing your document.
      6. When creating documents that mention several different services, you MUST add qualifiers to any attributes (handlers, properties, etc.) so that it is clear which service attribute belongs to. E.g. `service1.handler`.
      7. Maintain consistency in language of document. E.g. if document is written in English, then all its headings must also be in English.
   </guidelines>
   <constraints>
      1. When creating documents, avoid adding "historical" information such as "what was corrected" or other information that is not relevant to document's content unless it is required (for example, for a second review, it makes sense to indicate what was corrected). Otherwise, ONLY ACTUAL information should be included.
      2. MUST NOT read a template if it is not DIRECTLY relevant to current task. For example, if your task is to write code and then run a sub-agent with independent review, then review template does NOT need to be read, as that is a task for sub-agent, not for you!
      3. When generating README, FAQ, how-to guides, and other similar documentation, MUST NOT use templates from this list and MUST NOT use `ID prefix guideline`. Keep it simple and intuitive.
      4. These templates are ONLY for creating documents. NOT FOR user-facing messages!
   </constraints>
   <document_creation_rules>
      Immediately before starting the creation of the document (not earlier), you MUST read `rules/doc_rules.md` located in subdirectory RELATIVE to this file.
   </document_creation_rules>
   <template_selection>
      <mapping note="Choose and read a template from list below. Choose by goal">
         1. Documents that require following `ID prefix guideline` (`assets/IDS.md`):
            1) Request some changes -> `<change_request_template>`: `assets/CHANGE_REQUEST_TEMPLATE.md`
               For changes in documents, architecture, plans, specifications, and other similar artifacts.
            2) Define what/why -> `<specification_report_template>`: `assets/SPECIFICATION_TEMPLATE.md`
            3) High level system design -> `<architecture_report_template>`: `assets/ARCHITECTURE_TEMPLATE.md`
            4) Plan phases/tasks -> `<plan_report_template>`: `assets/PLAN_TEMPLATE.md`
            5) Detail a development task for one plan phase/stage -> `<task_template>`: `assets/TASK_TEMPLATE.md`
            6) Investigate/compare -> `<analysis_report_template>`: `assets/ANALYSIS_RESULTS_TEMPLATE.md`
            7) Review existing artifact -> `<review_report_template>`: `assets/REVIEW_RESULTS_TEMPLATE.md`
            8) Extract facts only -> `<extraction_report_template>`: `assets/EXTRACTION_TEMPLATE.md`
            9) Describe step-by-step logic -> `<algorithm_report_template>`: `assets/ALGORITHM_TEMPLATE.md`
            10) Identification of open questions and blockers -> `<discrepancy_report_template>`: `assets/DISCREPANCY_TEMPLATE.md`.
            11) Standalone technical solution description -> `<technical_solution_template>`: `assets/TECHNICAL_SOLUTION.md`
            12) Implementation supplement. Can be used as addition to a main technical solution document -> no specific template, but MUST follow `ID prefix guideline`
            13) Ticket (requirements for task implementation) -> `<ticket_template>`: `assets/TICKET_TEMPLATE.md`
            14) Test scenarios -> `<test_scenario_template>`: `assets/TEST_SCENARIO_TEMPLATE.md`
               For describing test scenarios for a specific task, feature, or project.
         2. Documents for which `ID prefix guideline` is PROHIBITED:
            1) ADR (Architecture Decision Record) -> `<adr_template>`: `assets/ADR_TEMPLATE.md`
            2) Problem Statement -> `<problem_statement_template>`: `assets/PROBLEM_STATEMENT.md`
               Used for preliminary description of a problem that requires further research and analysis. It is not a PRD or technical design.
            3) README, FAQ, how-to guides, etc. -> NO template.
            4) Brainstorming results, ideas, other documents that are not final and not intended for publication -> NO template.
            5) Tickets index (list of tickets with their dependencies and execution order) -> `<tickets_index_template>`: `assets/TICKETS_INDEX_TEMPLATE.md`
            6) Domain terms and definitions -> Follow simple structure: `term: definition`. MUST NOT add business logic description, etc., ONLY list of terms and their definitions. No template.
               If project contains artifacts in several languages, then term name must be specified in ALL languages of project, and definition of term must be in language in which document is written: `- term eng / term rus: definition rus`.
            7) Light Project Requirements Document (for quick idea evaluation) -> `<light_prd_report_template>`: `assets/LIGHT_PRD.md`. For quick idea evaluation, not a full specification. Use by default to capture requirements unless user explicitly specifies otherwise. Don't add "light" attribute to a file or document - it's a property, not a name.
            8) Analysis and feedback on the reviewer's comments on PR/MR -> `<review_feedback_template>`: `assets/REVIEW_FEEDBACK_TEMPLATE.md`.
         3. If document does not fit any template and user does not ask for a specific template, follow `<format_evaluation_workflow>` below.
      </mapping>
      <template_choice>
         1. Architecture vs Technical Solution:
            1) Architecture document describes high-level system design, its components, interfaces, and contracts. It is a blueprint for system.
            2) Technical Solution document describes a specific solution to a problem, including implementation details.
            3) Technical Solution can appear as a detailed part of architecture or be a standalone document, but you MUST NEVER turn architecture into a detailed technical solution.
         2. Specification vs Plan:
            1) Specification document describes what needs to be done and why, but does not describe how to do it.
            2) Plan document describes how to do something, including phases, tasks, and their execution order.
         3. Ticket vs Task:
            1) Ticket describes requirements for implementing a specific task, but does not describe in detail how to implement task.
            2) Task describes in detail how to implement a specific task, including code, configuration, scripts, tests, etc.
         4. Problem Statement vs Specification:
            1) Problem Statement document describes a problem, not to how to solve it.
            2) In opposite, Specification document describes requirements for a solution to a problem.
         5. ADR vs Architecture:
            1) ADR document describes a decision made regarding architecture. It is base for architecture, but not architecture itself.
            2) Architecture document describes overall system design.
      </template_choice>
      <format_preservation>
         If an existing document fits a template type but already has a certain format, it MUST BE PRESERVED.
         For example, if document does not have IDs, they MUST NOT be added!
      </format_preservation>
      <open_questions_handling>
         Any research document (brainstorming, analysis, ideas, etc.), even if it does not have a template, MUST have an "Open Questions" section.
      </open_questions_handling>
      <format_evaluation_workflow>
         1. Evaluate how well target document fits one of templates.
         2. If there is even a slight doubt, IMMEDIATELY STOP AND ASK USER. GUESSING DOCUMENT FORMAT IS PROHIBITED!
         3. If document fits one of templates, then read template and follow it RIGOROUSLY.
            Otherwise, MUST NOT READ ANY TEMPLATE AND `ID prefix guideline`!
      </format_evaluation_workflow>
   </template_selection>
<templates>