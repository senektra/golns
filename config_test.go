package main

import (
	"os"
	"testing"
)

func TestNewConfig(t *testing.T) {
	// Set GOLNS_DIR to an empty string.
	os.Setenv(GolnsDirKey, "")
	c, err := NewConfig()
	if err == nil {
		t.Fatalf("Error wasn't set on empty %s key (value of %s: %s)",
			GolnsDirKey, GolnsDirKey, os.Getenv(GolnsDirKey))
	}

	// Set GOLNS_DIR to nonempty "path."
	os.Setenv(GolnsDirKey, "a/dummy/path")
	c, err = NewConfig()
	if err != nil {
		t.Fatalf("Error was set on nonempty %s key (value of %s: %s)",
			GolnsDirKey, GolnsDirKey, os.Getenv(GolnsDirKey))
	}
	if c.targetDir != "a/dummy/path" {
		t.Fatalf("Incorrect targetDir set in config, %s got set instead.", c.targetDir)
	}
}
