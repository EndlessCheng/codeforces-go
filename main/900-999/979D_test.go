package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/979/D
// https://codeforces.com/problemset/status/979/problem/D
func TestCF979D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5
1 1
1 2
2 1 1 3
2 1 1 2
2 1 1 1
outputCopy
2
1
-1
inputCopy
10
1 9
2 9 9 22
2 3 3 18
1 25
2 9 9 20
2 25 25 14
1 20
2 26 26 3
1 14
2 20 20 9
outputCopy
9
9
9
-1
-1
-1`
	testutil.AssertEqualCase(t, rawText, 0, CF979D)
}
