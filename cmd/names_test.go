package cmd

import (
	"os/exec"
	"strings"
	"testing"
)

// TestNoForeignProductNames fails when a tracked file names another object
// store. LICENSE, NOTICE and CREDITS carry legally required attribution.
func TestNoForeignProductNames(t *testing.T) {
	out, err := exec.Command("git", "-C", "..", "grep", "-lIE", "minio|MinIO|Minio|MINIO|min\\.io").Output()
	if err != nil && len(out) == 0 {
		t.Skip("not a git checkout")
	}
	for _, name := range strings.Fields(string(out)) {
		base := name[strings.LastIndex(name, "/")+1:]
		switch base {
		case "LICENSE", "NOTICE", "CREDITS", "go.sum", "names_test.go":
			continue
		}
		t.Errorf("%s names another product", name)
	}
}
