package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1552/D
// https://codeforces.com/problemset/status/1552/problem/D
func TestCF1552D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5
5
4 -7 -1 5 10
1
0
3
1 10 100
4
-3 2 10 2
9
25 -171 250 174 152 242 100 -205 -258
outputCopy
YES
YES
NO
YES
YES`
	testutil.AssertEqualCase(t, rawText, 0, CF1552D)
}
