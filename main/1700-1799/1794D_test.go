package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1794/D
// https://codeforces.com/problemset/status/1794/problem/D
func TestCF1794D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
2
1 3 2 3
outputCopy
2
inputCopy
2
2 2 3 5
outputCopy
5
inputCopy
1
1 4
outputCopy
0`
	testutil.AssertEqualCase(t, rawText, 0, CF1794D)
}
