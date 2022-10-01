package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/767/C
// https://codeforces.com/problemset/status/767/problem/C
func TestCF767C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
6
2 4
0 5
4 2
2 1
1 1
4 2
outputCopy
1 4
inputCopy
6
2 4
0 6
4 2
2 1
1 1
4 2
outputCopy
-1`
	testutil.AssertEqualCase(t, rawText, 0, CF767C)
}
