package memda

import (
	"os"
	"testing"
)

// TestParseArgs Test parseArgs function
func TestParseArgs(t *testing.T) {
	os.Args = append(os.Args, "--profile=test")
	os.Args = append(os.Args, "--region=us-west-2")
	os.Args = append(os.Args, `--lambda=test`)
	os.Args = append(os.Args, "--limit=20")

	profile, region, lambda, limit := parseArgs()

	if profile != "test" {
		t.Fail()
	}
	if region != "us-west-2" {
		t.Fail()
	}
	if lambda != "test" {
		t.Fail()
	}
	if limit != 20 {
		t.Fail()
	}
}
