package main

import (
	"errors"
	"os"
)

// GolnsDirKey is the environment identifier for the target directory.
const GolnsDirKey = "GOLNS_DIR"

// Config the target and temporary directory. The temporary directory is used for
// the 'dir' command.
type Config struct {
	targetDir string
}

// NewConfig returns a new config.
func NewConfig() (*Config, error) {
	golnsDir := os.Getenv(GolnsDirKey)
	if golnsDir == "" {
		return nil, errors.New(GolnsDirKey + " is empty. (is it exported?)")
	}
	return &Config{targetDir: golnsDir}, nil
}
