package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1383B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
3
1 2 2
3
2 2 3
5
0 0 0 2 2
outputCopy
WIN
LOSE
DRAW
inputCopy
4
5
4 1 5 1 3
4
1 0 1 6
1
0
2
5 4
outputCopy
WIN
WIN
DRAW
WIN`
	testutil.AssertEqualCase(t, rawText, 0, CF1383B)
}
