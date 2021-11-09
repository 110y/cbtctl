package command

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"strings"

	"github.com/110y/cbtctl/internal/gcp/build"
)

const runCmd = "run"

func run(ctx context.Context, args []string) error {
	fs := flag.NewFlagSet(runCmd, flag.ContinueOnError)

	project := fs.String("project", "", "GCP Project ID")
	trigger := fs.String("trigger", "", "Trigger ID")
	substitutions := fs.String("substitutions", "", "Substitutions")

	if err := fs.Parse(args); err != nil {
		return fmt.Errorf("failed to parse flags: %w", err)
	}

	substitutionsMap := make(map[string]string)
	if *substitutions != "" {
		for _, substitution := range strings.Split(*substitutions, ",") {
			kv := strings.Split(substitution, "=")
			if len(kv) != 2 {
				return errors.New("invalid substitutions")
			}

			substitutionsMap[kv[0]] = kv[1]
		}
	}

	if err := build.Run(ctx, *project, *trigger, substitutionsMap); err != nil {
		return fmt.Errorf("failed to run trigger: %w", err)
	}

	return nil
}
