package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/712/D
// https://codeforces.com/problemset/status/712/problem/D
func TestCF712D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
1 2 2 1
outputCopy
6
inputCopy
1 1 1 2
outputCopy
31
inputCopy
2 12 3 1
outputCopy
0`
	testutil.AssertEqualCase(t, rawText, 0, CF712D)
}
