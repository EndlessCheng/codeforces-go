package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/500/B
// https://codeforces.com/problemset/status/500/problem/B
func TestCF500B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
7
5 2 4 3 6 7 1
0001001
0000000
0000010
1000001
0000000
0010000
1001000
outputCopy
1 2 4 3 6 7 5
inputCopy
5
4 2 1 5 3
00100
00011
10010
01101
01010
outputCopy
1 2 3 4 5
inputCopy
7
1 7 6 4 2 3 5
0000100
0000010
0000001
0000000
1000000
0100000
0010000
outputCopy
1 3 5 4 2 7 6`
	testutil.AssertEqualCase(t, rawText, -1, CF500B)
}
