package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestSol1249E(t *testing.T) {
	// just copy from website
	rawText := `
10 2
7 6 18 6 16 18 1 17 17
6 9 3 10 9 1 10 1 5
outputCopy
0 7 13 18 24 35 36 37 40 45 
inputCopy
10 1
3 2 3 1 3 3 1 4 1
1 2 3 4 4 1 2 1 3
outputCopy
0 2 4 7 8 11 13 14 16 17 `
	testutil.AssertEqualCase(t, rawText, 0, Sol1249E)
}
