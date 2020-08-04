package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestSol4D(t *testing.T) {
	// just copy from website
	rawText := `
2 1 1
2 2
2 2
outputCopy
1
1 
inputCopy
3 3 3
5 4
12 11
9 8
outputCopy
3
1 3 2 `
	testutil.AssertEqualCase(t, rawText, 0, Sol4D)
}
