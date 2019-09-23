package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestSol257C(t *testing.T) {
	// just copy from website
	rawText := `2
2 0
0 2
outputCopy
90.0000000000
inputCopy
3
2 0
0 2
-2 2
outputCopy
135.0000000000
inputCopy
4
2 0
0 2
-2 0
0 -2
outputCopy
270.0000000000
inputCopy
2
2 1
1 2
outputCopy
36.8698976458`
	testutil.AssertEqual(t, rawText, Sol257C)
}
