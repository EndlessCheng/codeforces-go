package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1042/C
// https://codeforces.com/problemset/status/1042/problem/C
func TestCF1042C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5
5 -2 0 1 -3
outputCopy
2 3
1 1 2
1 2 4
1 4 5
inputCopy
5
5 2 0 4 0
outputCopy
1 3 5
2 5
1 1 2
1 2 4
inputCopy
2
2 -1
outputCopy
2 2
inputCopy
4
0 -10 0 0
outputCopy
1 1 2
1 2 3
1 3 4
inputCopy
4
0 0 0 0
outputCopy
1 1 2
1 2 3
1 3 4`
	testutil.AssertEqualCase(t, rawText, 0, CF1042C)
}
