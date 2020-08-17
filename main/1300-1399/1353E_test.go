package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1353E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
6
9 2
010001010
9 3
111100000
7 4
1111111
10 3
1001110101
1 1
1
1 1
0
outputCopy
1
2
5
4
0
0`
	testutil.AssertEqualCase(t, rawText, 0, CF1353E)
}
