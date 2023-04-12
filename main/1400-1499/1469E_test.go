package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1469/E
// https://codeforces.com/problemset/status/1469/problem/E
func TestCF1469E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
7
4 2
0110
4 2
1001
9 3
010001110
9 3
101110001
10 3
0101110001
10 10
1111111111
11 10
11111111110
outputCopy
YES
11
YES
00
YES
010
YES
101
NO
YES
0000000001
YES
0000000010`
	testutil.AssertEqualCase(t, rawText, 0, CF1469E)
}
