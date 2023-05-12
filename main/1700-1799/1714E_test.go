package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1714/E
// https://codeforces.com/problemset/status/1714/problem/E
func TestCF1714E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
10
2
6 11
3
2 18 22
5
5 10 5 10 5
4
1 2 4 8
2
4 5
3
93 96 102
2
40 6
2
50 30
2
22 44
2
1 5
outputCopy
Yes
No
Yes
Yes
No
Yes
No
No
Yes
No`
	testutil.AssertEqualCase(t, rawText, 0, CF1714E)
}
