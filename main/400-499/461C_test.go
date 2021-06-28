package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/461/C
// https://codeforces.com/problemset/status/461/problem/C
func TestCF461C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
7 4
1 3
1 2
2 0 1
2 1 2
outputCopy
4
3
inputCopy
10 9
2 2 9
1 1
2 0 1
1 8
2 0 8
1 2
2 1 3
1 4
2 2 4
outputCopy
7
2
10
4
5`
	testutil.AssertEqualCase(t, rawText, 0, CF461C)
}
