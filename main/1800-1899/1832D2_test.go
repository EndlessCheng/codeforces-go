package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/1832/problem/D2
// https://codeforces.com/problemset/status/1832/problem/D2
func TestCF1832D2(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4 10
5 2 8 4
1 2 3 4 5 6 7 8 9 10
outputCopy
3 4 5 6 7 8 8 10 8 12
inputCopy
5 10
5 2 8 4 4
1 2 3 4 5 6 7 8 9 10
outputCopy
3 4 5 6 7 8 9 8 11 8
inputCopy
2 5
2 3
10 6 8 1 3
outputCopy
10 7 8 3 3`
	testutil.AssertEqualCase(t, rawText, -1, CF1832D2)
}
