package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1523/D
// https://codeforces.com/problemset/status/1523/problem/D
func TestCF1523D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3 4 3
1000
0110
1001
outputCopy
1000
inputCopy
5 5 4
11001
10101
10010
01110
11011
outputCopy
10001`
	testutil.DebugTLE = 0
	testutil.AssertEqualCase(t, rawText, 0, CF1523D)
}
