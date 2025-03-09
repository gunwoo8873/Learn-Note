package bin

import (
	"fmt"
	"testing"
)

var (
	Version    = "dev"
	CommitHash = "n/a"
	BuildTime  = "n/a"
)

func TestBuildTimestamp(t *testing.T) {
	v := fmt.Sprintf("Version: %s\nCommit Hash: %s\nBuild Time: %s\n", Version, CommitHash, BuildTime)
	fmt.Printf(v)
	if gotest := "Version: dev\nCommit Hash: n/a\nBuild Time: n/a\n"; v != gotest {
		t.Errorf("Expected %s, got %s", gotest, v)
	}
}
