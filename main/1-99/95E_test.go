package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/95/E
// https://codeforces.com/problemset/status/95/problem/E
func TestCF95E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4 3
1 2
2 3
1 3
outputCopy
1
inputCopy
5 4
1 2
3 4
4 5
3 5
outputCopy
-1`
	testutil.AssertEqualCase(t, rawText, 0, CF95E)
}
