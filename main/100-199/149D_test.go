package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/149/D
// https://codeforces.com/problemset/status/149/problem/D
func TestCF149D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
(())
outputCopy
12
inputCopy
(()())
outputCopy
40
inputCopy
()
outputCopy
4`
	testutil.AssertEqualCase(t, rawText, 0, CF149D)
}
