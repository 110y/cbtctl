package command

import (
	"context"
	"fmt"
	"os"
)

func Exec() {
	os.Exit(exec(context.Background()))
}

func exec(ctx context.Context) int {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "must specify sub command")
		return 1
	}

	switch os.Args[1] {
	case runCmd:
		if err := run(ctx, os.Args[2:]); err != nil {
			fmt.Fprintf(os.Stderr, "failed to execute `run`: %s\n", err.Error())
			return 1
		}

	default:
		fmt.Fprintln(os.Stderr, "must specify valid sub command")
		return 1
	}

	return 0
}
