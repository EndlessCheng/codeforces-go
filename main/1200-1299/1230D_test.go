package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestSol1230D(t *testing.T) {
	// just copy from website
	rawText := `
4
3 2 3 6
2 8 5 10
outputCopy
15
inputCopy
3
1 2 3
1 2 3
outputCopy
0
inputCopy
1
0
1
outputCopy
0`
	testutil.AssertEqual(t, rawText, Sol1230D)
}
