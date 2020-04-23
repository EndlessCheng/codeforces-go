package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"path/filepath"
	"testing"
)

func Test_run(t *testing.T) {
	dir, _ := filepath.Abs(".")
	testutil.AssertEqualFileCaseWithName(t, dir, "*.in", "*.out", 0, run)
}
