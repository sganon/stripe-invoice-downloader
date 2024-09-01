package main

import (
	"fmt"
	"os"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func run() error {
	app := initCliApp()

	if err := app.Run(os.Args); err != nil {
		return fmt.Errorf("error running app: %w", err)
	}

	return nil
}
