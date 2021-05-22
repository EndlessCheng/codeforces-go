package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/163/A
// https://codeforces.com/problemset/status/163/problem/A
func TestCF163A(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
aa
aa
outputCopy
5
inputCopy
codeforces
forceofcode
outputCopy
60`
	testutil.AssertEqualCase(t, rawText, 0, CF163A)
}
