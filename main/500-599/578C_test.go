package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestSol578C(t *testing.T) {
	// just copy from website
	rawText := `
3
1 2 3
outputCopy
1.000000000000000
inputCopy
4
1 2 3 4
outputCopy
2.000000000000000
inputCopy
10
1 10 2 9 3 8 4 7 5 6
outputCopy
4.500000000000000
inputCopy
12
0 0 0 0 0 0 0 0 0 0 0 0
outputCopy
0.000000000000000`
	testutil.AssertEqualCase(t, rawText, 0, Sol578C)
}
