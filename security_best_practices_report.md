# Security Best-Practices Review

Scope: Lexvia branch `codex/lexvia-enterprise-gateway`, based on new-api `v1.0.0-rc.21`.

## Executive summary

The review found four high-impact deployment weaknesses in the inherited HTTP bootstrap/realtime boundary and one stored configuration injection risk. All five are remediated in this branch. Existing request-body limits, SSRF-protected video fetching, strict session cookies, token authentication, quota overflow saturation, and admin-only audit fields were retained.

## Findings

### 1. [High, remediated] Credentialed wildcard CORS

- Evidence: `middleware/cors.go:13-39`.
- Previous behavior: every origin and every request header were accepted while credentials were enabled.
- Risk: a browser could attach session credentials to cross-origin requests, greatly expanding the impact of an untrusted or compromised origin.
- Remediation: credentials now default to disabled; credentialed CORS requires an explicit `CORS_ALLOWED_ORIGINS` allowlist. Allowed headers are enumerated, configuration values reject CR/LF, and wildcard-plus-credentials fails closed at startup.
- Verification: `middleware/cors_test.go` covers the safe default, explicit allowlist, and invalid wildcard/credentials combination.

### 2. [High, remediated] Diagnostic listener exposed on all interfaces

- Evidence: `main.go:160-171`.
- Previous behavior: pprof listened on `0.0.0.0:8005` whenever enabled.
- Risk: runtime profiles and process diagnostics could be reachable outside the host or container network boundary.
- Remediation: pprof now defaults to `127.0.0.1:8005`, is overrideable through `PPROF_ADDR`, and uses bounded header/idle settings.

### 3. [Medium, remediated] Unbounded public HTTP headers and idle connections

- Evidence: `main.go:223-229`.
- Previous behavior: the main `http.Server` configured only address and handler.
- Risk: slow-header and oversized-header traffic could retain resources longer than necessary. Global read/write timeouts were intentionally avoided because they would break legitimate SSE and streaming responses.
- Remediation: configurable `ReadHeaderTimeout`, `IdleTimeout`, and `MaxHeaderBytes` were added with conservative defaults.

### 4. [Medium, remediated] Analytics configuration injected into HTML without validation

- Evidence: `main.go:262-335`.
- Previous behavior: environment-provided analytics IDs and the Umami script URL were concatenated directly into HTML and JavaScript.
- Risk: a compromised deployment configuration could inject markup or script into every served frontend page.
- Remediation: analytics IDs now use a strict bounded character allowlist, script URLs require an absolute HTTP(S) URL, and HTML attribute values are escaped. Invalid configuration is ignored and audited.
- Verification: `main_security_test.go` covers script-like IDs, newlines, unsafe schemes, and scheme-relative URLs.

### 5. [High, remediated] Realtime WebSocket accepted every browser origin

- Evidence: `controller/relay.go` (`websocketOriginAllowed`).
- Previous behavior: the WebSocket upgrader returned `true` for every `Origin`.
- Risk: a malicious site could initiate a cross-origin realtime connection from a victim's browser, increasing exposure to cross-site WebSocket hijacking when browser-held credentials or ambient authorization are present.
- Remediation: non-browser clients without `Origin` remain supported; browser connections default to same-origin and cross-origin consoles require an explicit `CORS_ALLOWED_ORIGINS` entry. A wildcard never weakens the WebSocket check.
- Verification: `controller/websocket_origin_test.go` covers same-origin, allowlisted, untrusted, wildcard, invalid, and non-browser requests.

## Additional hardening delivered

- Browser responses now include `X-Content-Type-Options`, `X-Frame-Options`, `Referrer-Policy`, and a restrictive `Permissions-Policy` (`middleware/security_headers.go:7-14`).
- Seedance native requests enforce a non-empty model/text prompt and the existing bounded billing duration before forwarding (`relay/channel/task/doubao/adaptor.go:120-194`).
- Retry routing excludes already-attempted channels, reducing repeated requests to the same failing upstream (`controller/relay.go:294-338`).
- Subscription purchase/management routes are not registered, and new requests use wallet-only funding (`router/api-router.go:150-163`, `service/billing_session.go:337-375`).

## Residual considerations

- A deployment-specific Content Security Policy is not forced by the application because operators may configure different analytics and embedded console origins. Configure CSP at the ingress after enumerating the required origins.
- TLS termination is expected at the reverse proxy/load balancer. Keep the application listener on a private network and enforce HTTPS/HSTS at the edge.
- Legacy subscription tables and settlement helpers remain migration-readable for safe upgrades and in-flight historical task reconciliation; no public subscription endpoints or new subscription-funded sessions remain.
