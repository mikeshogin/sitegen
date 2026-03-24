# sitegen

Static site generator for AI agent projects. Zero API surface, pure HTML/CSS/JS output. Built on Hugo.

## Why Static

- No API endpoints = no attack surface
- No server runtime = no runtime vulnerabilities
- CDN-deployable = fast, global, free hosting
- Git-native = every change is tracked

If you need interactivity - use Telegram bots or third-party providers, not direct ports.

## What It Does

Generates documentation and product sites for the linter ecosystem:

```bash
# Generate site for a project
sitegen build --project archlint --output ./dist/

# Generate ecosystem landing page
sitegen ecosystem --config ecosystem.yaml --output ./dist/

# Generate all project sites at once
sitegen build-all --config ecosystem.yaml --output ./dist/
```

## Features

### Project Sites
- Auto-generated from README.md + ECOSYSTEM.md
- CLI reference from --help output
- Architecture diagrams (Mermaid -> SVG)
- Metrics dashboards (static charts from telemetry JSONL)
- Issue tracker summary (from GitHub API at build time)

### Ecosystem Landing Page
- All projects on one page with connections
- Pipeline diagram
- Quick start guide
- Cost savings calculator (static, pre-computed)

### Hugo Integration
- Hugo theme templates for consistent styling
- Shortcodes for Mermaid, code examples, metrics
- Multilingual (RU/EN) out of the box
- Zero JavaScript requirement (progressive enhancement)

## Architecture

```
Source files (README, ECOSYSTEM, telemetry)
    |
    v
sitegen (Go CLI)
    |
    +-- Parse README.md -> structured data
    +-- Parse ECOSYSTEM.md -> project graph
    +-- Parse telemetry.jsonl -> static charts
    +-- Fetch GitHub issues (at build time only)
    |
    v
Hugo templates + content
    |
    v
Static HTML/CSS/JS
    |
    v
Deploy: GitHub Pages / Cloudflare Pages / Nginx
```

## Config

```yaml
# ecosystem.yaml
name: "AI Agent Linter Ecosystem"
domain: "mshogin.ru"

projects:
  - name: archlint
    repo: mshogin/archlint
    description: "Architecture quality gate"
    docs: true

  - name: promptlint
    repo: mikeshogin/promptlint
    description: "Complexity-based model routing"
    docs: true

  - name: costlint
    repo: mikeshogin/costlint
    description: "Token cost analysis"
    docs: true

  - name: seclint
    repo: mikeshogin/seclint
    description: "Content safety filter"
    docs: true

theme: "linter-docs"
languages: [ru, en]
```

## Security Model

- Build time only: all external data fetched during `sitegen build`
- No runtime API calls
- No server-side processing
- No database
- No user input handling
- Output is pure static files

If dynamic interaction needed:
- Telegram bot for feedback/support
- GitHub Issues for bug reports
- Google Forms for surveys
- Third-party comment systems (optional)

## Integration with Ecosystem

sitegen consumes data from all other linters:
- archlint scan results -> architecture violation charts
- promptlint telemetry -> routing distribution graphs
- costlint reports -> cost savings dashboards
- seclint ratings -> content safety stats

All data is pre-computed at build time and embedded as static assets.

## Related Projects

- [archlint](https://github.com/mshogin/archlint) - architecture linter
- [promptlint](https://github.com/mikeshogin/promptlint) - complexity router
- [costlint](https://github.com/mikeshogin/costlint) - cost analysis
- [seclint](https://github.com/mikeshogin/seclint) - content filter
- [myhome](https://github.com/kgatilin/myhome) - agent orchestration

## Contributing

1. Fork and send a PR
2. Open an issue with ideas
3. Barter: review exchange
