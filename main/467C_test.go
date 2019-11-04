package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestSol467C(t *testing.T) {
	// just copy from website
	rawText := `
6 2 2
6 5 4 3 2 1
outputCopy
18
inputCopy
5 2 1
1 2 3 4 5
outputCopy
9
inputCopy
7 1 3
2 10 7 18 5 33 0
outputCopy
61`
	testutil.AssertEqualCase(t, rawText, 0, Sol467C)
}
