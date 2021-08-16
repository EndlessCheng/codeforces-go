package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/383/D
// https://codeforces.com/problemset/status/383/problem/D
func TestCF383D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4
1 1 1 1
outputCopy
12
inputCopy
2
1 1
outputCopy
2
inputCopy
3
1 1 1
outputCopy
4`
	testutil.AssertEqualCase(t, rawText, -1, CF383D)
}
