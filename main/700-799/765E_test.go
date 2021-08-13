package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/765/E
// https://codeforces.com/problemset/status/765/problem/E
func TestCF765E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
6
1 2
2 3
2 4
4 5
1 6
outputCopy
3
inputCopy
7
1 2
1 3
3 4
1 5
5 6
6 7
outputCopy
-1`
	testutil.AssertEqualCase(t, rawText, 0, CF765E)
}
