package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1716/C
// https://codeforces.com/problemset/status/1716/problem/C
func TestCF1716C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4
3
0 0 1
4 3 2
5
0 4 8 12 16
2 6 10 14 18
4
0 10 10 10
10 10 10 10
2
0 0
0 0
outputCopy
5
19
17
3`
	testutil.AssertEqualCase(t, rawText, 0, CF1716C)
}
