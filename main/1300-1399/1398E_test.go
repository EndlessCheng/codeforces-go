package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1398/E
// https://codeforces.com/problemset/status/1398/problem/E
func TestCF1398E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
6
1 5
0 10
1 -5
0 5
1 11
0 -10
outputCopy
5
25
10
15
36
21
inputCopy
7
0 136177412
0 -136177412
0 455326434
1 14442996
0 -455326434
0 958682748
1 104290903
outputCopy
136177412
0
455326434
925095864
14442996
1931808492
2140390298`
	testutil.AssertEqualCase(t, rawText, -1, CF1398E)
}
