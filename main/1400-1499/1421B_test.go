package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1421B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
4
S010
0001
1000
111F
3
S10
101
01F
5
S0101
00000
01111
11111
0001F
outputCopy
1
3 4
2
1 2
2 1
0`
	testutil.AssertEqualCase(t, rawText, 0, CF1421B)
}
