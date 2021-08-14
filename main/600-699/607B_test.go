package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/607/B
// https://codeforces.com/problemset/status/607/problem/B
func TestCF607B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
1 2 1
outputCopy
1
inputCopy
3
1 2 3
outputCopy
3
inputCopy
7
1 4 4 2 3 2 1
outputCopy
2`
	testutil.AssertEqualCase(t, rawText, 0, CF607B)
}
