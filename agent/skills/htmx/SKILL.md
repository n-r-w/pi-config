---
name: htmx
description: Use when designing, implementing, reviewing, or testing htmx user interfaces served by Go HTTP handlers and html/template.
---

<htmx_go_guidelines>
  <scope>
    Use these rules for server-rendered Go applications that use htmx 2.x. Follow project rules and `golang` skill first when they are stricter.
  </scope>

  <core_model>
    1. Treat HTML as application protocol. Go handlers SHOULD return rendered HTML, not JSON that browser code converts to HTML.
    2. Keep authoritative application state and business rules on server. Use htmx attributes for requests, targets, swaps, and events.
    3. Prefer standard links and forms with htmx enhancement. Application SHOULD keep core navigation and form submission usable when JavaScript is unavailable.
    4. Use semantic HTTP methods:
      1) `GET` reads state and MUST NOT change server state.
      2) `POST`, `PUT`, `PATCH`, and `DELETE` change state.
    5. Keep htmx behavior near HTML element that owns it. Move behavior to JavaScript only when htmx attributes and server responses cannot express it clearly.
    6. When swapped content contains third-party JavaScript components, initialize them with `htmx.onLoad` and release resources on `htmx:beforeCleanupElement`.
  </core_model>

  <library_delivery>
    1. Pin an exact htmx 2.x version.
    2. Production applications SHOULD serve pinned file from their own static assets. If a CDN is required, use HTTPS, an exact version, Subresource Integrity, and `crossorigin="anonymous"`.
    3. Load htmx once. Load each extension after htmx and enable it with `hx-ext` only in smallest required DOM subtree.
    4. Do not add npm or a JavaScript build step only to load htmx.
  </library_delivery>

  <templates>
    1. Use `html/template` for all HTML responses. MUST NOT use `text/template` for HTML.
    2. Parse templates at startup. A parse error MUST stop startup.
    3. Use named templates for pages, layouts, and fragments. Render a full page for normal navigation and matching fragment for an htmx request.
    4. Render fragment markup through one component template. A full-page template SHOULD include that component and MAY add layout data. Do not copy component markup into separate implementations.
    5. Treat template source as trusted and template data as untrusted. MUST NOT convert untrusted data to `template.HTML`, `template.JS`, `template.CSS`, `template.URL`, or related trusted types.
    6. Render to a buffer before writing response headers. If template execution fails, return an error response instead of a partial HTML response.
    7. Set `Content-Type: text/html; charset=utf-8` for HTML responses.
  </templates>

  <request_handling>
    1. Detect htmx requests with an exact case-insensitive check of `HX-Request: true`. Treat all `HX-*` headers as untrusted input.
    2. `HX-Request` MAY select a response representation. It MUST NOT bypass authentication, authorization, CSRF checks, validation, rate limits, or audit rules.
    3. If one URL returns a full page without `HX-Request` and a fragment with `HX-Request`, add `Vary: HX-Request` before writing status.
    4. Parse and validate form, path, and query values on server. Browser validation is a usability feature, not a security boundary.
    5. Keep handlers thin. A handler SHOULD parse input, call application logic, select a template, and map result to an HTTP response.
    6. Pass `r.Context()` to database and upstream calls so canceled htmx requests stop server work.
  </request_handling>

  <response_contracts>
    1. Define each interaction as a contract between requesting element and Go handler:
      1) request method and URL;
      2) submitted values;
      3) target selector;
      4) swap mode;
      5) success fragment;
      6) validation and failure behavior.
    2. Return smallest complete fragment that owns changed UI state. Fragment MUST contain every element required by its target and swap mode.
    3. Keep element `id` values unique. Keep an element ID stable across replacements when focus, CSS transitions, out-of-band swaps, or client code depends on it.
    4. Use `hx-swap-oob` only when one request must update independent page regions. Prefer one normal target when it represents interaction.
    5. Use `204 No Content` when client must not swap response content.
    6. htmx does not swap `4xx` and `5xx` responses by default. For validation errors, either return a rendered form fragment with `200`, or define and test a project-wide `responseHandling` rule that swaps chosen error status.
    7. htmx does not process its response headers from a `3xx` response. For an htmx redirect, return a non-3xx response with:
      1) `HX-Location` for htmx navigation and history update;
      2) `HX-Redirect` for a full browser reload, including destinations that require different scripts, styles, or other `<head>` content.
    8. For a non-htmx form submission, use normal Post/Redirect/Get flow with `303 See Other` after success.
    9. Use `HX-Trigger`, `HX-Trigger-After-Swap`, or `HX-Trigger-After-Settle` for small cross-component notifications. Event names and JSON payload fields MUST be stable response contracts.
    10. Validate every dynamic URL used in `HX-Location`, `HX-Redirect`, `HX-Push-Url`, or `HX-Replace-Url` to prevent open redirects.
  </response_contracts>

  <history_and_caching>
    1. Every URL added by `hx-push-url` MUST return a complete page when opened directly.
    2. When `HX-Request` selects fragments, set `htmx.config.historyRestoreAsHxRequest` to `false`. A history cache miss must receive a complete page.
    3. If a response varies by another request header, add that header to `Vary` too.
    4. If responses use `ETag`, generate different values for full-page and fragment representations.
    5. Use `hx-history="false"` on pages whose HTML MUST NOT enter htmx `localStorage` history cache.
  </history_and_caching>

  <forms_and_concurrency>
    1. Use native form controls, `name` attributes, labels, `action`, and `method`. Add htmx attributes as enhancement.
    2. Use `hx-encoding="multipart/form-data"` for file uploads. Enforce body size, file count, media type, and storage limits in Go.
    3. Show request state with `hx-indicator`. Use `hx-disabled-elt` for controls that MUST not be submitted twice.
    4. Use a delayed change trigger such as `hx-trigger="keyup changed delay:300ms"` for live search. Choose delay from measured request cost and expected typing rate.
    5. For live search, use `hx-sync="this:replace"` when newest result must win. For field validation, use `hx-sync="closest form:abort"` when form submission must win. Server mutations MUST remain safe under duplicate or reordered requests.
  </forms_and_concurrency>

  <security>
    1. Use relative same-origin URLs for htmx requests. Keep `htmx.config.selfRequestsOnly` enabled.
    2. Apply CSRF protection to state-changing requests when browsers attach credentials automatically. Prefer a hidden form field. If application uses `hx-headers`, render token from trusted server state and validate it in Go.
    3. Authentication cookies MUST use `Secure`, `HttpOnly`, and an explicit `SameSite` policy in production. Cookie flags do not replace CSRF checks required by application threat model.
    4. Escape untrusted data with `html/template`. If product accepts user HTML, sanitize it with an allowlist and wrap it in `hx-disable` so injected `hx-*` and `data-hx-*` attributes cannot run.
    5. Do not place untrusted data in tag names, attribute names, inline script, inline style, event handlers, or htmx expressions. Validate untrusted link and media URLs against allowed schemes and hosts.
    6. Prefer a Content Security Policy that limits scripts and connections to approved origins. If `htmx.config.allowEval` is `false`, do not use trigger filters, `hx-on`, or `js:` values in `hx-vals` and `hx-headers`.
    7. Set `htmx.config.allowScriptTags` to `false` unless swapped responses require trusted script execution.
    8. Authenticate and authorize target resource in every handler. Do not infer permission from `HX-Current-URL`, `HX-Target`, `HX-Trigger`, or `HX-Trigger-Name`.
  </security>

  <accessibility_and_ux>
    1. Use semantic HTML before adding ARIA attributes.
    2. Every form control MUST have an associated label. Keyboard operation and visible focus MUST work after each swap.
    3. Enable `htmx.config.reportValidityOfForms` so invalid forms report errors and focus first invalid control.
    4. Put validation messages next to their controls and provide an accessible summary when form has multiple errors.
    5. Use an `aria-live` region for status messages that change outside current focus.
    6. After `outerHTML` swaps, return required focus target or move focus through a tested htmx event handler.
  </accessibility_and_ux>

  <examples>
    Example directories are ACCESSIBLE and located in subdirectories RELATIVE to this file (no need to go fetch them from web!):
      1. [Full-page and fragment handler](examples/internal/basic)
      2. [Form validation, CSRF, and redirects](examples/internal/form)
  </examples>

  <testing>
    1. Test business rules without HTTP or templates. Assert returned values, domain errors, and state changes.
    2. Use `httptest` for handler contracts:
      1) request without `HX-Request` returns a full document;
      2) request with `HX-Request: true` returns expected fragment;
      3) `Content-Type`, `Vary`, status, and htmx response headers are correct;
      4) invalid input and unauthorized access cannot mutate state.
    3. Parse HTML responses into a DOM. Assert stable semantics: document or fragment, target root, form `action` and `method`, htmx attributes, ARIA relations, and escaped data.
    4. MUST NOT compare full HTML, whitespace, attribute order, class order, or decorative copy. Assert exact text only when wording is a user-visible contract.
    5. Test normal form submission and htmx submission for same business result. Test redirect behavior separately for both request types.
    6. Keep browser suite small. Test swaps, history, focus, indicators, double submission, request races, and JavaScript-disabled fallback.
    7. Use screenshot tests only for intentional visual regression. MUST NOT use screenshots to prove business behavior.
    8. Run Go race tests when handlers or application services use shared mutable state.
  </testing>

  <completion_checks>
    1. Direct navigation returns a complete page.
    2. htmx requests return valid fragments for their declared targets and swap modes.
    3. Full-page and fragment responses use one component implementation.
    4. Caching cannot mix full pages and fragments.
    5. Authentication, authorization, CSRF, validation, and escaping apply to htmx and normal requests.
    6. Error, redirect, history, focus, and concurrent-request behavior have tests.
  </completion_checks>

  <references>
    - https://htmx.org/docs/
    - https://htmx.org/reference/
    - https://htmx.org/essays/web-security-basics-with-htmx/
    - https://pkg.go.dev/html/template
    - https://pkg.go.dev/net/http/httptest
  </references>
</htmx_go_guidelines>
