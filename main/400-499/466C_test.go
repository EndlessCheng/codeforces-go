package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestSol466C(t *testing.T) {
	// just copy from website
	rawText := `
4
0 0 0 0
outputCopy
3
inputCopy
5
1 2 3 0 3
outputCopy
2
inputCopy
4
0 1 -1 0
outputCopy
1
inputCopy
2
4 1
outputCopy
0`
	testutil.AssertEqual(t, rawText, Sol466C)
}
