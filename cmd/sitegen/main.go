package main

import (
	"fmt"
	"os"

	"github.com/mikeshogin/sitegen/pkg/generator"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Usage: sitegen {build|ecosystem|build-all}\n\n")
		fmt.Fprintf(os.Stderr, "Commands:\n")
		fmt.Fprintf(os.Stderr, "  build --project NAME --output DIR    Build site for one project\n")
		fmt.Fprintf(os.Stderr, "  ecosystem --config FILE --output DIR Build ecosystem landing page\n")
		fmt.Fprintf(os.Stderr, "  build-all --config FILE --output DIR Build all project sites\n")
		os.Exit(1)
	}

	cmd := os.Args[1]

	switch cmd {
	case "build":
		project := flagValue("--project")
		output := flagValue("--output")
		if project == "" || output == "" {
			fmt.Fprintf(os.Stderr, "Usage: sitegen build --project NAME --output DIR\n")
			os.Exit(1)
		}
		if err := generator.BuildProject(project, output); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("Site generated: %s/%s/\n", output, project)

	case "ecosystem":
		config := flagValue("--config")
		output := flagValue("--output")
		if config == "" || output == "" {
			fmt.Fprintf(os.Stderr, "Usage: sitegen ecosystem --config FILE --output DIR\n")
			os.Exit(1)
		}
		if err := generator.BuildEcosystem(config, output); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("Ecosystem site generated: %s/\n", output)

	case "build-all":
		config := flagValue("--config")
		output := flagValue("--output")
		if config == "" || output == "" {
			fmt.Fprintf(os.Stderr, "Usage: sitegen build-all --config FILE --output DIR\n")
			os.Exit(1)
		}
		if err := generator.BuildAll(config, output); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("All sites generated: %s/\n", output)

	default:
		fmt.Fprintf(os.Stderr, "Unknown command: %s\n", cmd)
		os.Exit(1)
	}
}

func flagValue(name string) string {
	for i, arg := range os.Args {
		if arg == name && i+1 < len(os.Args) {
			return os.Args[i+1]
		}
	}
	return ""
}
