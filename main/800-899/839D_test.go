package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/839/D
// https://codeforces.com/problemset/status/839/problem/D
func TestCF839D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
3 3 1
outputCopy
12
inputCopy
4
2 3 4 6
outputCopy
39`
	testutil.AssertEqualCase(t, rawText, 0, CF839D)
}
