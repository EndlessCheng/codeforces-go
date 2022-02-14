package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/1638/problem/C
// https://codeforces.com/problemset/status/1638/problem/C
func TestCF1638C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
6
3
1 2 3
5
2 1 4 3 5
6
6 1 4 2 5 3
1
1
6
3 2 1 6 5 4
5
3 1 5 2 4
outputCopy
3
3
1
1
2
1`
	testutil.AssertEqualCase(t, rawText, 0, CF1638C)
}
