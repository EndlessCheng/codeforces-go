package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1633/A
// https://codeforces.com/problemset/status/1633/problem/A
func TestCF1633A(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
42
23
377
outputCopy
42
28
777`
	testutil.AssertEqualCase(t, rawText, 0, CF1633A)
}
