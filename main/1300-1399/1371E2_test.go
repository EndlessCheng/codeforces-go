package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1371/E2
// https://codeforces.com/problemset/status/1371/problem/E2
func TestCF1371E2(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3 2
3 4 5
outputCopy
1
3
inputCopy
4 3
2 3 5 6
outputCopy
2
3 4
inputCopy
4 3
9 1 1 1
outputCopy
0

inputCopy
3 2
1000000000 1 999999999
outputCopy
1
999999998
inputCopy
7 5
1 1 1 1 4 4 4
outputCopy
`
	testutil.AssertEqualCase(t, rawText, 0, CF1371E2)
}
