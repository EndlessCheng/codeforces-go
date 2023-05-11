package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1809/D
// https://codeforces.com/problemset/status/1809/problem/D
func TestCF1809D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
6
100
0
0101
00101101
1001101
11111
outputCopy
1000000000001
0
1000000000000
2000000000001
2000000000002
0`
	testutil.AssertEqualCase(t, rawText, 0, CF1809D)
}
