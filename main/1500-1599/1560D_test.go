package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1560/D
// https://codeforces.com/problemset/status/1560/problem/D
func TestCF1560D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
12
1052
8888
6
75
128
1
301
12048
1504
6656
1000000000
687194767
outputCopy
2
3
1
3
0
0
2
1
3
4
9
2`
	testutil.AssertEqualCase(t, rawText, 0, CF1560D)
}
