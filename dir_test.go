package main

import (
	"os"
	"testing"

	cli "gopkg.in/urfave/cli.v1"
)

func TestNewCommandDir(t *testing.T) {
	app := cli.NewApp()
	app.Name = "TestApp"
	NewCommandDir(app)
	if len(app.Commands) != 1 {
		t.Fatalf("NewCommandDir did not correctly append a new command to app")
	}
}

func TestDir(t *testing.T) {
	testApp := cli.NewApp()
	NewCommandDir(testApp)
	testApp.Run([]string{"testApp", "dir", "./tmp"})
	if os.Getenv(GolnsDirKey) != "./tmp" {
		t.Fatalf("Dir command failed to correctly set the environment variable")
	}
}
