package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestSol772B(t *testing.T) {
	// just copy from website
	rawText := `
4
0 0
0 1
1 1
1 0
outputCopy
0.3535533906
inputCopy
6
5 0
10 0
12 -4
10 -8
5 -8
3 -4
outputCopy
1.0000000000`
	testutil.AssertEqualCase(t, rawText, 0, Sol772B)
}
