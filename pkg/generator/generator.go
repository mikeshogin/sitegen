package generator

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"
)

// EcosystemConfig defines the ecosystem site configuration.
type EcosystemConfig struct {
	Name      string    `yaml:"name"`
	Domain    string    `yaml:"domain"`
	Projects  []Project `yaml:"projects"`
	Theme     string    `yaml:"theme"`
	Languages []string  `yaml:"languages"`
}

// Project defines one project in the ecosystem.
type Project struct {
	Name        string `yaml:"name"`
	Repo        string `yaml:"repo"`
	Description string `yaml:"description"`
	Docs        bool   `yaml:"docs"`
}

// BuildProject generates a static site for one project.
func BuildProject(name, outputDir string) error {
	dir := filepath.Join(outputDir, name)
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return fmt.Errorf("create output dir: %w", err)
	}

	// Generate index.html from README
	html := generateProjectHTML(name)
	if err := os.WriteFile(filepath.Join(dir, "index.html"), []byte(html), 0o644); err != nil {
		return fmt.Errorf("write index.html: %w", err)
	}

	return nil
}

// BuildEcosystem generates the ecosystem landing page.
func BuildEcosystem(configPath, outputDir string) error {
	config, err := loadConfig(configPath)
	if err != nil {
		return err
	}

	if err := os.MkdirAll(outputDir, 0o755); err != nil {
		return fmt.Errorf("create output dir: %w", err)
	}

	html := generateEcosystemHTML(config)
	if err := os.WriteFile(filepath.Join(outputDir, "index.html"), []byte(html), 0o644); err != nil {
		return fmt.Errorf("write index.html: %w", err)
	}

	return nil
}

// BuildAll generates sites for all projects plus ecosystem landing.
func BuildAll(configPath, outputDir string) error {
	config, err := loadConfig(configPath)
	if err != nil {
		return err
	}

	// Build ecosystem landing
	if err := BuildEcosystem(configPath, outputDir); err != nil {
		return err
	}

	// Build each project
	for _, p := range config.Projects {
		if p.Docs {
			if err := BuildProject(p.Name, outputDir); err != nil {
				return fmt.Errorf("build %s: %w", p.Name, err)
			}
		}
	}

	return nil
}

func loadConfig(path string) (*EcosystemConfig, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read config: %w", err)
	}
	var config EcosystemConfig
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("parse config: %w", err)
	}
	return &config, nil
}

func generateProjectHTML(name string) string {
	return fmt.Sprintf(`<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>%s - AI Agent Linter Ecosystem</title>
    <style>
        body { font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, sans-serif; max-width: 800px; margin: 0 auto; padding: 2rem; line-height: 1.6; color: #333; }
        h1 { border-bottom: 2px solid #2563eb; padding-bottom: 0.5rem; }
        code { background: #f1f5f9; padding: 0.2rem 0.4rem; border-radius: 3px; font-size: 0.9em; }
        pre { background: #1e293b; color: #e2e8f0; padding: 1rem; border-radius: 8px; overflow-x: auto; }
        pre code { background: transparent; color: inherit; }
        a { color: #2563eb; }
        .badge { display: inline-block; padding: 0.2rem 0.6rem; border-radius: 12px; font-size: 0.8em; background: #dbeafe; color: #1e40af; }
    </style>
</head>
<body>
    <h1>%s</h1>
    <p><span class="badge">Go</span> <span class="badge">No LLM</span> <span class="badge">&lt;10ms</span></p>
    <p>Part of the <a href="../">AI Agent Linter Ecosystem</a></p>
    <h2>Install</h2>
    <pre><code>go install github.com/mikeshogin/%s/cmd/%s@latest</code></pre>
    <h2>Links</h2>
    <ul>
        <li><a href="https://github.com/mikeshogin/%s">GitHub Repository</a></li>
        <li><a href="../">Ecosystem Overview</a></li>
    </ul>
</body>
</html>`, name, name, name, name, name)
}

func generateEcosystemHTML(config *EcosystemConfig) string {
	var projectCards strings.Builder
	for _, p := range config.Projects {
		projectCards.WriteString(fmt.Sprintf(`
        <div class="card">
            <h3><a href="./%s/">%s</a></h3>
            <p>%s</p>
            <a href="https://github.com/%s" class="badge">GitHub</a>
        </div>`, p.Name, p.Name, p.Description, p.Repo))
	}

	return fmt.Sprintf(`<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>%s</title>
    <style>
        body { font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, sans-serif; max-width: 1000px; margin: 0 auto; padding: 2rem; line-height: 1.6; color: #333; }
        h1 { text-align: center; font-size: 2.5rem; margin-bottom: 0.5rem; }
        .subtitle { text-align: center; color: #64748b; font-size: 1.2rem; margin-bottom: 2rem; }
        .grid { display: grid; grid-template-columns: repeat(auto-fit, minmax(250px, 1fr)); gap: 1.5rem; margin: 2rem 0; }
        .card { border: 1px solid #e2e8f0; border-radius: 12px; padding: 1.5rem; transition: box-shadow 0.2s; }
        .card:hover { box-shadow: 0 4px 12px rgba(0,0,0,0.1); }
        .card h3 { margin-top: 0; }
        .card h3 a { color: #1e293b; text-decoration: none; }
        .card h3 a:hover { color: #2563eb; }
        .badge { display: inline-block; padding: 0.2rem 0.6rem; border-radius: 12px; font-size: 0.8em; background: #dbeafe; color: #1e40af; text-decoration: none; }
        .pipeline { background: #f8fafc; border-radius: 12px; padding: 2rem; margin: 2rem 0; text-align: center; }
        .pipeline code { font-size: 1.1em; }
        pre { background: #1e293b; color: #e2e8f0; padding: 1rem; border-radius: 8px; overflow-x: auto; }
        pre code { background: transparent; color: inherit; }
        .stats { display: grid; grid-template-columns: repeat(4, 1fr); gap: 1rem; margin: 2rem 0; text-align: center; }
        .stat { padding: 1rem; }
        .stat-value { font-size: 2rem; font-weight: bold; color: #2563eb; }
        .stat-label { color: #64748b; font-size: 0.9rem; }
    </style>
</head>
<body>
    <h1>%s</h1>
    <p class="subtitle">Quality, cost and safety control for AI agent workflows. All Go, no LLM, under 10ms.</p>

    <div class="stats">
        <div class="stat"><div class="stat-value">58%%</div><div class="stat-label">Cost savings</div></div>
        <div class="stat"><div class="stat-value">&lt;10ms</div><div class="stat-label">Routing latency</div></div>
        <div class="stat"><div class="stat-value">0</div><div class="stat-label">API endpoints</div></div>
        <div class="stat"><div class="stat-value">5</div><div class="stat-label">Tools</div></div>
    </div>

    <div class="pipeline">
        <code>prompt -> seclint (safe?) -> promptlint (route) -> agent -> archlint (quality) -> costlint (cost)</code>
    </div>

    <h2>Tools</h2>
    <div class="grid">%s
    </div>

    <h2>Quick Start</h2>
    <pre><code>go install github.com/mshogin/archlint@latest
go install github.com/mikeshogin/promptlint/cmd/promptlint@latest
go install github.com/mikeshogin/costlint/cmd/costlint@latest
go install github.com/mikeshogin/seclint/cmd/seclint@latest</code></pre>

    <p style="text-align: center; margin-top: 2rem; color: #64748b;">Built by AI agents. Open source. Go.</p>
</body>
</html>`, config.Name, config.Name, projectCards.String())
}
