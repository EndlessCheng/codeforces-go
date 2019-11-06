package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestSol546D(t *testing.T) {
	// just copy from website
	rawText := `
2
3 1
6 3
outputCopy
2
5`
	testutil.AssertEqualCase(t, rawText, 0, Sol546D)
}
