package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestSol913D(t *testing.T) {
	// just copy from website
	rawText := `
5 300
3 100
4 150
4 80
2 90
2 300
outputCopy
2
3
3 1 4
inputCopy
2 100
1 787
2 788
outputCopy
0
0

inputCopy
2 100
2 42
2 58
outputCopy
2
2
1 2`
	testutil.AssertEqualCase(t, rawText, -1, Sol913D)
}
