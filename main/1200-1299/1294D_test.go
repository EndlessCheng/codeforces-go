package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1294/D
// https://codeforces.com/problemset/status/1294/problem/D
func TestCF1294D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
7 3
0
1
2
2
0
0
10
outputCopy
1
2
3
3
4
4
7
inputCopy
4 3
1
2
1
2
outputCopy
0
0
0
0`
	testutil.AssertEqualCase(t, rawText, 0, CF1294D)
}
