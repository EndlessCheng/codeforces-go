package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/464/problem/B
// https://codeforces.com/problemset/status/464/problem/B
func TestCF464B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
0 0 0
0 0 1
0 0 1
0 0 1
0 1 1
0 1 1
0 1 1
1 1 1
outputCopy
YES
0 0 0
0 0 1
0 1 0
1 0 0
0 1 1
1 0 1
1 1 0
1 1 1
inputCopy
0 0 0
0 0 0
0 0 0
0 0 0
1 1 1
1 1 1
1 1 1
1 1 1
outputCopy
NO`
	testutil.AssertEqualCase(t, rawText, 0, CF464B)
}
