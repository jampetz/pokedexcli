package main

import (
	"os"
)

func exitCommand(cfg *config, args ...string) error {
	os.Exit(0)
	return nil
}
