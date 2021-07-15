package mainutil

import (
	"context"
	"fmt"
	"os"
)

// Main is a helper for main functions.
func Main(run func(context.Context) error) {
	if err := run(context.Background()); err != nil {
		if errString := err.Error(); errString != "" {
			fmt.Fprintln(os.Stderr, errString)
		}
		os.Exit(1)
	}
}
