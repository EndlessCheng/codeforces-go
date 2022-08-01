package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/175/C
// https://codeforces.com/problemset/status/175/problem/C
func TestCF175C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
1
5 10
2
3 6
outputCopy
70
inputCopy
2
3 8
5 10
1
20
outputCopy
74
inputCopy
3
10 3
20 2
30 1
3
30 50 60
outputCopy
200
inputCopy
2
1000000000 1000
1 1
1
10
outputCopy
1999999991001`
	testutil.AssertEqualCase(t, rawText, -1, CF175C)
}
