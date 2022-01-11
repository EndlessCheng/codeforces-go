package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/1027/problem/F
// https://codeforces.com/problemset/status/1027/problem/F
func TestCF1027F(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
2
1 5
1 7
outputCopy
5
inputCopy
3
5 13
1 5
1 7
outputCopy
7
inputCopy
3
10 40
40 80
10 80
outputCopy
80
inputCopy
3
99 100
99 100
99 100
outputCopy
-1`
	testutil.AssertEqualCase(t, rawText, 0, CF1027F)
}
